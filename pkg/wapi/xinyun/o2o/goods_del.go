package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsDel
// @id 65
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除菜品)
func (client *Client) GoodsDel(request *GoodsDelRequest) (response *GoodsDelResponse, err error) {
	rpcResponse := CreateGoodsDelResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsDelRequest struct {
	*api.BaseRequest

	GoodsId int64 `json:"goodsId,omitempty"`
	StoreId int64 `json:"storeId,omitempty"`
}

type GoodsDelResponse struct {
}

func CreateGoodsDelRequest() (request *GoodsDelRequest) {
	request = &GoodsDelRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goods/del", "api")
	request.Method = api.POST
	return
}

func CreateGoodsDelResponse() (response *api.BaseResponse[GoodsDelResponse]) {
	response = api.CreateResponse[GoodsDelResponse](&GoodsDelResponse{})
	return
}
