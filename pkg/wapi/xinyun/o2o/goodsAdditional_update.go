package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsAdditionalUpdate
// @id 149
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新加值)
func (client *Client) GoodsAdditionalUpdate(request *GoodsAdditionalUpdateRequest) (response *GoodsAdditionalUpdateResponse, err error) {
	rpcResponse := CreateGoodsAdditionalUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsAdditionalUpdateRequest struct {
	*api.BaseRequest

	MerchantId        int64   `json:"merchantId,omitempty"`
	StoreId           int64   `json:"storeId,omitempty"`
	AdditionalName    string  `json:"additionalName,omitempty"`
	Price             float64 `json:"price,omitempty"`
	StockNum          int64   `json:"stockNum,omitempty"`
	ThirdGoodsId      string  `json:"thirdGoodsId,omitempty"`
	GoodsAdditionalId int64   `json:"goodsAdditionalId,omitempty"`
}

type GoodsAdditionalUpdateResponse struct {
}

func CreateGoodsAdditionalUpdateRequest() (request *GoodsAdditionalUpdateRequest) {
	request = &GoodsAdditionalUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goodsAdditional/update", "api")
	request.Method = api.POST
	return
}

func CreateGoodsAdditionalUpdateResponse() (response *api.BaseResponse[GoodsAdditionalUpdateResponse]) {
	response = api.CreateResponse[GoodsAdditionalUpdateResponse](&GoodsAdditionalUpdateResponse{})
	return
}
