package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsUpdateGoodsShelfStatus
// @id 348
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量上下架)
func (client *Client) GoodsUpdateGoodsShelfStatus(request *GoodsUpdateGoodsShelfStatusRequest) (response *GoodsUpdateGoodsShelfStatusResponse, err error) {
	rpcResponse := CreateGoodsUpdateGoodsShelfStatusResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsUpdateGoodsShelfStatusRequest struct {
	*api.BaseRequest

	GoodsIdList []int64 `json:"goodsIdList,omitempty"`
	IsPutAway   int     `json:"isPutAway,omitempty"`
	StoreId     int64   `json:"storeId,omitempty"`
}

type GoodsUpdateGoodsShelfStatusResponse struct {
}

func CreateGoodsUpdateGoodsShelfStatusRequest() (request *GoodsUpdateGoodsShelfStatusRequest) {
	request = &GoodsUpdateGoodsShelfStatusRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goods/updateGoodsShelfStatus", "api")
	request.Method = api.POST
	return
}

func CreateGoodsUpdateGoodsShelfStatusResponse() (response *api.BaseResponse[GoodsUpdateGoodsShelfStatusResponse]) {
	response = api.CreateResponse[GoodsUpdateGoodsShelfStatusResponse](&GoodsUpdateGoodsShelfStatusResponse{})
	return
}
