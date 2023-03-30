package weimob_guide

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerCardGuiderGetList
// @id 1379
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1379?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=批量查询开卡导购)
func (client *Client) CustomerCardGuiderGetList(request *CustomerCardGuiderGetListRequest) (response *CustomerCardGuiderGetListResponse, err error) {
	rpcResponse := CreateCustomerCardGuiderGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerCardGuiderGetListRequest struct {
	*api.BaseRequest

	CardExtendList []CustomerCardGuiderGetListRequestCardExtendList `json:"cardExtendList,omitempty"`
	Tcode          string                                           `json:"tcode,omitempty"`
}

type CustomerCardGuiderGetListResponse struct {
	List []CustomerCardGuiderGetListResponselist `json:"list,omitempty"`
}

func CreateCustomerCardGuiderGetListRequest() (request *CustomerCardGuiderGetListRequest) {
	request = &CustomerCardGuiderGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_guide", "v2.0", "customer/card/guider/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerCardGuiderGetListResponse() (response *api.BaseResponse[CustomerCardGuiderGetListResponse]) {
	response = api.CreateResponse[CustomerCardGuiderGetListResponse](&CustomerCardGuiderGetListResponse{})
	return
}

type CustomerCardGuiderGetListRequestCardExtendList struct {
	CustomerWid int64  `json:"customerWid,omitempty"`
	Extension   string `json:"extension,omitempty"`
}

type CustomerCardGuiderGetListResponselist struct {
	CustomerWid int64  `json:"customerWid,omitempty"`
	GuiderWid   int64  `json:"guiderWid,omitempty"`
	GuiderId    string `json:"guiderId,omitempty"`
	GuiderVid   int64  `json:"guiderVid,omitempty"`
	GuiderName  string `json:"guiderName,omitempty"`
	GuiderPhone string `json:"guiderPhone,omitempty"`
	Extension   string `json:"extension,omitempty"`
}
