# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: Data.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\nData.proto\x12\x02pb\"\x96\x01\n\x0b\x44\x61taRequest\x12\x0e\n\x06\x63logid\x18\x01 \x01(\x05\x12\r\n\x05\x63neid\x18\x02 \x01(\x05\x12\x11\n\tclineport\x18\x03 \x01(\t\x12\x15\n\rclocationinfo\x18\x04 \x01(\t\x12\x15\n\rcoccurutctime\x18\x05 \x01(\t\x12\x12\n\ncalarmcode\x18\x06 \x01(\x05\x12\x13\n\x0b\x63\x61larmlevel\x18\x07 \x01(\x05\"*\n\tDataReply\x12\x1d\n\x04\x64\x61ta\x18\x01 \x03(\x0b\x32\x0f.pb.DataRequest2;\n\x0b\x44\x61taService\x12,\n\x08SendData\x12\x0f.pb.DataRequest\x1a\r.pb.DataReply\"\x00\x62\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'Data_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  _DATAREQUEST._serialized_start=19
  _DATAREQUEST._serialized_end=169
  _DATAREPLY._serialized_start=171
  _DATAREPLY._serialized_end=213
  _DATASERVICE._serialized_start=215
  _DATASERVICE._serialized_end=274
# @@protoc_insertion_point(module_scope)
