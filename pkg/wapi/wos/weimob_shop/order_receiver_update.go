package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderReceiverUpdate
// @id 1814
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1814?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=修改收货人信息)
func (client *Client) OrderReceiverUpdate(request *OrderReceiverUpdateRequest) (response *OrderReceiverUpdateResponse, err error) {
	rpcResponse := CreateOrderReceiverUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderReceiverUpdateRequest struct {
	*api.BaseRequest

	ReceiverInfo OrderReceiverUpdateRequestReceiverInfo `json:"receiverInfo,omitempty"`
	OrderNo      int64                                  `json:"orderNo,omitempty"`
}

type OrderReceiverUpdateResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateOrderReceiverUpdateRequest() (request *OrderReceiverUpdateRequest) {
	request = &OrderReceiverUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "order/receiver/update", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderReceiverUpdateResponse() (response *api.BaseResponse[OrderReceiverUpdateResponse]) {
	response = api.CreateResponse[OrderReceiverUpdateResponse](&OrderReceiverUpdateResponse{})
	return
}

type OrderReceiverUpdateRequestReceiverInfo struct {
	Name         string `json:"name,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Province     string `json:"province,omitempty"`
	City         string `json:"city,omitempty"`
	County       string `json:"county,omitempty"`
	Area         string `json:"area,omitempty"`
	Address      string `json:"address,omitempty"`
	Zip          string `json:"zip,omitempty"`
	ProvinceCode string `json:"provinceCode,omitempty"`
	CityCode     string `json:"cityCode,omitempty"`
	CountyCode   string `json:"countyCode,omitempty"`
	AreaCode     string `json:"areaCode,omitempty"`
}
