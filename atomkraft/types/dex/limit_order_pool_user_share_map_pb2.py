# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: dex/limit_order_pool_user_share_map.proto

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
  name='dex/limit_order_pool_user_share_map.proto',
  package='nicholasdotsol.duality.dex',
  syntax='proto3',
  serialized_pb=_b('\n)dex/limit_order_pool_user_share_map.proto\x12\x1anicholasdotsol.duality.dex\x1a\x14gogoproto/gogo.proto\"\xd8\x01\n\x1aLimitOrderPoolUserShareMap\x12\x0e\n\x06pairId\x18\x01 \x01(\t\x12\r\n\x05token\x18\x02 \x01(\t\x12\x11\n\ttickIndex\x18\x03 \x01(\x03\x12\r\n\x05\x63ount\x18\x04 \x01(\x04\x12\x0f\n\x07\x61\x64\x64ress\x18\x05 \x01(\t\x12h\n\x0bsharesOwned\x18\x06 \x01(\tBS\xf2\xde\x1f\x12yaml:\"sharesOwned\"\xda\xde\x1f&github.com/cosmos/cosmos-sdk/types.Dec\xc8\xde\x1f\x00\xea\xde\x1f\x0bsharesOwnedB/Z-github.com/NicholasDotSol/duality/x/dex/typesb\x06proto3')
  ,
  dependencies=[gogoproto_dot_gogo__pb2.DESCRIPTOR,])
_sym_db.RegisterFileDescriptor(DESCRIPTOR)




_LIMITORDERPOOLUSERSHAREMAP = _descriptor.Descriptor(
  name='LimitOrderPoolUserShareMap',
  full_name='nicholasdotsol.duality.dex.LimitOrderPoolUserShareMap',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='pairId', full_name='nicholasdotsol.duality.dex.LimitOrderPoolUserShareMap.pairId', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='token', full_name='nicholasdotsol.duality.dex.LimitOrderPoolUserShareMap.token', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='tickIndex', full_name='nicholasdotsol.duality.dex.LimitOrderPoolUserShareMap.tickIndex', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='count', full_name='nicholasdotsol.duality.dex.LimitOrderPoolUserShareMap.count', index=3,
      number=4, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='address', full_name='nicholasdotsol.duality.dex.LimitOrderPoolUserShareMap.address', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='sharesOwned', full_name='nicholasdotsol.duality.dex.LimitOrderPoolUserShareMap.sharesOwned', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\362\336\037\022yaml:\"sharesOwned\"\332\336\037&github.com/cosmos/cosmos-sdk/types.Dec\310\336\037\000\352\336\037\013sharesOwned'))),
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
  serialized_start=96,
  serialized_end=312,
)

DESCRIPTOR.message_types_by_name['LimitOrderPoolUserShareMap'] = _LIMITORDERPOOLUSERSHAREMAP

LimitOrderPoolUserShareMap = _reflection.GeneratedProtocolMessageType('LimitOrderPoolUserShareMap', (_message.Message,), dict(
  DESCRIPTOR = _LIMITORDERPOOLUSERSHAREMAP,
  __module__ = 'dex.limit_order_pool_user_share_map_pb2'
  # @@protoc_insertion_point(class_scope:nicholasdotsol.duality.dex.LimitOrderPoolUserShareMap)
  ))
_sym_db.RegisterMessage(LimitOrderPoolUserShareMap)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z-github.com/NicholasDotSol/duality/x/dex/types'))
_LIMITORDERPOOLUSERSHAREMAP.fields_by_name['sharesOwned'].has_options = True
_LIMITORDERPOOLUSERSHAREMAP.fields_by_name['sharesOwned']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\362\336\037\022yaml:\"sharesOwned\"\332\336\037&github.com/cosmos/cosmos-sdk/types.Dec\310\336\037\000\352\336\037\013sharesOwned'))
# @@protoc_insertion_point(module_scope)
