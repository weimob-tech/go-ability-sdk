package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StockQueryAllocateOutStorageOrderListWithPage
// @id 557
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询调拨出库单)
func (client *Client) StockQueryAllocateOutStorageOrderListWithPage(request *StockQueryAllocateOutStorageOrderListWithPageRequest) (response *StockQueryAllocateOutStorageOrderListWithPageResponse, err error) {
	rpcResponse := CreateStockQueryAllocateOutStorageOrderListWithPageResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StockQueryAllocateOutStorageOrderListWithPageRequest struct {
	*api.BaseRequest

	QueryParameter StockQueryAllocateOutStorageOrderListWithPageRequestQueryParameter `json:"queryParameter,omitempty"`
	PageNum        int                                                                `json:"pageNum,omitempty"`
	PageSize       int                                                                `json:"pageSize,omitempty"`
	SearchStoreId  string                                                             `json:"searchStoreId,omitempty"`
	SearchType     int                                                                `json:"searchType,omitempty"`
	Search         string                                                             `json:"search,omitempty"`
	StoreId        string                                                             `json:"storeId,omitempty"`
}

type StockQueryAllocateOutStorageOrderListWithPageResponse struct {
}

func CreateStockQueryAllocateOutStorageOrderListWithPageRequest() (request *StockQueryAllocateOutStorageOrderListWithPageRequest) {
	request = &StockQueryAllocateOutStorageOrderListWithPageRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "stock/queryAllocateOutStorageOrderListWithPage", "api")
	request.Method = api.POST
	return
}

func CreateStockQueryAllocateOutStorageOrderListWithPageResponse() (response *api.BaseResponse[StockQueryAllocateOutStorageOrderListWithPageResponse]) {
	response = api.CreateResponse[StockQueryAllocateOutStorageOrderListWithPageResponse](&StockQueryAllocateOutStorageOrderListWithPageResponse{})
	return
}

type StockQueryAllocateOutStorageOrderListWithPageRequestQueryParameter struct {
}
