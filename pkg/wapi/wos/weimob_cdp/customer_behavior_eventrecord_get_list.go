package weimob_cdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerBehaviorEventrecordGetList
// @id 2095
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2095?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询客户行为明细（滚动）)
func (client *Client) CustomerBehaviorEventrecordGetList(request *CustomerBehaviorEventrecordGetListRequest) (response *CustomerBehaviorEventrecordGetListResponse, err error) {
	rpcResponse := CreateCustomerBehaviorEventrecordGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerBehaviorEventrecordGetListRequest struct {
	*api.BaseRequest

	BosId                   int64  `json:"bosId,omitempty"`
	TargetProductInstanceId int64  `json:"targetProductInstanceId,omitempty"`
	Limit                   int64  `json:"limit,omitempty"`
	SearchDate              string `json:"searchDate,omitempty"`
	CursorId                string `json:"cursorId,omitempty"`
}

type CustomerBehaviorEventrecordGetListResponse struct {
	List         []CustomerBehaviorEventrecordGetListResponselist `json:"list,omitempty"`
	HasNext      bool                                             `json:"hasNext,omitempty"`
	NextCursorId string                                           `json:"nextCursorId,omitempty"`
	Total        int64                                            `json:"total,omitempty"`
}

func CreateCustomerBehaviorEventrecordGetListRequest() (request *CustomerBehaviorEventrecordGetListRequest) {
	request = &CustomerBehaviorEventrecordGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_cdp", "v2.0", "customer/behavior/eventrecord/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerBehaviorEventrecordGetListResponse() (response *api.BaseResponse[CustomerBehaviorEventrecordGetListResponse]) {
	response = api.CreateResponse[CustomerBehaviorEventrecordGetListResponse](&CustomerBehaviorEventrecordGetListResponse{})
	return
}

type CustomerBehaviorEventrecordGetListResponselist struct {
	BosId                   string `json:"bosId,omitempty"`
	TargetProductInstanceId string `json:"targetProductInstanceId,omitempty"`
	Vid                     string `json:"vid,omitempty"`
	Ukey                    string `json:"ukey,omitempty"`
	SessionId               string `json:"sessionId,omitempty"`
	PageName                string `json:"pageName,omitempty"`
	SourceEnd               string `json:"sourceEnd,omitempty"`
	SourceName              string `json:"sourceName,omitempty"`
	VisitDuration           string `json:"visitDuration,omitempty"`
	GuideId                 string `json:"guideId,omitempty"`
	InType                  string `json:"inType,omitempty"`
	En                      string `json:"en,omitempty"`
	JsonData                string `json:"jsonData,omitempty"`
	CreateTime              int64  `json:"createTime,omitempty"`
}
