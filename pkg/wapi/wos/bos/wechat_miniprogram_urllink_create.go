package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// WechatMiniprogramUrllinkCreate
// @id 2116
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2116?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=通过appid创建微信小程序UrlLink)
func (client *Client) WechatMiniprogramUrllinkCreate(request *WechatMiniprogramUrllinkCreateRequest) (response *WechatMiniprogramUrllinkCreateResponse, err error) {
	rpcResponse := CreateWechatMiniprogramUrllinkCreateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WechatMiniprogramUrllinkCreateRequest struct {
	*api.BaseRequest

	AuthorizerAppid string `json:"authorizerAppid,omitempty"`
	Path            string `json:"path,omitempty"`
	Query           string `json:"query,omitempty"`
	ExpireType      int64  `json:"expireType,omitempty"`
	ExpireTime      int64  `json:"expireTime,omitempty"`
	ExpireInterval  int64  `json:"expireInterval,omitempty"`
}

type WechatMiniprogramUrllinkCreateResponse struct {
	UrlLink string `json:"urlLink,omitempty"`
}

func CreateWechatMiniprogramUrllinkCreateRequest() (request *WechatMiniprogramUrllinkCreateRequest) {
	request = &WechatMiniprogramUrllinkCreateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "wechat/miniprogram/urllink/create", "apigw")
	request.Method = api.POST
	return
}

func CreateWechatMiniprogramUrllinkCreateResponse() (response *api.BaseResponse[WechatMiniprogramUrllinkCreateResponse]) {
	response = api.CreateResponse[WechatMiniprogramUrllinkCreateResponse](&WechatMiniprogramUrllinkCreateResponse{})
	return
}
