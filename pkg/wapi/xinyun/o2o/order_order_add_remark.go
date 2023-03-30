package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderOrderAddRemark
// @id 156
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=订单商户添加备注)
func (client *Client) OrderOrderAddRemark(request *OrderOrderAddRemarkRequest) (response *OrderOrderAddRemarkResponse, err error) {
	rpcResponse := CreateOrderOrderAddRemarkResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderOrderAddRemarkRequest struct {
	*api.BaseRequest

	OrderId int64  `json:"orderId,omitempty"`
	OrderNo string `json:"orderNo,omitempty"`
	Remark  string `json:"remark,omitempty"`
}

type OrderOrderAddRemarkResponse struct {
}

func CreateOrderOrderAddRemarkRequest() (request *OrderOrderAddRemarkRequest) {
	request = &OrderOrderAddRemarkRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "order/orderAddRemark", "api")
	request.Method = api.POST
	return
}

func CreateOrderOrderAddRemarkResponse() (response *api.BaseResponse[OrderOrderAddRemarkResponse]) {
	response = api.CreateResponse[OrderOrderAddRemarkResponse](&OrderOrderAddRemarkResponse{})
	return
}
