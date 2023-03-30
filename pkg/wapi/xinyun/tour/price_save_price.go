package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PriceSavePrice
// @id 1018
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增/修改价格日历)
func (client *Client) PriceSavePrice(request *PriceSavePriceRequest) (response *PriceSavePriceResponse, err error) {
	rpcResponse := CreatePriceSavePriceResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PriceSavePriceRequest struct {
	*api.BaseRequest

	PriceItem       []map[string]any                  `json:"priceItem,omitempty"`
	PriceCalendar   []map[string]any                  `json:"priceCalendar,omitempty"`
	PriceItemVO     PriceSavePriceRequestPriceItemVO  `json:"PriceItemVO,omitempty"`
	CalendarInVO    PriceSavePriceRequestCalendarInVO `json:"CalendarInVO,omitempty"`
	Codes           []map[string]any                  `json:"codes,omitempty"`
	GoodsCode       string                            `json:"goodsCode,omitempty"`
	Type            int                               `json:"type,omitempty"`
	IsPublic        int                               `json:"isPublic,omitempty"`
	PublicStockCode string                            `json:"publicStockCode,omitempty"`
	Price           float64                           `json:"price,omitempty"`
	Stock           int                               `json:"stock,omitempty"`
	Remark          string                            `json:"remark,omitempty"`
	StartTime       int                               `json:"startTime,omitempty"`
	EndTime         int                               `json:"endTime,omitempty"`
}

type PriceSavePriceResponse struct {
}

func CreatePriceSavePriceRequest() (request *PriceSavePriceRequest) {
	request = &PriceSavePriceRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "price/savePrice", "api")
	request.Method = api.POST
	return
}

func CreatePriceSavePriceResponse() (response *api.BaseResponse[PriceSavePriceResponse]) {
	response = api.CreateResponse[PriceSavePriceResponse](&PriceSavePriceResponse{})
	return
}

type PriceSavePriceRequestPriceItemVO struct {
	SkuCode string `json:"skuCode,omitempty"`
	Price   int64  `json:"price,omitempty"`
	IsMain  int64  `json:"isMain,omitempty"`
	SkuName string `json:"skuName,omitempty"`
}

type PriceSavePriceRequestCalendarInVO struct {
}
