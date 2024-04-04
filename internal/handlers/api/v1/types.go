package v1

type Code int

func (o Code) ApplyToResponse(r Responser) {
	r.SetCode(int(o))
}

type Message string

func (o Message) ApplyToResponse(r Responser) {
	r.SetMessage(string(o))
}

type ResponseExtend struct {
	Data interface{} `json:"data,omitempty"`
}

func (e *ResponseExtend) ApplyToResponse(r Responser) {
	r.SetExtend(e)
}

func Response(r Responser, opts ...Option) Responser {
	for _, opt := range opts {
		opt.ApplyToResponse(r)
	}
	return r
}

type simpleResponse struct {
	Code            int    `json:"code,omitempty"`
	Message         string `json:"message,omitempty"`
	*ResponseExtend `json:",inline,omitempty"`
}

func (r *simpleResponse) SetCode(code int) {
	r.Code = code
}

func (r *simpleResponse) SetMessage(message string) {
	r.Message = message
}

func (r *simpleResponse) SetExtend(ext *ResponseExtend) {
	r.ResponseExtend = ext
}

type detailResponse[T any] struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

type listResponse[T []E, E any] struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Len     int    `json:"len"`
	Data    T      `json:"data"`
}
