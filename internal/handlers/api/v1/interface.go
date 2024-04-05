package v1

type Responser interface {
	SetCode(code int)
	SetMessage(message string)
}

type Option interface {
	ApplyToResponse(r Responser)
}
