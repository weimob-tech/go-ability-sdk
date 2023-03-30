package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FulfillDeliverygoodsGetList
// @id 1904
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1904?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=批量查询商品履约关系)
func (client *Client) FulfillDeliverygoodsGetList(request *FulfillDeliverygoodsGetListRequest) (response *FulfillDeliverygoodsGetListResponse, err error) {
	rpcResponse := CreateFulfillDeliverygoodsGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FulfillDeliverygoodsGetListRequest struct {
	*api.BaseRequest

	GoodsNodeDtoList []FulfillDeliverygoodsGetListRequestGoodsNodeDtoList `json:"goodsNodeDtoList,omitempty"`
	BasicInfo        FulfillDeliverygoodsGetListRequestBasicInfo          `json:"basicInfo,omitempty"`
}

type FulfillDeliverygoodsGetListResponse struct {
	GoodsDeliveryShipVos []FulfillDeliverygoodsGetListResponseGoodsDeliveryShipVos `json:"goodsDeliveryShipVos,omitempty"`
}

func CreateFulfillDeliverygoodsGetListRequest() (request *FulfillDeliverygoodsGetListRequest) {
	request = &FulfillDeliverygoodsGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "fulfill/deliverygoods/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateFulfillDeliverygoodsGetListResponse() (response *api.BaseResponse[FulfillDeliverygoodsGetListResponse]) {
	response = api.CreateResponse[FulfillDeliverygoodsGetListResponse](&FulfillDeliverygoodsGetListResponse{})
	return
}

type FulfillDeliverygoodsGetListRequestGoodsNodeDtoList struct {
	Vid     int64 `json:"vid,omitempty"`
	GoodsId int64 `json:"goodsId,omitempty"`
}

type FulfillDeliverygoodsGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type FulfillDeliverygoodsGetListResponseGoodsDeliveryShipVos struct {
	Vid                int64 `json:"vid,omitempty"`
	DeliveryId         int64 `json:"deliveryId,omitempty"`
	DeliveryNodeShipId int64 `json:"deliveryNodeShipId,omitempty"`
	GoodsId            int64 `json:"goodsId,omitempty"`
	ProductInstanceId  int64 `json:"productInstanceId,omitempty"`
	DeliveryType       int64 `json:"deliveryType,omitempty"`
	TemplateId         int64 `json:"templateId,omitempty"`
}
