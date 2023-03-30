package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsGroupDel
// @id 60
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除菜品分组)
func (client *Client) GoodsGroupDel(request *GoodsGroupDelRequest) (response *GoodsGroupDelResponse, err error) {
	rpcResponse := CreateGoodsGroupDelResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsGroupDelRequest struct {
	*api.BaseRequest

	GoodsGroupId int64 `json:"goodsGroupId,omitempty"`
	StoreId      int64 `json:"storeId,omitempty"`
	MerchantId   int64 `json:"merchantId,omitempty"`
}

type GoodsGroupDelResponse struct {
}

func CreateGoodsGroupDelRequest() (request *GoodsGroupDelRequest) {
	request = &GoodsGroupDelRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goodsGroup/del", "api")
	request.Method = api.POST
	return
}

func CreateGoodsGroupDelResponse() (response *api.BaseResponse[GoodsGroupDelResponse]) {
	response = api.CreateResponse[GoodsGroupDelResponse](&GoodsGroupDelResponse{})
	return
}
