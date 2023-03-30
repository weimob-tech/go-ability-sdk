package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsGetGoodsList
// @id 502
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=商品列表)
func (client *Client) GoodsGetGoodsList(request *GoodsGetGoodsListRequest) (response *GoodsGetGoodsListResponse, err error) {
	rpcResponse := CreateGoodsGetGoodsListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsGetGoodsListRequest struct {
	*api.BaseRequest

	Name            string `json:"name,omitempty"`
	PageIndex       int    `json:"pageIndex,omitempty"`
	PageSize        int    `json:"pageSize,omitempty"`
	GoodsType       []int  `json:"goodsType,omitempty"`
	IsSellOut       bool   `json:"isSellOut,omitempty"`
	IsCheckSellTime bool   `json:"isCheckSellTime,omitempty"`
	IsNeedStoreId   bool   `json:"isNeedStoreId,omitempty"`
	IsInventory     bool   `json:"isInventory,omitempty"`
	IsPutAway       bool   `json:"isPutAway,omitempty"`
	OrderBy         int    `json:"orderBy,omitempty"`
	AscOrDesc       int    `json:"ascOrDesc,omitempty"`
}

type GoodsGetGoodsListResponse struct {
}

func CreateGoodsGetGoodsListRequest() (request *GoodsGetGoodsListRequest) {
	request = &GoodsGetGoodsListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "goods/getGoodsList", "api")
	request.Method = api.POST
	return
}

func CreateGoodsGetGoodsListResponse() (response *api.BaseResponse[GoodsGetGoodsListResponse]) {
	response = api.CreateResponse[GoodsGetGoodsListResponse](&GoodsGetGoodsListResponse{})
	return
}
