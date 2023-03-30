package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SingleProductUpdateSingleProductCoastPrice
// @id 545
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改成本价)
func (client *Client) SingleProductUpdateSingleProductCoastPrice(request *SingleProductUpdateSingleProductCoastPriceRequest) (response *SingleProductUpdateSingleProductCoastPriceResponse, err error) {
	rpcResponse := CreateSingleProductUpdateSingleProductCoastPriceResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SingleProductUpdateSingleProductCoastPriceRequest struct {
	*api.BaseRequest

	StoreId   int64   `json:"storeId,omitempty"`
	ProductId int64   `json:"productId,omitempty"`
	CostPrice float64 `json:"costPrice,omitempty"`
}

type SingleProductUpdateSingleProductCoastPriceResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateSingleProductUpdateSingleProductCoastPriceRequest() (request *SingleProductUpdateSingleProductCoastPriceRequest) {
	request = &SingleProductUpdateSingleProductCoastPriceRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "singleProduct/updateSingleProductCoastPrice", "api")
	request.Method = api.POST
	return
}

func CreateSingleProductUpdateSingleProductCoastPriceResponse() (response *api.BaseResponse[SingleProductUpdateSingleProductCoastPriceResponse]) {
	response = api.CreateResponse[SingleProductUpdateSingleProductCoastPriceResponse](&SingleProductUpdateSingleProductCoastPriceResponse{})
	return
}
