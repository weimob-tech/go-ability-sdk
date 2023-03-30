package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantBatchUpdateStockConfig
// @id 1142
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量变更仓库模式)
func (client *Client) MerchantBatchUpdateStockConfig(request *MerchantBatchUpdateStockConfigRequest) (response *MerchantBatchUpdateStockConfigResponse, err error) {
	rpcResponse := CreateMerchantBatchUpdateStockConfigResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantBatchUpdateStockConfigRequest struct {
	*api.BaseRequest

	StoreIdList  []int64 `json:"storeIdList,omitempty"`
	StockUseType int     `json:"stockUseType,omitempty"`
	StockId      int     `json:"stockId,omitempty"`
}

type MerchantBatchUpdateStockConfigResponse struct {
}

func CreateMerchantBatchUpdateStockConfigRequest() (request *MerchantBatchUpdateStockConfigRequest) {
	request = &MerchantBatchUpdateStockConfigRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchant/batchUpdateStockConfig", "api")
	request.Method = api.POST
	return
}

func CreateMerchantBatchUpdateStockConfigResponse() (response *api.BaseResponse[MerchantBatchUpdateStockConfigResponse]) {
	response = api.CreateResponse[MerchantBatchUpdateStockConfigResponse](&MerchantBatchUpdateStockConfigResponse{})
	return
}
