package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CanteenModifyOrderDetailList
// @id 1411
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=点餐订单改单)
func (client *Client) CanteenModifyOrderDetailList(request *CanteenModifyOrderDetailListRequest) (response *CanteenModifyOrderDetailListResponse, err error) {
	rpcResponse := CreateCanteenModifyOrderDetailListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CanteenModifyOrderDetailListRequest struct {
	*api.BaseRequest

	OrderDetailList []CanteenModifyOrderDetailListRequestOrderDetailList `json:"orderDetailList,omitempty"`
	ConsumerWid     int64                                                `json:"consumerWid,omitempty"`
	LastestTime     int64                                                `json:"lastestTime,omitempty"`
	Operator        string                                               `json:"operator,omitempty"`
	OrderNo         string                                               `json:"orderNo,omitempty"`
	PayAmount       float64                                              `json:"payAmount,omitempty"`
	StoreId         int64                                                `json:"storeId,omitempty"`
	TotalAmount     float64                                              `json:"totalAmount,omitempty"`
}

type CanteenModifyOrderDetailListResponse struct {
}

func CreateCanteenModifyOrderDetailListRequest() (request *CanteenModifyOrderDetailListRequest) {
	request = &CanteenModifyOrderDetailListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "canteen/modifyOrderDetailList", "api")
	request.Method = api.POST
	return
}

func CreateCanteenModifyOrderDetailListResponse() (response *api.BaseResponse[CanteenModifyOrderDetailListResponse]) {
	response = api.CreateResponse[CanteenModifyOrderDetailListResponse](&CanteenModifyOrderDetailListResponse{})
	return
}

type CanteenModifyOrderDetailListRequestOrderDetailList struct {
	AdditionList   []CanteenModifyOrderDetailListRequestAdditionList   `json:"additionList,omitempty"`
	GoodsComboList []CanteenModifyOrderDetailListRequestGoodsComboList `json:"goodsComboList,omitempty"`
	SkuInfo        CanteenModifyOrderDetailListRequestSkuInfo          `json:"skuInfo,omitempty"`
	TasteList      []CanteenModifyOrderDetailListRequestTasteList      `json:"tasteList,omitempty"`
	GoodsId        int64                                               `json:"goodsId,omitempty"`
	GoodsName      string                                              `json:"goodsName,omitempty"`
	GoodsNum       int                                                 `json:"goodsNum,omitempty"`
	GoodsPrice     float64                                             `json:"goodsPrice,omitempty"`
	GoodsType      int                                                 `json:"goodsType,omitempty"`
	ThirdGoodsId   string                                              `json:"thirdGoodsId,omitempty"`
}

type CanteenModifyOrderDetailListRequestAdditionList struct {
	AdditionId      int64   `json:"additionId,omitempty"`
	AdditionName    string  `json:"additionName,omitempty"`
	AdditionNum     int     `json:"additionNum,omitempty"`
	AdditionPrice   float64 `json:"additionPrice,omitempty"`
	AdditionThirdId string  `json:"additionThirdId,omitempty"`
}

type CanteenModifyOrderDetailListRequestGoodsComboList struct {
	GoodsComboId    int64   `json:"goodsComboId,omitempty"`
	GoodsComboName  string  `json:"goodsComboName,omitempty"`
	GoodsComboPrice float64 `json:"goodsComboPrice,omitempty"`
	ThirdGoodsId    string  `json:"thirdGoodsId,omitempty"`
}

type CanteenModifyOrderDetailListRequestSkuInfo struct {
	SkuId      int64   `json:"skuId,omitempty"`
	SkuName    string  `json:"skuName,omitempty"`
	SkuPrice   float64 `json:"skuPrice,omitempty"`
	ThirdSkuId string  `json:"thirdSkuId,omitempty"`
}

type CanteenModifyOrderDetailListRequestTasteList struct {
	TasteId   int64  `json:"tasteId,omitempty"`
	TasteName string `json:"tasteName,omitempty"`
}
