package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// LogisticsDelivery
// @id 23
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=发货接口)
func (client *Client) LogisticsDelivery(request *LogisticsDeliveryRequest) (response *LogisticsDeliveryResponse, err error) {
	rpcResponse := CreateLogisticsDeliveryResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type LogisticsDeliveryRequest struct {
	*api.BaseRequest

	Deliveries         []map[string]any `json:"deliveries,omitempty"`
	OrderNo            string           `json:"order_no,omitempty"`
	NeedDelivery       bool             `json:"need_delivery,omitempty"`
	NeedCustomDelivery bool             `json:"need_custom_delivery,omitempty"`
	CarrierCode        string           `json:"carrier_code,omitempty"`
	CarrierName        string           `json:"carrier_name,omitempty"`
	ExpressNo          string           `json:"express_no,omitempty"`
	Remark             string           `json:"remark,omitempty"`
	SenderAddress      string           `json:"sender_address,omitempty"`
	SenderName         string           `json:"sender_name,omitempty"`
	SenderTel          string           `json:"sender_tel,omitempty"`
}

type LogisticsDeliveryResponse struct {
}

func CreateLogisticsDeliveryRequest() (request *LogisticsDeliveryRequest) {
	request = &LogisticsDeliveryRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "Logistics/Delivery", "api")
	request.Method = api.POST
	return
}

func CreateLogisticsDeliveryResponse() (response *api.BaseResponse[LogisticsDeliveryResponse]) {
	response = api.CreateResponse[LogisticsDeliveryResponse](&LogisticsDeliveryResponse{})
	return
}
