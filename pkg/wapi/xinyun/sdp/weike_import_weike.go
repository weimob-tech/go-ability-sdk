package sdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// WeikeImportWeike
// @id 1287
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=指定成为微客)
func (client *Client) WeikeImportWeike(request *WeikeImportWeikeRequest) (response *WeikeImportWeikeResponse, err error) {
	rpcResponse := CreateWeikeImportWeikeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WeikeImportWeikeRequest struct {
	*api.BaseRequest

	WidList []int `json:"widList,omitempty"`
	BizType int   `json:"bizType,omitempty"`
	BizId   int64 `json:"bizId,omitempty"`
}

type WeikeImportWeikeResponse struct {
}

func CreateWeikeImportWeikeRequest() (request *WeikeImportWeikeRequest) {
	request = &WeikeImportWeikeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("sdp", "1_0", "weike/importWeike", "api")
	request.Method = api.POST
	return
}

func CreateWeikeImportWeikeResponse() (response *api.BaseResponse[WeikeImportWeikeResponse]) {
	response = api.CreateResponse[WeikeImportWeikeResponse](&WeikeImportWeikeResponse{})
	return
}
