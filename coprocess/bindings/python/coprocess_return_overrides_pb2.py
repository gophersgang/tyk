# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: coprocess_return_overrides.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='coprocess_return_overrides.proto',
  package='coprocess',
  syntax='proto3',
  serialized_pb=_b('\n coprocess_return_overrides.proto\x12\tcoprocess\"@\n\x0fReturnOverrides\x12\x15\n\rresponse_code\x18\x01 \x01(\x05\x12\x16\n\x0eresponse_error\x18\x02 \x01(\tB\x05H\x03\xf8\x01\x01\x62\x06proto3')
)
_sym_db.RegisterFileDescriptor(DESCRIPTOR)




_RETURNOVERRIDES = _descriptor.Descriptor(
  name='ReturnOverrides',
  full_name='coprocess.ReturnOverrides',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='response_code', full_name='coprocess.ReturnOverrides.response_code', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='response_error', full_name='coprocess.ReturnOverrides.response_error', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
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
  serialized_start=47,
  serialized_end=111,
)

DESCRIPTOR.message_types_by_name['ReturnOverrides'] = _RETURNOVERRIDES

ReturnOverrides = _reflection.GeneratedProtocolMessageType('ReturnOverrides', (_message.Message,), dict(
  DESCRIPTOR = _RETURNOVERRIDES,
  __module__ = 'coprocess_return_overrides_pb2'
  # @@protoc_insertion_point(class_scope:coprocess.ReturnOverrides)
  ))
_sym_db.RegisterMessage(ReturnOverrides)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('H\003\370\001\001'))
# @@protoc_insertion_point(module_scope)
