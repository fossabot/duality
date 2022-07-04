package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/NicholasDotSol/duality/testutil/keeper"
	"github.com/NicholasDotSol/duality/testutil/nullify"
	"github.com/NicholasDotSol/duality/x/dex/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestShareQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNShare(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetShareRequest
		response *types.QueryGetShareResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetShareRequest{
				Owner:  msgs[0].Owner,
				Token0: msgs[0].Token0,
				Token1: msgs[0].Token1,
				Price:  msgs[0].Price,
				Fee:    msgs[0].Fee,
			},
			response: &types.QueryGetShareResponse{Share: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetShareRequest{
				Owner:  msgs[1].Owner,
				Token0: msgs[1].Token0,
				Token1: msgs[1].Token1,
				Price:  msgs[1].Price,
				Fee:    msgs[1].Fee,
			},
			response: &types.QueryGetShareResponse{Share: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetShareRequest{
				Owner:  strconv.Itoa(100000),
				Token0: strconv.Itoa(100000),
				Token1: strconv.Itoa(100000),
				Price:  strconv.Itoa(100000),
				Fee:    100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Share(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestShareQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNShare(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllShareRequest {
		return &types.QueryAllShareRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ShareAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Share), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Share),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ShareAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Share), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Share),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ShareAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Share),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ShareAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
