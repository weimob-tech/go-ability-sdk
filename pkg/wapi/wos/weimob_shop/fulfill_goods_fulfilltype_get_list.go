package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FulfillGoodsFulfilltypeGetList
// @id 1431
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1431?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询配送方式列表)
func (client *Client) FulfillGoodsFulfilltypeGetList(request *FulfillGoodsFulfilltypeGetListRequest) (response *FulfillGoodsFulfilltypeGetListResponse, err error) {
	rpcResponse := CreateFulfillGoodsFulfilltypeGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FulfillGoodsFulfilltypeGetListRequest struct {
	*api.BaseRequest

	BasicInfo FulfillGoodsFulfilltypeGetListRequestBasicInfo `json:"basicInfo,omitempty"`
}

type FulfillGoodsFulfilltypeGetListResponse struct {
	NodeDeliveryDtoList []FulfillGoodsFulfilltypeGetListResponseNodeDeliveryDtoList `json:"nodeDeliveryDtoList,omitempty"`
}

func CreateFulfillGoodsFulfilltypeGetListRequest() (request *FulfillGoodsFulfilltypeGetListRequest) {
	request = &FulfillGoodsFulfilltypeGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "fulfill/goods/fulfilltype/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateFulfillGoodsFulfilltypeGetListResponse() (response *api.BaseResponse[FulfillGoodsFulfilltypeGetListResponse]) {
	response = api.CreateResponse[FulfillGoodsFulfilltypeGetListResponse](&FulfillGoodsFulfilltypeGetListResponse{})
	return
}

type FulfillGoodsFulfilltypeGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type FulfillGoodsFulfilltypeGetListResponseNodeDeliveryDtoList struct {
	Id               int64  `json:"id,omitempty"`
	DeliveryId       int64  `json:"deliveryId,omitempty"`
	DeliveryType     int64  `json:"deliveryType,omitempty"`
	DeliveryTypeName string `json:"deliveryTypeName,omitempty"`
	IsDefault        int64  `json:"isDefault,omitempty"`
	IsSupported      int64  `json:"isSupported,omitempty"`
}
