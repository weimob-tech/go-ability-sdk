package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsList
// @id 61
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询菜品列表)
func (client *Client) GoodsList(request *GoodsListRequest) (response *GoodsListResponse, err error) {
	rpcResponse := CreateGoodsListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsListRequest struct {
	*api.BaseRequest

	StoreId    int64 `json:"storeId,omitempty"`
	PageNum    int   `json:"pageNum,omitempty"`
	PageSize   int   `json:"pageSize,omitempty"`
	MerchantId int64 `json:"merchantId,omitempty"`
	Status     int   `json:"status,omitempty"`
}

type GoodsListResponse struct {
}

func CreateGoodsListRequest() (request *GoodsListRequest) {
	request = &GoodsListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goods/list", "api")
	request.Method = api.POST
	return
}

func CreateGoodsListResponse() (response *api.BaseResponse[GoodsListResponse]) {
	response = api.CreateResponse[GoodsListResponse](&GoodsListResponse{})
	return
}
