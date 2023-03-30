package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsFindStockBySkuIdList
// @id 1837
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=根据skuIdList查询库存信息)
func (client *Client) GoodsFindStockBySkuIdList(request *GoodsFindStockBySkuIdListRequest) (response *GoodsFindStockBySkuIdListResponse, err error) {
	rpcResponse := CreateGoodsFindStockBySkuIdListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsFindStockBySkuIdListRequest struct {
	*api.BaseRequest

	SkuIdList []int `json:"skuIdList,omitempty"`
	StoreId   int64 `json:"storeId,omitempty"`
	PId       int64 `json:"pId,omitempty"`
}

type GoodsFindStockBySkuIdListResponse struct {
}

func CreateGoodsFindStockBySkuIdListRequest() (request *GoodsFindStockBySkuIdListRequest) {
	request = &GoodsFindStockBySkuIdListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goods/findStockBySkuIdList", "api")
	request.Method = api.POST
	return
}

func CreateGoodsFindStockBySkuIdListResponse() (response *api.BaseResponse[GoodsFindStockBySkuIdListResponse]) {
	response = api.CreateResponse[GoodsFindStockBySkuIdListResponse](&GoodsFindStockBySkuIdListResponse{})
	return
}
