# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: dex/reserve_0_and_shares_type.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='dex/reserve_0_and_shares_type.proto',
  package='nicholasdotsol.duality.dex',
  syntax='proto3',
  serialized_pb=_b('\n#dex/reserve_0_and_shares_type.proto\x12\x1anicholasdotsol.duality.dex\x1a\x14gogoproto/gogo.proto\"\xe2\x01\n\x15Reserve0AndSharesType\x12_\n\x08reserve0\x18\x01 \x01(\tBM\xf2\xde\x1f\x0fyaml:\"reserve0\"\xda\xde\x1f&github.com/cosmos/cosmos-sdk/types.Dec\xc8\xde\x1f\x00\xea\xde\x1f\x08reserve0\x12h\n\x0btotalShares\x18\x02 \x01(\tBS\xf2\xde\x1f\x12yaml:\"totalShares\"\xda\xde\x1f&github.com/cosmos/cosmos-sdk/types.Dec\xc8\xde\x1f\x00\xea\xde\x1f\x0btotalSharesB/Z-github.com/NicholasDotSol/duality/x/dex/typesb\x06proto3')
  ,
  dependencies=[gogoproto_dot_gogo__pb2.DESCRIPTOR,])
_sym_db.RegisterFileDescriptor(DESCRIPTOR)




_RESERVE0ANDSHARESTYPE = _descriptor.Descriptor(
  name='Reserve0AndSharesType',
  full_name='nicholasdotsol.duality.dex.Reserve0AndSharesType',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='reserve0', full_name='nicholasdotsol.duality.dex.Reserve0AndSharesType.reserve0', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\362\336\037\017yaml:\"reserve0\"\332\336\037&github.com/cosmos/cosmos-sdk/types.Dec\310\336\037\000\352\336\037\010reserve0'))),
    _descriptor.FieldDescriptor(
      name='totalShares', full_name='nicholasdotsol.duality.dex.Reserve0AndSharesType.totalShares', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\362\336\037\022yaml:\"totalShares\"\332\336\037&github.com/cosmos/cosmos-sdk/types.Dec\310\336\037\000\352\336\037\013totalShares'))),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=90,
  serialized_end=316,
)

DESCRIPTOR.message_types_by_name['Reserve0AndSharesType'] = _RESERVE0ANDSHARESTYPE

Reserve0AndSharesType = _reflection.GeneratedProtocolMessageType('Reserve0AndSharesType', (_message.Message,), dict(
  DESCRIPTOR = _RESERVE0ANDSHARESTYPE,
  __module__ = 'dex.reserve_0_and_shares_type_pb2'
  # @@protoc_insertion_point(class_scope:nicholasdotsol.duality.dex.Reserve0AndSharesType)
  ))
_sym_db.RegisterMessage(Reserve0AndSharesType)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z-github.com/NicholasDotSol/duality/x/dex/types'))
_RESERVE0ANDSHARESTYPE.fields_by_name['reserve0'].has_options = True
_RESERVE0ANDSHARESTYPE.fields_by_name['reserve0']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\362\336\037\017yaml:\"reserve0\"\332\336\037&github.com/cosmos/cosmos-sdk/types.Dec\310\336\037\000\352\336\037\010reserve0'))
_RESERVE0ANDSHARESTYPE.fields_by_name['totalShares'].has_options = True
_RESERVE0ANDSHARESTYPE.fields_by_name['totalShares']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\362\336\037\022yaml:\"totalShares\"\332\336\037&github.com/cosmos/cosmos-sdk/types.Dec\310\336\037\000\352\336\037\013totalShares'))
# @@protoc_insertion_point(module_scope)
