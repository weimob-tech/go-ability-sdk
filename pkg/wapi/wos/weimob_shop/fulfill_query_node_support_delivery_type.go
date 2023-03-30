package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FulfillQueryNodeSupportDeliveryType
// @id 1236
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1236?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取配送方式列表)
func (client *Client) FulfillQueryNodeSupportDeliveryType(request *FulfillQueryNodeSupportDeliveryTypeRequest) (response *FulfillQueryNodeSupportDeliveryTypeResponse, err error) {
	rpcResponse := CreateFulfillQueryNodeSupportDeliveryTypeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FulfillQueryNodeSupportDeliveryTypeRequest struct {
	*api.BaseRequest

	DeliveryTypeList []int64 `json:"deliveryTypeList,omitempty"`
	BizStoreVid      int64   `json:"bizStoreVid,omitempty"`
}

type FulfillQueryNodeSupportDeliveryTypeResponse struct {
	NodeDeliveryDtoList []FulfillQueryNodeSupportDeliveryTypeResponseNodeDeliveryDtoList `json:"nodeDeliveryDtoList,omitempty"`
}

func CreateFulfillQueryNodeSupportDeliveryTypeRequest() (request *FulfillQueryNodeSupportDeliveryTypeRequest) {
	request = &FulfillQueryNodeSupportDeliveryTypeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "fulfill/queryNodeSupportDeliveryType", "apigw")
	request.Method = api.POST
	return
}

func CreateFulfillQueryNodeSupportDeliveryTypeResponse() (response *api.BaseResponse[FulfillQueryNodeSupportDeliveryTypeResponse]) {
	response = api.CreateResponse[FulfillQueryNodeSupportDeliveryTypeResponse](&FulfillQueryNodeSupportDeliveryTypeResponse{})
	return
}

type FulfillQueryNodeSupportDeliveryTypeResponseNodeDeliveryDtoList struct {
	DeliveryId       int64  `json:"deliveryId,omitempty"`
	DeliveryTypeName string `json:"deliveryTypeName,omitempty"`
	IsDefault        int64  `json:"isDefault,omitempty"`
	DeliveryType     int64  `json:"deliveryType,omitempty"`
	Id               int64  `json:"id,omitempty"`
	IsSupported      int64  `json:"isSupported,omitempty"`
}
