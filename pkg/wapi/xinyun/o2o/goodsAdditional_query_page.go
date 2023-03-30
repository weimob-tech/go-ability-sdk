package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsAdditionalQueryPage
// @id 150
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询加值)
func (client *Client) GoodsAdditionalQueryPage(request *GoodsAdditionalQueryPageRequest) (response *GoodsAdditionalQueryPageResponse, err error) {
	rpcResponse := CreateGoodsAdditionalQueryPageResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsAdditionalQueryPageRequest struct {
	*api.BaseRequest

	MerchantId int64 `json:"merchantId,omitempty"`
	StoreId    int64 `json:"storeId,omitempty"`
	PageNum    int   `json:"pageNum,omitempty"`
	PageSize   int   `json:"pageSize,omitempty"`
}

type GoodsAdditionalQueryPageResponse struct {
}

func CreateGoodsAdditionalQueryPageRequest() (request *GoodsAdditionalQueryPageRequest) {
	request = &GoodsAdditionalQueryPageRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goodsAdditional/queryPage", "api")
	request.Method = api.POST
	return
}

func CreateGoodsAdditionalQueryPageResponse() (response *api.BaseResponse[GoodsAdditionalQueryPageResponse]) {
	response = api.CreateResponse[GoodsAdditionalQueryPageResponse](&GoodsAdditionalQueryPageResponse{})
	return
}
