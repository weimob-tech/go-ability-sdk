package sdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// WeikeGetLevelList
// @id 1328
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查看微客等级列表)
func (client *Client) WeikeGetLevelList(request *WeikeGetLevelListRequest) (response *WeikeGetLevelListResponse, err error) {
	rpcResponse := CreateWeikeGetLevelListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WeikeGetLevelListRequest struct {
	*api.BaseRequest
}

type WeikeGetLevelListResponse struct {
	LevelList []int64 `json:"levelList,omitempty"`
	LevelId   int     `json:"levelId,omitempty"`
	LevelName string  `json:"levelName,omitempty"`
}

func CreateWeikeGetLevelListRequest() (request *WeikeGetLevelListRequest) {
	request = &WeikeGetLevelListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("sdp", "1_0", "weike/getLevelList", "api")
	request.Method = api.POST
	return
}

func CreateWeikeGetLevelListResponse() (response *api.BaseResponse[WeikeGetLevelListResponse]) {
	response = api.CreateResponse[WeikeGetLevelListResponse](&WeikeGetLevelListResponse{})
	return
}
