package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SpuShopGet
// @id 289
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取门店商品信息)
func (client *Client) SpuShopGet(request *SpuShopGetRequest) (response *SpuShopGetResponse, err error) {
	rpcResponse := CreateSpuShopGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SpuShopGetRequest struct {
	*api.BaseRequest

	CategoryId string `json:"category_id,omitempty"`
	ClassifyId string `json:"classify_id,omitempty"`
	IsOnsale   int    `json:"is_onsale,omitempty"`
	PageNo     int    `json:"page_no,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
	SpuCode    string `json:"spu_code,omitempty"`
	ShopId     int64  `json:"shop_id,omitempty"`
}

type SpuShopGetResponse struct {
}

func CreateSpuShopGetRequest() (request *SpuShopGetRequest) {
	request = &SpuShopGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "spu/ShopGet", "api")
	request.Method = api.POST
	return
}

func CreateSpuShopGetResponse() (response *api.BaseResponse[SpuShopGetResponse]) {
	response = api.CreateResponse[SpuShopGetResponse](&SpuShopGetResponse{})
	return
}
