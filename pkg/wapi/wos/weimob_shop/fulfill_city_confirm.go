package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FulfillCityConfirm
// @id 1731
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1731?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=同城确认送达)
func (client *Client) FulfillCityConfirm(request *FulfillCityConfirmRequest) (response *FulfillCityConfirmResponse, err error) {
	rpcResponse := CreateFulfillCityConfirmResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FulfillCityConfirmRequest struct {
	*api.BaseRequest

	BasicInfo FulfillCityConfirmRequestBasicInfo `json:"basicInfo,omitempty"`
	OrderNo   int64                              `json:"orderNo,omitempty"`
	FulfillNo int64                              `json:"fulfillNo,omitempty"`
}

type FulfillCityConfirmResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateFulfillCityConfirmRequest() (request *FulfillCityConfirmRequest) {
	request = &FulfillCityConfirmRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "fulfill/city/confirm", "apigw")
	request.Method = api.POST
	return
}

func CreateFulfillCityConfirmResponse() (response *api.BaseResponse[FulfillCityConfirmResponse]) {
	response = api.CreateResponse[FulfillCityConfirmResponse](&FulfillCityConfirmResponse{})
	return
}

type FulfillCityConfirmRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}
