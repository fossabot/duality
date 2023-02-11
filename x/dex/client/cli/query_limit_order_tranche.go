package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/spf13/cobra"
)

func CmdListLimitOrderTranche() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-limit-order-tranche [pair-id] [token-in]",
		Short: "list all LimitOrderTranches",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			argPairId := args[0]
			argTokenIn := args[1]

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllLimitOrderTrancheRequest{
				Pagination: pageReq,
				PairId:     argPairId,
				TokenIn:    argTokenIn,
			}

			res, err := queryClient.LimitOrderTrancheAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowLimitOrderTranche() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-limit-order-tranche [pair-id] [tick-index] [token-in] [tranche-index]",
		Short: "shows a LimitOrderTranche",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argPairId := args[0]
			argTickIndex := args[1]
			argTokenIn := args[2]
			argTrancheIndex := args[3]

			argTrancheIndexInt, err := strconv.Atoi(argTrancheIndex)

			if err != nil {
				return err
			}

			argTickIndexInt, err := strconv.Atoi(argTickIndex)

			if err != nil {
				return err
			}

			params := &types.QueryGetLimitOrderTrancheRequest{
				PairId:       argPairId,
				TickIndex:    int64(argTickIndexInt),
				TokenIn:      argTokenIn,
				TrancheIndex: uint64(argTrancheIndexInt),
			}

			res, err := queryClient.LimitOrderTranche(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}