package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SettingSaveWayTag
// @id 1012
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增/修改线路产品标签)
func (client *Client) SettingSaveWayTag(request *SettingSaveWayTagRequest) (response *SettingSaveWayTagResponse, err error) {
	rpcResponse := CreateSettingSaveWayTagResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SettingSaveWayTagRequest struct {
	*api.BaseRequest

	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type int    `json:"type,omitempty"`
}

type SettingSaveWayTagResponse struct {
}

func CreateSettingSaveWayTagRequest() (request *SettingSaveWayTagRequest) {
	request = &SettingSaveWayTagRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "setting/saveWayTag", "api")
	request.Method = api.POST
	return
}

func CreateSettingSaveWayTagResponse() (response *api.BaseResponse[SettingSaveWayTagResponse]) {
	response = api.CreateResponse[SettingSaveWayTagResponse](&SettingSaveWayTagResponse{})
	return
}
