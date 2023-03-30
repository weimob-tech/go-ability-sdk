package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NicheSalesStageList
// @id 1733
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询商机销售阶段)
func (client *Client) NicheSalesStageList(request *NicheSalesStageListRequest) (response *NicheSalesStageListResponse, err error) {
	rpcResponse := CreateNicheSalesStageListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NicheSalesStageListRequest struct {
	*api.BaseRequest

	StageProcessId int   `json:"stageProcessId,omitempty"`
	Wid            int64 `json:"wid,omitempty"`
}

type NicheSalesStageListResponse struct {
}

func CreateNicheSalesStageListRequest() (request *NicheSalesStageListRequest) {
	request = &NicheSalesStageListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "niche/salesStageList", "api")
	request.Method = api.POST
	return
}

func CreateNicheSalesStageListResponse() (response *api.BaseResponse[NicheSalesStageListResponse]) {
	response = api.CreateResponse[NicheSalesStageListResponse](&NicheSalesStageListResponse{})
	return
}
