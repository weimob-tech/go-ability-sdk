package spi

type InvocationCode struct {
	Errcode string `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

var SuccessCode = InvocationCode{
	Errcode: "0",
	Errmsg:  "success",
}

type InvocationResponse[T any] struct {
	Code InvocationCode `json:"code,omitempty"`
	Data *T             `json:"data,omitempty"`
}

func Ok[T any](v *T) InvocationResponse[T] {
	return InvocationResponse[T]{
		Code: SuccessCode,
		Data: v,
	}
}
