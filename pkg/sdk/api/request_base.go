package api

import (
	"strings"
)

type BaseRequest struct {
	*baseRequest
	product      string
	version      string
	actionName   string
	endpointType string
}

func (request *BaseRequest) initRequest() {
	request.baseRequest = defaultBaseRequest()
	request.Method = POST
}

// /api/1_0/open/usercenter/getWeimobUserInfo
// /apigw/weimob_shop/v2.0/goods/unit/delete?accesstoken=ACCESS_TOKEN
// endpointType: api apigw
// product: weimob_shop
// version: V2.0
// action: goods/unit/delete
func (request *BaseRequest) InitWithApiInfo(product, version, action, endpointType string) {
	request.initRequest()
	request.product = product
	request.version = version
	request.actionName = action
	request.endpointType = endpointType
}

func (request *BaseRequest) BuildUrl(config *Config) string {
	schema := or(request.GetScheme(), config.defaultSchema)
	domain := or(request.GetDomain(), config.defaultDomain)
	path := request.buildPath()

	var sb strings.Builder
	sb.WriteString(schema)
	sb.WriteString("://")
	sb.WriteString(domain)
	sb.WriteString(path)
	return sb.String()
}

func (request *BaseRequest) BuildQueries() string {
	if len(request.QueryParams) == 0 {
		return ""
	}
	var head bool
	var sb strings.Builder
	for k, v := range request.QueryParams {
		if !head {
			head = true
		} else {
			sb.WriteString("&")
		}
		sb.WriteString(k)
		sb.WriteString("=")
		sb.WriteString(v)
	}
	return sb.String()
}

func (request *BaseRequest) buildPath() string {
	var sb strings.Builder
	sb.WriteString(PathSeparator)
	sb.WriteString(request.endpointType)
	sb.WriteString(PathSeparator)
	if len(request.product) > 0 {
		sb.WriteString(request.product)
		sb.WriteString(PathSeparator)
	}
	sb.WriteString(request.version)
	sb.WriteString(PathSeparator)
	sb.WriteString(request.actionName)
	return sb.String()
}

func or(a, b string) string {
	if a != "" {
		return a
	}
	return b
}

func (request *BaseRequest) GetRPCType() RPCType {
	return RPCTypeJson
}

type BaseMultipartRequest struct {
	BaseRequest
}

func (request *BaseMultipartRequest) GetRPCType() RPCType {
	return RPCTypeMultiPart
}

func (request *BaseMultipartRequest) initRequest() {
	request.baseRequest = defaultBaseRequest()
	request.ContentType = MultiPart
	request.ContentTypeByte = MultiPartByte
}
