package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PriceGetPriceList
// @id 1021
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取价格日历)
func (client *Client) PriceGetPriceList(request *PriceGetPriceListRequest) (response *PriceGetPriceListResponse, err error) {
	rpcResponse := CreatePriceGetPriceListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PriceGetPriceListRequest struct {
	*api.BaseRequest

	Type      int    `json:"type,omitempty"`
	GoodsCode string `json:"goodsCode,omitempty"`
	StartTime int    `json:"startTime,omitempty"`
	EndTime   int    `json:"endTime,omitempty"`
}

type PriceGetPriceListResponse struct {
	PriceItem       PriceGetPriceListResponsePriceItem `json:"priceItem,omitempty"`
	PriceItemOutVO  []map[string]any                   `json:"PriceItemOutVO,omitempty"`
	GoodsCode       string                             `json:"goodsCode,omitempty"`
	OtherPrice      int64                              `json:"otherPrice,omitempty"`
	PublicStockCode string                             `json:"publicStockCode,omitempty"`
	Sales           int64                              `json:"sales,omitempty"`
	Stock           int64                              `json:"stock,omitempty"`
	LimitType       int64                              `json:"limitType,omitempty"`
	StartTime       string                             `json:"startTime,omitempty"`
	EndTime         string                             `json:"endTime,omitempty"`
	Remark          string                             `json:"remark,omitempty"`
	SkuCode         string                             `json:"skuCode,omitempty"`
	SkuName         string                             `json:"skuName,omitempty"`
	Code            string                             `json:"code,omitempty"`
	Price           string                             `json:"price,omitempty"`
	IsMain          int64                              `json:"isMain,omitempty"`
}

func CreatePriceGetPriceListRequest() (request *PriceGetPriceListRequest) {
	request = &PriceGetPriceListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "price/getPriceList", "api")
	request.Method = api.POST
	return
}

func CreatePriceGetPriceListResponse() (response *api.BaseResponse[PriceGetPriceListResponse]) {
	response = api.CreateResponse[PriceGetPriceListResponse](&PriceGetPriceListResponse{})
	return
}

type PriceGetPriceListResponsePriceItem struct {
}
