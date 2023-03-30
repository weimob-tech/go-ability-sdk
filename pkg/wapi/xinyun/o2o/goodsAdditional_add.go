package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsAdditionalAdd
// @id 148
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=添加加值)
func (client *Client) GoodsAdditionalAdd(request *GoodsAdditionalAddRequest) (response *GoodsAdditionalAddResponse, err error) {
	rpcResponse := CreateGoodsAdditionalAddResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsAdditionalAddRequest struct {
	*api.BaseRequest

	MerchantId     int64   `json:"merchantId,omitempty"`
	StoreId        int64   `json:"storeId,omitempty"`
	AdditionalName string  `json:"additionalName,omitempty"`
	Price          float64 `json:"price,omitempty"`
	StockNum       int64   `json:"stockNum,omitempty"`
	ThirdGoodsId   string  `json:"thirdGoodsId,omitempty"`
}

type GoodsAdditionalAddResponse struct {
}

func CreateGoodsAdditionalAddRequest() (request *GoodsAdditionalAddRequest) {
	request = &GoodsAdditionalAddRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goodsAdditional/add", "api")
	request.Method = api.POST
	return
}

func CreateGoodsAdditionalAddResponse() (response *api.BaseResponse[GoodsAdditionalAddResponse]) {
	response = api.CreateResponse[GoodsAdditionalAddResponse](&GoodsAdditionalAddResponse{})
	return
}
