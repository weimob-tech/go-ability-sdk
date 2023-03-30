package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ComboUpdate
// @id 259
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新菜品套餐)
func (client *Client) ComboUpdate(request *ComboUpdateRequest) (response *ComboUpdateResponse, err error) {
	rpcResponse := CreateComboUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ComboUpdateRequest struct {
	*api.BaseRequest

	GoodsGroupId       []map[string]any               `json:"goodsGroupId,omitempty"`
	GoodsInfos         []ComboUpdateRequestGoodsInfos `json:"goodsInfos,omitempty"`
	Groups             []ComboUpdateRequestGroups     `json:"groups,omitempty"`
	SellTimes          []ComboUpdateRequestSellTimes  `json:"sellTimes,omitempty"`
	BowlFee            float64                        `json:"bowlFee,omitempty"`
	ComboId            int64                          `json:"comboId,omitempty"`
	ComboName          string                         `json:"comboName,omitempty"`
	ComboType          int                            `json:"comboType,omitempty"`
	DisplayPrice       float64                        `json:"displayPrice,omitempty"`
	GoodsTagId         int64                          `json:"goodsTagId,omitempty"`
	IncreaseUnit       int                            `json:"increaseUnit,omitempty"`
	OffShelfTime       int                            `json:"offShelfTime,omitempty"`
	OnShelfTime        int                            `json:"onShelfTime,omitempty"`
	OrderInPrice       float64                        `json:"orderInPrice,omitempty"`
	OrderInTotalSales  int64                          `json:"orderInTotalSales,omitempty"`
	OrderOutPrice      float64                        `json:"orderOutPrice,omitempty"`
	OrderOutTotalSales int64                          `json:"orderOutTotalSales,omitempty"`
	PicUrl             string                         `json:"picUrl,omitempty"`
	RecmdPriceType     int                            `json:"recmdPriceType,omitempty"`
	Sales              int64                          `json:"sales,omitempty"`
	Status             int                            `json:"status,omitempty"`
	StockNum           int64                          `json:"stockNum,omitempty"`
	StoreId            int64                          `json:"storeId,omitempty"`
	SuppOrderTypes     int                            `json:"suppOrderTypes,omitempty"`
	ThirdGoodsId       string                         `json:"thirdGoodsId,omitempty"`
}

type ComboUpdateResponse struct {
}

func CreateComboUpdateRequest() (request *ComboUpdateRequest) {
	request = &ComboUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "combo/update", "api")
	request.Method = api.POST
	return
}

func CreateComboUpdateResponse() (response *api.BaseResponse[ComboUpdateResponse]) {
	response = api.CreateResponse[ComboUpdateResponse](&ComboUpdateResponse{})
	return
}

type ComboUpdateRequestGoodsInfos struct {
	Comment  string `json:"comment,omitempty"`
	PhotoUrl string `json:"photoUrl,omitempty"`
}

type ComboUpdateRequestGroups struct {
	Items          []ComboUpdateRequestItems `json:"items,omitempty"`
	AvailableCount int                       `json:"availableCount,omitempty"`
	GoodsComboId   int64                     `json:"goodsComboId,omitempty"`
	GoodsOrderType int                       `json:"goodsOrderType,omitempty"`
	GroupName      string                    `json:"groupName,omitempty"`
	Sequence       int                       `json:"sequence,omitempty"`
}

type ComboUpdateRequestItems struct {
	Count             int     `json:"count,omitempty"`
	GoodsComboGroupId int64   `json:"goodsComboGroupId,omitempty"`
	GoodsComboId      int64   `json:"goodsComboId,omitempty"`
	GoodsId           int64   `json:"goodsId,omitempty"`
	GoodsName         string  `json:"goodsName,omitempty"`
	GoodsSkuIds       string  `json:"goodsSkuIds,omitempty"`
	Id                int64   `json:"id,omitempty"`
	IsForce           int     `json:"isForce,omitempty"`
	Price             float64 `json:"price,omitempty"`
	Sequence          int     `json:"sequence,omitempty"`
}

type ComboUpdateRequestSellTimes struct {
	TimeList []ComboUpdateRequestTimeList `json:"timeList,omitempty"`
	WeekDay  int                          `json:"weekDay,omitempty"`
}

type ComboUpdateRequestTimeList struct {
	EndTime   string `json:"endTime,omitempty"`
	StartTime string `json:"startTime,omitempty"`
}
