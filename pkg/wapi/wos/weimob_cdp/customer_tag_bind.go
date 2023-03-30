package weimob_cdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerTagBind
// @id 1558
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1558?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=给用户打标签)
func (client *Client) CustomerTagBind(request *CustomerTagBindRequest) (response *CustomerTagBindResponse, err error) {
	rpcResponse := CreateCustomerTagBindResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerTagBindRequest struct {
	*api.BaseRequest

	TagList           []CustomerTagBindRequestTagList `json:"tagList,omitempty"`
	WidList           []int64                         `json:"widList,omitempty"`
	NeedCheckExist    bool                            `json:"needCheckExist,omitempty"`
	VidType           int64                           `json:"vidType,omitempty"`
	Vid               int64                           `json:"vid,omitempty"`
	BosId             int64                           `json:"bosId,omitempty"`
	ProductInstanceId int64                           `json:"productInstanceId,omitempty"`
	MerchantId        int64                           `json:"merchantId,omitempty"`
	Tcode             string                          `json:"tcode,omitempty"`
}

type CustomerTagBindResponse struct {
	IsSuccess bool `json:"isSuccess,omitempty"`
}

func CreateCustomerTagBindRequest() (request *CustomerTagBindRequest) {
	request = &CustomerTagBindRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_cdp", "v2.0", "customer/tag/bind", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerTagBindResponse() (response *api.BaseResponse[CustomerTagBindResponse]) {
	response = api.CreateResponse[CustomerTagBindResponse](&CustomerTagBindResponse{})
	return
}

type CustomerTagBindRequestTagList struct {
	TagId      int64   `json:"tagId,omitempty"`
	AttrIdList []int64 `json:"attrIdList,omitempty"`
}
