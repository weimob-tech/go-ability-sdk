package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GiftcardCheckCardPassword
// @id 1836
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=实体卡卡密校验)
func (client *Client) GiftcardCheckCardPassword(request *GiftcardCheckCardPasswordRequest) (response *GiftcardCheckCardPasswordResponse, err error) {
	rpcResponse := CreateGiftcardCheckCardPasswordResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GiftcardCheckCardPasswordRequest struct {
	*api.BaseRequest

	Data GiftcardCheckCardPasswordRequestData `json:"data,omitempty"`
}

type GiftcardCheckCardPasswordResponse struct {
}

func CreateGiftcardCheckCardPasswordRequest() (request *GiftcardCheckCardPasswordRequest) {
	request = &GiftcardCheckCardPasswordRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "giftcard/checkCardPassword", "api")
	request.Method = api.POST
	return
}

func CreateGiftcardCheckCardPasswordResponse() (response *api.BaseResponse[GiftcardCheckCardPasswordResponse]) {
	response = api.CreateResponse[GiftcardCheckCardPasswordResponse](&GiftcardCheckCardPasswordResponse{})
	return
}

type GiftcardCheckCardPasswordRequestData struct {
	CardNo  string `json:"cardNo,omitempty"`
	CardPwd string `json:"cardPwd,omitempty"`
}
