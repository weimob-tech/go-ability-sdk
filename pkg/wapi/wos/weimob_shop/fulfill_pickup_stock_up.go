package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FulfillPickupStockUp
// @id 1734
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1734?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=自提订单备货)
func (client *Client) FulfillPickupStockUp(request *FulfillPickupStockUpRequest) (response *FulfillPickupStockUpResponse, err error) {
	rpcResponse := CreateFulfillPickupStockUpResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FulfillPickupStockUpRequest struct {
	*api.BaseRequest

	OperatorVo     FulfillPickupStockUpRequestOperatorVo     `json:"operatorVo,omitempty"`
	FulfillItems   []FulfillPickupStockUpRequestFulfillItems `json:"fulfillItems,omitempty"`
	BasicInfo      FulfillPickupStockUpRequestBasicInfo      `json:"basicInfo,omitempty"`
	OrderNo        int64                                     `json:"orderNo,omitempty"`
	IsSplitPackage bool                                      `json:"isSplitPackage,omitempty"`
	FulfillMethod  int64                                     `json:"fulfillMethod,omitempty"`
}

type FulfillPickupStockUpResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateFulfillPickupStockUpRequest() (request *FulfillPickupStockUpRequest) {
	request = &FulfillPickupStockUpRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "fulfill/pickup/stockUp", "apigw")
	request.Method = api.POST
	return
}

func CreateFulfillPickupStockUpResponse() (response *api.BaseResponse[FulfillPickupStockUpResponse]) {
	response = api.CreateResponse[FulfillPickupStockUpResponse](&FulfillPickupStockUpResponse{})
	return
}

type FulfillPickupStockUpRequestOperatorVo struct {
	OperatorId    string `json:"operatorId,omitempty"`
	OperatorName  string `json:"operatorName,omitempty"`
	OperatorPhone string `json:"operatorPhone,omitempty"`
}

type FulfillPickupStockUpRequestFulfillItems struct {
	OrderItemId int64 `json:"orderItemId,omitempty"`
	SkuNum      int64 `json:"skuNum,omitempty"`
}

type FulfillPickupStockUpRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}
