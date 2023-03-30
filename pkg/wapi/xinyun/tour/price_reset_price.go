package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PriceResetPrice
// @id 1019
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=重置价格日历)
func (client *Client) PriceResetPrice(request *PriceResetPriceRequest) (response *PriceResetPriceResponse, err error) {
	rpcResponse := CreatePriceResetPriceResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PriceResetPriceRequest struct {
	*api.BaseRequest

	GoodsCode string `json:"goodsCode,omitempty"`
	Type      int    `json:"type,omitempty"`
}

type PriceResetPriceResponse struct {
}

func CreatePriceResetPriceRequest() (request *PriceResetPriceRequest) {
	request = &PriceResetPriceRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "price/resetPrice", "api")
	request.Method = api.POST
	return
}

func CreatePriceResetPriceResponse() (response *api.BaseResponse[PriceResetPriceResponse]) {
	response = api.CreateResponse[PriceResetPriceResponse](&PriceResetPriceResponse{})
	return
}
