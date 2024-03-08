package v1

type ResponseExtend struct {
	Data interface{} `json:"data"`
}

func (e *ResponseExtend) ApplyToResponse(r Responser) {
	r.SetExtend(e)
}

type simpleResponse struct {
	Code    int              `json:"code,omitempty"`
	Message string           `json:"message,omitempty"`
	Extend  ResponseExtender `json:"extend,omitempty"`
}

func (r *simpleResponse) SetCode(code int) {
	r.Code = code
}

func (r *simpleResponse) SetMessage(message string) {
	r.Message = message
}

func (r *simpleResponse) SetExtend(ext ResponseExtender) {
	r.Extend = ext
}

type Code int

func (o Code) ApplyToResponse(r Responser) {
	r.SetCode(int(o))
}

type Message string

func (o Message) ApplyToResponse(r Responser) {
	r.SetMessage(string(o))
}

func Response(r Responser, opts ...Option) Responser {
	for _, opt := range opts {
		opt.ApplyToResponse(r)
	}
	return r
}
