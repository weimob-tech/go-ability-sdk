package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsConfirmReceivedRightsGoods
// @id 382
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=退货退款/换货时商家确认收货)
func (client *Client) RightsConfirmReceivedRightsGoods(request *RightsConfirmReceivedRightsGoodsRequest) (response *RightsConfirmReceivedRightsGoodsResponse, err error) {
	rpcResponse := CreateRightsConfirmReceivedRightsGoodsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsConfirmReceivedRightsGoodsRequest struct {
	*api.BaseRequest

	Id int64 `json:"id,omitempty"`
}

type RightsConfirmReceivedRightsGoodsResponse struct {
}

func CreateRightsConfirmReceivedRightsGoodsRequest() (request *RightsConfirmReceivedRightsGoodsRequest) {
	request = &RightsConfirmReceivedRightsGoodsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "rights/confirmReceivedRightsGoods", "api")
	request.Method = api.POST
	return
}

func CreateRightsConfirmReceivedRightsGoodsResponse() (response *api.BaseResponse[RightsConfirmReceivedRightsGoodsResponse]) {
	response = api.CreateResponse[RightsConfirmReceivedRightsGoodsResponse](&RightsConfirmReceivedRightsGoodsResponse{})
	return
}
