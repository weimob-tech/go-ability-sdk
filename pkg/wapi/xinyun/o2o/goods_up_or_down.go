package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsUpOrDown
// @id 66
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=上/下架菜品)
func (client *Client) GoodsUpOrDown(request *GoodsUpOrDownRequest) (response *GoodsUpOrDownResponse, err error) {
	rpcResponse := CreateGoodsUpOrDownResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsUpOrDownRequest struct {
	*api.BaseRequest

	GoodsId int64 `json:"goodsId,omitempty"`
	Status  int   `json:"status,omitempty"`
	StoreId int64 `json:"storeId,omitempty"`
}

type GoodsUpOrDownResponse struct {
}

func CreateGoodsUpOrDownRequest() (request *GoodsUpOrDownRequest) {
	request = &GoodsUpOrDownRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goods/upOrDown", "api")
	request.Method = api.POST
	return
}

func CreateGoodsUpOrDownResponse() (response *api.BaseResponse[GoodsUpOrDownResponse]) {
	response = api.CreateResponse[GoodsUpOrDownResponse](&GoodsUpOrDownResponse{})
	return
}
