package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UrlQrcodeGet
// @id 2209
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2209?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取链接及二维码)
func (client *Client) UrlQrcodeGet(request *UrlQrcodeGetRequest) (response *UrlQrcodeGetResponse, err error) {
	rpcResponse := CreateUrlQrcodeGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UrlQrcodeGetRequest struct {
	*api.BaseRequest

	BasicInfo UrlQrcodeGetRequestBasicInfo `json:"basicInfo,omitempty"`
	ParamMap  UrlQrcodeGetRequestParamMap  `json:"paramMap,omitempty"`
	ExtendMap UrlQrcodeGetRequestExtendMap `json:"extendMap,omitempty"`
	Refer     string                       `json:"refer,omitempty"`
	Channel   int64                        `json:"channel,omitempty"`
}

type UrlQrcodeGetResponse struct {
	Url    string `json:"url,omitempty"`
	QrCode string `json:"qrCode,omitempty"`
}

func CreateUrlQrcodeGetRequest() (request *UrlQrcodeGetRequest) {
	request = &UrlQrcodeGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "url/qrcode/get", "apigw")
	request.Method = api.POST
	return
}

func CreateUrlQrcodeGetResponse() (response *api.BaseResponse[UrlQrcodeGetResponse]) {
	response = api.CreateResponse[UrlQrcodeGetResponse](&UrlQrcodeGetResponse{})
	return
}

type UrlQrcodeGetRequestBasicInfo struct {
	Vid             int64 `json:"vid,omitempty"`
	BosId           int64 `json:"bosId,omitempty"`
	OperationSource int64 `json:"operationSource,omitempty"`
}

type UrlQrcodeGetRequestParamMap struct {
	VID int64 `json:"VID,omitempty"`
}

type UrlQrcodeGetRequestExtendMap struct {
}
