package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RetailGoodsFindStoreByGoodsIdWithPage
// @id 2577
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=滚动查询商品可售门店)
func (client *Client) RetailGoodsFindStoreByGoodsIdWithPage(request *RetailGoodsFindStoreByGoodsIdWithPageRequest) (response *RetailGoodsFindStoreByGoodsIdWithPageResponse, err error) {
	rpcResponse := CreateRetailGoodsFindStoreByGoodsIdWithPageResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RetailGoodsFindStoreByGoodsIdWithPageRequest struct {
	*api.BaseRequest

	GoodsId      int64 `json:"goodsId,omitempty"`
	StartStoreId int   `json:"startStoreId,omitempty"`
	PageSize     int   `json:"pageSize,omitempty"`
	Pid          int64 `json:"pid,omitempty"`
	StoreId      int64 `json:"storeId,omitempty"`
	Wid          int64 `json:"wid,omitempty"`
}

type RetailGoodsFindStoreByGoodsIdWithPageResponse struct {
}

func CreateRetailGoodsFindStoreByGoodsIdWithPageRequest() (request *RetailGoodsFindStoreByGoodsIdWithPageRequest) {
	request = &RetailGoodsFindStoreByGoodsIdWithPageRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "retailGoods/findStoreByGoodsIdWithPage", "api")
	request.Method = api.POST
	return
}

func CreateRetailGoodsFindStoreByGoodsIdWithPageResponse() (response *api.BaseResponse[RetailGoodsFindStoreByGoodsIdWithPageResponse]) {
	response = api.CreateResponse[RetailGoodsFindStoreByGoodsIdWithPageResponse](&RetailGoodsFindStoreByGoodsIdWithPageResponse{})
	return
}
