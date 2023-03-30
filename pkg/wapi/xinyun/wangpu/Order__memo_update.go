package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderMemoUpdate
// @id 33
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改订单旗帜和备注)
func (client *Client) OrderMemoUpdate(request *OrderMemoUpdateRequest) (response *OrderMemoUpdateResponse, err error) {
	rpcResponse := CreateOrderMemoUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderMemoUpdateRequest struct {
	*api.BaseRequest

	OrderNo       string `json:"order_no,omitempty"`
	ShopId        string `json:"shop_id,omitempty"`
	UpdateMan     string `json:"update_man,omitempty"`
	OrderFlag     int    `json:"order_flag,omitempty"`
	OrderFlagMemo string `json:"order_flag_memo,omitempty"`
	Remark        string `json:"remark,omitempty"`
}

type OrderMemoUpdateResponse struct {
}

func CreateOrderMemoUpdateRequest() (request *OrderMemoUpdateRequest) {
	request = &OrderMemoUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "Order/MemoUpdate", "api")
	request.Method = api.POST
	return
}

func CreateOrderMemoUpdateResponse() (response *api.BaseResponse[OrderMemoUpdateResponse]) {
	response = api.CreateResponse[OrderMemoUpdateResponse](&OrderMemoUpdateResponse{})
	return
}
