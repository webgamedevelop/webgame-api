package v1

type Responser interface {
	SetCode(code int)
	SetMessage(message string)
	SetExtend(ext *ResponseExtend)
}

type Option interface {
	ApplyToResponse(r Responser)
}
