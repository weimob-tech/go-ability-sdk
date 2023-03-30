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
	Code InvocationCode
	Data *T
}

func Ok[T any](v *T) InvocationResponse[T] {
	return InvocationResponse[T]{
		Code: SuccessCode,
		Data: v,
	}
}
