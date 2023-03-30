package weimob_aewc

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerGet
// @id 2222
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2222?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查看客户详情)
func (client *Client) CustomerGet(request *CustomerGetRequest) (response *CustomerGetResponse, err error) {
	rpcResponse := CreateCustomerGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerGetRequest struct {
	*api.BaseRequest

	ExternalUserId string `json:"externalUserId,omitempty"`
	CustomerId     string `json:"customerId,omitempty"`
	PageNum        int64  `json:"pageNum,omitempty"`
	PageSize       int64  `json:"pageSize,omitempty"`
}

type CustomerGetResponse struct {
	TagDetailList    []CustomerGetResponseTagDetailList `json:"tagDetailList,omitempty"`
	CustomerId       string                             `json:"customerId,omitempty"`
	ExternalUserId   string                             `json:"externalUserId,omitempty"`
	ExternalUserType int64                              `json:"externalUserType,omitempty"`
	NickName         string                             `json:"nickName,omitempty"`
	Avatar           string                             `json:"avatar,omitempty"`
	Gender           int64                              `json:"gender,omitempty"`
	Phone            string                             `json:"phone,omitempty"`
	CustomerType     int64                              `json:"customerType,omitempty"`
	FlowStatus       bool                               `json:"flowStatus,omitempty"`
	LatestAddTime    int64                              `json:"latestAddTime,omitempty"`
	AddWayDesc       string                             `json:"addWayDesc,omitempty"`
	GroupCount       int64                              `json:"groupCount,omitempty"`
}

func CreateCustomerGetRequest() (request *CustomerGetRequest) {
	request = &CustomerGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_aewc", "v2.0", "customer/get", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerGetResponse() (response *api.BaseResponse[CustomerGetResponse]) {
	response = api.CreateResponse[CustomerGetResponse](&CustomerGetResponse{})
	return
}

type CustomerGetResponseTagDetailList struct {
	TagName      string `json:"tagName,omitempty"`
	TagSource    int64  `json:"tagSource,omitempty"`
	TagGroupName string `json:"tagGroupName,omitempty"`
}
