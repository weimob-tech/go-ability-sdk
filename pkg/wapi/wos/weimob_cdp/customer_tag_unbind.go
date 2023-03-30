package weimob_cdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerTagUnbind
// @id 1559
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1559?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除用户标签)
func (client *Client) CustomerTagUnbind(request *CustomerTagUnbindRequest) (response *CustomerTagUnbindResponse, err error) {
	rpcResponse := CreateCustomerTagUnbindResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerTagUnbindRequest struct {
	*api.BaseRequest

	TagList           []CustomerTagUnbindRequestTagList `json:"tagList,omitempty"`
	WidList           []int64                           `json:"widList,omitempty"`
	VidType           int64                             `json:"vidType,omitempty"`
	Vid               int64                             `json:"vid,omitempty"`
	BosId             int64                             `json:"bosId,omitempty"`
	ProductInstanceId int64                             `json:"productInstanceId,omitempty"`
	MerchantId        int64                             `json:"merchantId,omitempty"`
	Tcode             string                            `json:"tcode,omitempty"`
}

type CustomerTagUnbindResponse struct {
	IsSuccess bool `json:"isSuccess,omitempty"`
}

func CreateCustomerTagUnbindRequest() (request *CustomerTagUnbindRequest) {
	request = &CustomerTagUnbindRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_cdp", "v2.0", "customer/tag/unbind", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerTagUnbindResponse() (response *api.BaseResponse[CustomerTagUnbindResponse]) {
	response = api.CreateResponse[CustomerTagUnbindResponse](&CustomerTagUnbindResponse{})
	return
}

type CustomerTagUnbindRequestTagList struct {
	TagId      int64   `json:"tagId,omitempty"`
	AttrIdList []int64 `json:"attrIdList,omitempty"`
}
