package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderNiche
// @id 2663
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取商机下订单列表)
func (client *Client) OrderNiche(request *OrderNicheRequest) (response *OrderNicheResponse, err error) {
	rpcResponse := CreateOrderNicheResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderNicheRequest struct {
	*api.BaseRequest

	UserWid  int64  `json:"userWid,omitempty"`
	Key      string `json:"key,omitempty"`
	PageSize int64  `json:"pageSize,omitempty"`
	PageNum  int64  `json:"pageNum,omitempty"`
}

type OrderNicheResponse struct {
}

func CreateOrderNicheRequest() (request *OrderNicheRequest) {
	request = &OrderNicheRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "order/niche", "api")
	request.Method = api.POST
	return
}

func CreateOrderNicheResponse() (response *api.BaseResponse[OrderNicheResponse]) {
	response = api.CreateResponse[OrderNicheResponse](&OrderNicheResponse{})
	return
}
