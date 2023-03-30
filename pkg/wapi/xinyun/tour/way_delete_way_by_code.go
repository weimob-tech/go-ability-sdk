package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// WayDeleteWayByCode
// @id 995
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除线路产品)
func (client *Client) WayDeleteWayByCode(request *WayDeleteWayByCodeRequest) (response *WayDeleteWayByCodeResponse, err error) {
	rpcResponse := CreateWayDeleteWayByCodeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WayDeleteWayByCodeRequest struct {
	*api.BaseRequest

	GoodsCode string `json:"goodsCode,omitempty"`
	GoodsName string `json:"goodsName,omitempty"`
}

type WayDeleteWayByCodeResponse struct {
}

func CreateWayDeleteWayByCodeRequest() (request *WayDeleteWayByCodeRequest) {
	request = &WayDeleteWayByCodeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "way/deleteWayByCode", "api")
	request.Method = api.POST
	return
}

func CreateWayDeleteWayByCodeResponse() (response *api.BaseResponse[WayDeleteWayByCodeResponse]) {
	response = api.CreateResponse[WayDeleteWayByCodeResponse](&WayDeleteWayByCodeResponse{})
	return
}
