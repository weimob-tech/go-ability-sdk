package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// WechatMiniprogramQrcodeCreate
// @id 2115
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2115?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=通过appid创建微信小程序二维码)
func (client *Client) WechatMiniprogramQrcodeCreate(request *WechatMiniprogramQrcodeCreateRequest) (response *WechatMiniprogramQrcodeCreateResponse, err error) {
	rpcResponse := CreateWechatMiniprogramQrcodeCreateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WechatMiniprogramQrcodeCreateRequest struct {
	*api.BaseRequest

	AuthorizerAppid string `json:"authorizerAppid,omitempty"`
	Page            string `json:"page,omitempty"`
	Scene           string `json:"scene,omitempty"`
	Width           int64  `json:"width,omitempty"`
}

type WechatMiniprogramQrcodeCreateResponse struct {
	QrcodeUrl string `json:"qrcodeUrl,omitempty"`
}

func CreateWechatMiniprogramQrcodeCreateRequest() (request *WechatMiniprogramQrcodeCreateRequest) {
	request = &WechatMiniprogramQrcodeCreateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "wechat/miniprogram/qrcode/create", "apigw")
	request.Method = api.POST
	return
}

func CreateWechatMiniprogramQrcodeCreateResponse() (response *api.BaseResponse[WechatMiniprogramQrcodeCreateResponse]) {
	response = api.CreateResponse[WechatMiniprogramQrcodeCreateResponse](&WechatMiniprogramQrcodeCreateResponse{})
	return
}
