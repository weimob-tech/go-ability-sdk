package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ActivityOpenRechargeQuery
// @id 1911
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=储值规则查询接口)
func (client *Client) ActivityOpenRechargeQuery(request *ActivityOpenRechargeQueryRequest) (response *ActivityOpenRechargeQueryResponse, err error) {
	rpcResponse := CreateActivityOpenRechargeQueryResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ActivityOpenRechargeQueryRequest struct {
	*api.BaseRequest

	Wid     int64 `json:"wid,omitempty"`
	StoreId int64 `json:"storeId,omitempty"`
}

type ActivityOpenRechargeQueryResponse struct {
}

func CreateActivityOpenRechargeQueryRequest() (request *ActivityOpenRechargeQueryRequest) {
	request = &ActivityOpenRechargeQueryRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "activityOpen/rechargeQuery", "api")
	request.Method = api.POST
	return
}

func CreateActivityOpenRechargeQueryResponse() (response *api.BaseResponse[ActivityOpenRechargeQueryResponse]) {
	response = api.CreateResponse[ActivityOpenRechargeQueryResponse](&ActivityOpenRechargeQueryResponse{})
	return
}
