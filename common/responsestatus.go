package common

import (
	"metrics/utils/response"
)

var Success = &response.BaseResponse{Code: "000000", Desc: "success"}
var Failure = &response.BaseResponse{Code: "199999", Desc: "failure"}
var NoUpdate = &response.BaseResponse{Code: "000001", Desc: "no update"}
var UnknownErr = &response.BaseResponse{Code: "999999", Desc: "unknown error"}
var ServiceNotFound = &response.BaseResponse{Code: "100001", Desc: "service not found"}
var ProtocolNotFound = &response.BaseResponse{Code: "100002", Desc: "protocol not found"}
var BlackList = &response.BaseResponse{Code: "100003", Desc: "black list"}

var ParamInvalid = &response.BaseResponse{Code: "110001", Desc: "param invalid"}
var ParamMissing = &response.BaseResponse{Code: "110002", Desc: "miss required parameters"}
var ParamPostBodyEmpty = &response.BaseResponse{Code: "110003", Desc: "post body is empty"}
var ParamPostBodyReadError = &response.BaseResponse{Code: "110004", Desc: "post body read error"}

var DeserializeError = &response.BaseResponse{Code: "120001", Desc: "deserialize error"}
var SerializeError = &response.BaseResponse{Code: "120002", Desc: "serialize error"}
var DeCompressError = &response.BaseResponse{Code: "120003", Desc: "decompress error"}
var CompressError = &response.BaseResponse{Code: "120004", Desc: "compress error"}
var EncryptError = &response.BaseResponse{Code: "120005", Desc: "encrypt error"}
var DecryptError = &response.BaseResponse{Code: "120006", Desc: "decrypt error"}

var ThirdServiceError = &response.BaseResponse{Code: "130001", Desc: "third service call error"}
var CacheError = &response.BaseResponse{Code: "130002", Desc: "cache error"}
var PersistenceError = &response.BaseResponse{Code: "130003", Desc: "persistence error"}
