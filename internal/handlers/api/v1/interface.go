package v1

type Responser interface {
	SetCode(code int)
	SetMessage(message string)
	SetExtend(ext ResponseExtender)
}

type Option interface {
	ApplyToResponse(r Responser)
}

type ResponseExtender interface {
	Option
}
