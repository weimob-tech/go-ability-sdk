package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsFindDeliveryTypeList
// @id 354
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取配送列表)
func (client *Client) GoodsFindDeliveryTypeList(request *GoodsFindDeliveryTypeListRequest) (response *GoodsFindDeliveryTypeListResponse, err error) {
	rpcResponse := CreateGoodsFindDeliveryTypeListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsFindDeliveryTypeListRequest struct {
	*api.BaseRequest

	GoodsId int64 `json:"goodsId,omitempty"`
}

type GoodsFindDeliveryTypeListResponse struct {
}

func CreateGoodsFindDeliveryTypeListRequest() (request *GoodsFindDeliveryTypeListRequest) {
	request = &GoodsFindDeliveryTypeListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goods/findDeliveryTypeList", "api")
	request.Method = api.POST
	return
}

func CreateGoodsFindDeliveryTypeListResponse() (response *api.BaseResponse[GoodsFindDeliveryTypeListResponse]) {
	response = api.CreateResponse[GoodsFindDeliveryTypeListResponse](&GoodsFindDeliveryTypeListResponse{})
	return
}
