package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsAgree
// @id 1846
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1846?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商家同意售后)
func (client *Client) RightsAgree(request *RightsAgreeRequest) (response *RightsAgreeResponse, err error) {
	rpcResponse := CreateRightsAgreeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsAgreeRequest struct {
	*api.BaseRequest

	ReturnAddress RightsAgreeRequestReturnAddress `json:"returnAddress,omitempty"`
	Remark        string                          `json:"remark,omitempty"`
	RightsId      int64                           `json:"rightsId,omitempty"`
}

type RightsAgreeResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateRightsAgreeRequest() (request *RightsAgreeRequest) {
	request = &RightsAgreeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "rights/agree", "apigw")
	request.Method = api.POST
	return
}

func CreateRightsAgreeResponse() (response *api.BaseResponse[RightsAgreeResponse]) {
	response = api.CreateResponse[RightsAgreeResponse](&RightsAgreeResponse{})
	return
}

type RightsAgreeRequestReturnAddress struct {
	ReceiverAddress string `json:"receiverAddress,omitempty"`
	ReceiverName    string `json:"receiverName,omitempty"`
	ReceiverPhone   string `json:"receiverPhone,omitempty"`
	ProvinceName    string `json:"provinceName,omitempty"`
	CityName        string `json:"cityName,omitempty"`
	CountyName      string `json:"countyName,omitempty"`
	AreaName        string `json:"areaName,omitempty"`
}
