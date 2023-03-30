package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsSaveStatus
// @id 994
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=商品上下架)
func (client *Client) GoodsSaveStatus(request *GoodsSaveStatusRequest) (response *GoodsSaveStatusResponse, err error) {
	rpcResponse := CreateGoodsSaveStatusResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsSaveStatusRequest struct {
	*api.BaseRequest

	GoodsCode string `json:"goodsCode,omitempty"`
	Status    int    `json:"status,omitempty"`
}

type GoodsSaveStatusResponse struct {
}

func CreateGoodsSaveStatusRequest() (request *GoodsSaveStatusRequest) {
	request = &GoodsSaveStatusRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "goods/saveStatus", "api")
	request.Method = api.POST
	return
}

func CreateGoodsSaveStatusResponse() (response *api.BaseResponse[GoodsSaveStatusResponse]) {
	response = api.CreateResponse[GoodsSaveStatusResponse](&GoodsSaveStatusResponse{})
	return
}
