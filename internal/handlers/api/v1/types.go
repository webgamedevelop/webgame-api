package v1

type Code int

func (o Code) ApplyToResponse(r Responser) {
	r.SetCode(int(o))
}

type Message string

func (o Message) ApplyToResponse(r Responser) {
	r.SetMessage(string(o))
}

type simpleResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (r *simpleResponse) SetCode(code int) {
	r.Code = code
}

func (r *simpleResponse) SetMessage(message string) {
	r.Message = message
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
