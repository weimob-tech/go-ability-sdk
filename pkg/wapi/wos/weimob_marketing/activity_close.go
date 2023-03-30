package weimob_marketing

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ActivityClose
// @id 1765
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1765?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=关闭活动)
func (client *Client) ActivityClose(request *ActivityCloseRequest) (response *ActivityCloseResponse, err error) {
	rpcResponse := CreateActivityCloseResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ActivityCloseRequest struct {
	*api.BaseRequest

	BosId             int64 `json:"bosId,omitempty"`
	ProductId         int64 `json:"productId,omitempty"`
	ProductInstanceId int64 `json:"productInstanceId,omitempty"`
	Id                int64 `json:"id,omitempty"`
	Vid               int64 `json:"vid,omitempty"`
}

type ActivityCloseResponse struct {
}

func CreateActivityCloseRequest() (request *ActivityCloseRequest) {
	request = &ActivityCloseRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_marketing", "v2.0", "activity/close", "apigw")
	request.Method = api.POST
	return
}

func CreateActivityCloseResponse() (response *api.BaseResponse[ActivityCloseResponse]) {
	response = api.CreateResponse[ActivityCloseResponse](&ActivityCloseResponse{})
	return
}
