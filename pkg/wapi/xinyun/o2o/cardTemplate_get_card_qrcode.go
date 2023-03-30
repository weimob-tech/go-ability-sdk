package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CardTemplateGetCardQrcode
// @id 208
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取发优惠券二维码)
func (client *Client) CardTemplateGetCardQrcode(request *CardTemplateGetCardQrcodeRequest) (response *CardTemplateGetCardQrcodeResponse, err error) {
	rpcResponse := CreateCardTemplateGetCardQrcodeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CardTemplateGetCardQrcodeRequest struct {
	*api.BaseRequest

	MerchantId     int64 `json:"merchantId,omitempty"`
	CardTemplateId int64 `json:"cardTemplateId,omitempty"`
}

type CardTemplateGetCardQrcodeResponse struct {
}

func CreateCardTemplateGetCardQrcodeRequest() (request *CardTemplateGetCardQrcodeRequest) {
	request = &CardTemplateGetCardQrcodeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "cardTemplate/getCardQrcode", "api")
	request.Method = api.POST
	return
}

func CreateCardTemplateGetCardQrcodeResponse() (response *api.BaseResponse[CardTemplateGetCardQrcodeResponse]) {
	response = api.CreateResponse[CardTemplateGetCardQrcodeResponse](&CardTemplateGetCardQrcodeResponse{})
	return
}
