package weimob_marketing

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ActivityDelete
// @id 1763
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1763?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除活动)
func (client *Client) ActivityDelete(request *ActivityDeleteRequest) (response *ActivityDeleteResponse, err error) {
	rpcResponse := CreateActivityDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ActivityDeleteRequest struct {
	*api.BaseRequest

	BosId             int64 `json:"bosId,omitempty"`
	ProductId         int64 `json:"productId,omitempty"`
	ProductInstanceId int64 `json:"productInstanceId,omitempty"`
	Id                int64 `json:"id,omitempty"`
	Vid               int64 `json:"vid,omitempty"`
}

type ActivityDeleteResponse struct {
}

func CreateActivityDeleteRequest() (request *ActivityDeleteRequest) {
	request = &ActivityDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_marketing", "v2.0", "activity/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateActivityDeleteResponse() (response *api.BaseResponse[ActivityDeleteResponse]) {
	response = api.CreateResponse[ActivityDeleteResponse](&ActivityDeleteResponse{})
	return
}
