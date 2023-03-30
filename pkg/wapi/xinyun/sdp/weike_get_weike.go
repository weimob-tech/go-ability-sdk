package sdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// WeikeGetWeike
// @id 740
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取微客的基本信息)
func (client *Client) WeikeGetWeike(request *WeikeGetWeikeRequest) (response *WeikeGetWeikeResponse, err error) {
	rpcResponse := CreateWeikeGetWeikeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WeikeGetWeikeRequest struct {
	*api.BaseRequest

	Wid int64 `json:"wid,omitempty"`
}

type WeikeGetWeikeResponse struct {
	Level        WeikeGetWeikeResponseLevel   `json:"level,omitempty"`
	Inviter      WeikeGetWeikeResponseInviter `json:"inviter,omitempty"`
	State        byte                         `json:"state,omitempty"`
	RegisterTime int64                        `json:"registerTime,omitempty"`
}

func CreateWeikeGetWeikeRequest() (request *WeikeGetWeikeRequest) {
	request = &WeikeGetWeikeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("sdp", "1_0", "weike/getWeike", "api")
	request.Method = api.POST
	return
}

func CreateWeikeGetWeikeResponse() (response *api.BaseResponse[WeikeGetWeikeResponse]) {
	response = api.CreateResponse[WeikeGetWeikeResponse](&WeikeGetWeikeResponse{})
	return
}

type WeikeGetWeikeResponseLevel struct {
	Name string `json:"name,omitempty"`
	Id   int64  `json:"id,omitempty"`
}

type WeikeGetWeikeResponseInviter struct {
	Wid int64 `json:"wid,omitempty"`
}
