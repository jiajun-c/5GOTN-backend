from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AlertRequest(_message.Message):
    __slots__ = ["datarequest"]
    DATAREQUEST_FIELD_NUMBER: _ClassVar[int]
    datarequest: _containers.RepeatedCompositeFieldContainer[Datainput]
    def __init__(self, datarequest: _Optional[_Iterable[_Union[Datainput, _Mapping]]] = ...) -> None: ...

class AlertResponse(_message.Message):
    __slots__ = ["dataresponse"]
    DATARESPONSE_FIELD_NUMBER: _ClassVar[int]
    dataresponse: _containers.RepeatedCompositeFieldContainer[Dataoutput]
    def __init__(self, dataresponse: _Optional[_Iterable[_Union[Dataoutput, _Mapping]]] = ...) -> None: ...

class AnalyseRequest(_message.Message):
    __slots__ = ["data"]
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[DataStruct]
    def __init__(self, data: _Optional[_Iterable[_Union[DataStruct, _Mapping]]] = ...) -> None: ...

class AnalyseResponse(_message.Message):
    __slots__ = ["all_account", "level0", "level1", "level2", "level3"]
    ALL_ACCOUNT_FIELD_NUMBER: _ClassVar[int]
    LEVEL0_FIELD_NUMBER: _ClassVar[int]
    LEVEL1_FIELD_NUMBER: _ClassVar[int]
    LEVEL2_FIELD_NUMBER: _ClassVar[int]
    LEVEL3_FIELD_NUMBER: _ClassVar[int]
    all_account: int
    level0: int
    level1: int
    level2: int
    level3: int
    def __init__(self, all_account: _Optional[int] = ..., level0: _Optional[int] = ..., level1: _Optional[int] = ..., level2: _Optional[int] = ..., level3: _Optional[int] = ...) -> None: ...

class DataStruct(_message.Message):
    __slots__ = ["calarmcode", "calarmlevel", "clineport", "clocationinfo", "clogid", "cneid", "coccurutctime"]
    CALARMCODE_FIELD_NUMBER: _ClassVar[int]
    CALARMLEVEL_FIELD_NUMBER: _ClassVar[int]
    CLINEPORT_FIELD_NUMBER: _ClassVar[int]
    CLOCATIONINFO_FIELD_NUMBER: _ClassVar[int]
    CLOGID_FIELD_NUMBER: _ClassVar[int]
    CNEID_FIELD_NUMBER: _ClassVar[int]
    COCCURUTCTIME_FIELD_NUMBER: _ClassVar[int]
    calarmcode: int
    calarmlevel: int
    clineport: str
    clocationinfo: str
    clogid: int
    cneid: int
    coccurutctime: str
    def __init__(self, clogid: _Optional[int] = ..., calarmcode: _Optional[int] = ..., cneid: _Optional[int] = ..., calarmlevel: _Optional[int] = ..., coccurutctime: _Optional[str] = ..., clocationinfo: _Optional[str] = ..., clineport: _Optional[str] = ...) -> None: ...

class Datainput(_message.Message):
    __slots__ = ["clocationinfo"]
    CLOCATIONINFO_FIELD_NUMBER: _ClassVar[int]
    clocationinfo: str
    def __init__(self, clocationinfo: _Optional[str] = ...) -> None: ...

class Dataoutput(_message.Message):
    __slots__ = ["clocationinfo"]
    CLOCATIONINFO_FIELD_NUMBER: _ClassVar[int]
    clocationinfo: str
    def __init__(self, clocationinfo: _Optional[str] = ...) -> None: ...

class TrainRequest(_message.Message):
    __slots__ = ["active"]
    ACTIVE_FIELD_NUMBER: _ClassVar[int]
    active: bool
    def __init__(self, active: bool = ...) -> None: ...

class TrainResponse(_message.Message):
    __slots__ = ["success"]
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    success: bool
    def __init__(self, success: bool = ...) -> None: ...
