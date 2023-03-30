package api

const (
	HTTP  = "HTTP"
	HTTPS = "HTTPS"

	DefaultHttpPort = "80"

	GET     = "GET"
	PUT     = "PUT"
	POST    = "POST"
	DELETE  = "DELETE"
	PATCH   = "PATCH"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"

	Json      = "application/json"
	Xml       = "application/xml"
	Raw       = "application/octet-stream"
	Form      = "application/x-www-form-urlencoded"
	MultiPart = "multipart/form-data"

	PathSeparator = "/"
)

type RPCType int

const (
	RPCTypeJson RPCType = iota
	RPCTypeMultiPart
)

var (
	JsonByte      = []byte(Json)
	XmlByte       = []byte(Xml)
	MultiPartByte = []byte(MultiPart)
)

type RpcRequest interface {
	GetScheme() string
	GetMethod() string
	GetDomain() string
	GetContentType() string
	GetContentTypeByte() []byte
	GetHeaders() map[string]string
	GetQueryParams() map[string]string
	GetContent() []byte
	BuildUrl(*Config) string
	BuildQueries() string
	GetRPCType() RPCType
	GetShopId() string
	GetShopType() string
}

type baseRequest struct {
	// config
	ShopId          string
	ShopType        string
	Scheme          string
	Method          string
	Domain          string
	ContentType     string
	ContentTypeByte []byte
	// ReadTimeout    time.Duration
	// ConnectTimeout time.Duration

	// payload
	QueryParams map[string]string
	Headers     map[string]string
	Content     []byte
	// todo trace
	// span opentracing.Span
}

func (request *baseRequest) GetShopId() string {
	return request.ShopId
}

func (request *baseRequest) GetShopType() string {
	return request.ShopType
}

func (request *baseRequest) GetScheme() string {
	return request.Scheme
}

func (request *baseRequest) GetDomain() string {
	return request.Domain
}

func (request *baseRequest) GetMethod() string {
	return request.Method
}

func (request *baseRequest) GetHeaders() map[string]string {
	return request.Headers
}

func (request *baseRequest) GetContent() []byte {
	return request.Content
}

func (request *baseRequest) GetQueryParams() map[string]string {
	return request.QueryParams
}

func (request *baseRequest) GetContentType() string {
	return request.ContentType
}

func (request *baseRequest) GetContentTypeByte() []byte {
	return request.ContentTypeByte
}

func defaultBaseRequest() (request *baseRequest) {
	request = &baseRequest{
		Scheme:          "",
		Method:          POST,
		ContentType:     Json,
		ContentTypeByte: JsonByte,
		Headers:         map[string]string{},
		QueryParams:     make(map[string]string),
	}
	return
}
