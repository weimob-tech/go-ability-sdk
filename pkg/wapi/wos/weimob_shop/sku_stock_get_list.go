package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SkuStockGetList
// @id 1860
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1860?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=批量查询商品库存)
func (client *Client) SkuStockGetList(request *SkuStockGetListRequest) (response *SkuStockGetListResponse, err error) {
	rpcResponse := CreateSkuStockGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SkuStockGetListRequest struct {
	*api.BaseRequest

	BasicInfo SkuStockGetListRequestBasicInfo `json:"basicInfo,omitempty"`
	SkuIdList []int64                         `json:"skuIdList,omitempty"`
}

type SkuStockGetListResponse struct {
	SkuStoreStockVoList []SkuStockGetListResponseSkuStoreStockVoList `json:"skuStoreStockVoList,omitempty"`
}

func CreateSkuStockGetListRequest() (request *SkuStockGetListRequest) {
	request = &SkuStockGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "sku/stock/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateSkuStockGetListResponse() (response *api.BaseResponse[SkuStockGetListResponse]) {
	response = api.CreateResponse[SkuStockGetListResponse](&SkuStockGetListResponse{})
	return
}

type SkuStockGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type SkuStockGetListResponseSkuStoreStockVoList struct {
	GoodsId           int64 `json:"goodsId,omitempty"`
	Vid               int64 `json:"vid,omitempty"`
	RealStockNum      int64 `json:"realStockNum,omitempty"`
	FrozenStockNum    int64 `json:"frozenStockNum,omitempty"`
	AvailableStockNum int64 `json:"availableStockNum,omitempty"`
	SkuId             int64 `json:"skuId,omitempty"`
}
