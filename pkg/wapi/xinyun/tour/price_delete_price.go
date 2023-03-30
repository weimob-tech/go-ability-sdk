package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PriceDeletePrice
// @id 1020
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量删除价格日历)
func (client *Client) PriceDeletePrice(request *PriceDeletePriceRequest) (response *PriceDeletePriceResponse, err error) {
	rpcResponse := CreatePriceDeletePriceResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PriceDeletePriceRequest struct {
	*api.BaseRequest

	PriceItems         []map[string]any                          `json:"priceItems,omitempty"`
	PriceRequestVO     PriceDeletePriceRequestPriceRequestVO     `json:"PriceRequestVO,omitempty"`
	PriceDetails       []map[string]any                          `json:"priceDetails,omitempty"`
	PriceItemRequestVO PriceDeletePriceRequestPriceItemRequestVO `json:"PriceItemRequestVO,omitempty"`
	Type               int                                       `json:"type,omitempty"`
	GoodsCode          string                                    `json:"goodsCode,omitempty"`
	PublicStockCode    string                                    `json:"publicStockCode,omitempty"`
	StartTime          int                                       `json:"startTime,omitempty"`
	EndTime            int                                       `json:"endTime,omitempty"`
	Code               string                                    `json:"code,omitempty"`
	SkuCode            string                                    `json:"skuCode,omitempty"`
	IsMain             int                                       `json:"isMain,omitempty"`
}

type PriceDeletePriceResponse struct {
}

func CreatePriceDeletePriceRequest() (request *PriceDeletePriceRequest) {
	request = &PriceDeletePriceRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "price/deletePrice", "api")
	request.Method = api.POST
	return
}

func CreatePriceDeletePriceResponse() (response *api.BaseResponse[PriceDeletePriceResponse]) {
	response = api.CreateResponse[PriceDeletePriceResponse](&PriceDeletePriceResponse{})
	return
}

type PriceDeletePriceRequestPriceRequestVO struct {
}

type PriceDeletePriceRequestPriceItemRequestVO struct {
}
