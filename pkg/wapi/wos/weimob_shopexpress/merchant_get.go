package weimob_shopexpress

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantGet
// @id 1956
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1956?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商城信息)
func (client *Client) MerchantGet(request *MerchantGetRequest) (response *MerchantGetResponse, err error) {
	rpcResponse := CreateMerchantGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantGetRequest struct {
	*api.BaseRequest
}

type MerchantGetResponse struct {
	MerchantName      string `json:"merchantName,omitempty"`
	Mail              string `json:"mail,omitempty"`
	UtcOffset         string `json:"utcOffset,omitempty"`
	UnitSystem        int64  `json:"unitSystem,omitempty"`
	DefaultWeightUnit int64  `json:"defaultWeightUnit,omitempty"`
	CurrencyCode      string `json:"currencyCode,omitempty"`
	CurrencyName      string `json:"currencyName,omitempty"`
	CurrencySign      string `json:"currencySign,omitempty"`
}

func CreateMerchantGetRequest() (request *MerchantGetRequest) {
	request = &MerchantGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shopexpress", "v2.0", "merchant/get", "apigw")
	request.Method = api.POST
	return
}

func CreateMerchantGetResponse() (response *api.BaseResponse[MerchantGetResponse]) {
	response = api.CreateResponse[MerchantGetResponse](&MerchantGetResponse{})
	return
}
