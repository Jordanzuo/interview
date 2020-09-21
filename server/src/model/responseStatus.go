package model

// ResponseStatus ... Response status type, which is just an alias of string.
type ResponseStatus = string

// Add the response status for a request
const (
	Success              = ""
	DisconnectStatus     = "DisconnectStatus"
	ParamTypeError       = "ParamTypeError"
	ClientDataError      = "ClientDataError"
	ClientRequestExpired = "ClientRequestExpired"
	NoTargetMethod       = "NoTargetMethod"
	ParamNotMatch        = "ParamNotMatch"
	ParamInValid         = "ParamInValid"
	NoLogin              = "NoLogin"
	RuntimeException     = "RuntimeException"
	NoAvailableRoom      = "NoAvailableRoom"
	PlayerNameNotExists  = "PlayerNameNotExists"
)
