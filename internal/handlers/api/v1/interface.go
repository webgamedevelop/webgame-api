package v1

type Responser interface {
	SetCode(code int)
	SetMessage(message string)
}

type Option interface {
	ApplyToResponse(r Responser)
}

func SimpleResponse(r Responser, opts ...Option) Responser {
	for _, opt := range opts {
		opt.ApplyToResponse(r)
	}
	return r
}
