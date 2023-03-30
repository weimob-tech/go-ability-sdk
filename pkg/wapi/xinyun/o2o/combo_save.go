package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ComboSave
// @id 255
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新建菜品套餐)
func (client *Client) ComboSave(request *ComboSaveRequest) (response *ComboSaveResponse, err error) {
	rpcResponse := CreateComboSaveResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ComboSaveRequest struct {
	*api.BaseRequest

	GoodsGroupId       []map[string]any             `json:"goodsGroupId,omitempty"`
	GoodsInfos         []ComboSaveRequestGoodsInfos `json:"goodsInfos,omitempty"`
	Groups             []ComboSaveRequestGroups     `json:"groups,omitempty"`
	SellTimes          []ComboSaveRequestSellTimes  `json:"sellTimes,omitempty"`
	BowlFee            float64                      `json:"bowlFee,omitempty"`
	ComboId            int64                        `json:"comboId,omitempty"`
	ComboName          string                       `json:"comboName,omitempty"`
	ComboType          int                          `json:"comboType,omitempty"`
	DisplayPrice       float64                      `json:"displayPrice,omitempty"`
	GoodsTagId         int64                        `json:"goodsTagId,omitempty"`
	IncreaseUnit       int                          `json:"increaseUnit,omitempty"`
	OffShelfTime       int                          `json:"offShelfTime,omitempty"`
	OnShelfTime        int                          `json:"onShelfTime,omitempty"`
	OrderInPrice       float64                      `json:"orderInPrice,omitempty"`
	OrderInTotalSales  int64                        `json:"orderInTotalSales,omitempty"`
	OrderOutPrice      float64                      `json:"orderOutPrice,omitempty"`
	OrderOutTotalSales int64                        `json:"orderOutTotalSales,omitempty"`
	PicUrl             string                       `json:"picUrl,omitempty"`
	RecmdPriceType     int                          `json:"recmdPriceType,omitempty"`
	Sales              int64                        `json:"sales,omitempty"`
	Status             int                          `json:"status,omitempty"`
	StockNum           int64                        `json:"stockNum,omitempty"`
	StoreId            int64                        `json:"storeId,omitempty"`
	SuppOrderTypes     int                          `json:"suppOrderTypes,omitempty"`
	ThirdGoodsId       string                       `json:"thirdGoodsId,omitempty"`
}

type ComboSaveResponse struct {
}

func CreateComboSaveRequest() (request *ComboSaveRequest) {
	request = &ComboSaveRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "combo/save", "api")
	request.Method = api.POST
	return
}

func CreateComboSaveResponse() (response *api.BaseResponse[ComboSaveResponse]) {
	response = api.CreateResponse[ComboSaveResponse](&ComboSaveResponse{})
	return
}

type ComboSaveRequestGoodsInfos struct {
	Comment  string `json:"comment,omitempty"`
	PhotoUrl string `json:"photoUrl,omitempty"`
}

type ComboSaveRequestGroups struct {
	Items          []ComboSaveRequestItems `json:"items,omitempty"`
	AvailableCount int                     `json:"availableCount,omitempty"`
	GoodsComboId   int64                   `json:"goodsComboId,omitempty"`
	GoodsOrderType int                     `json:"goodsOrderType,omitempty"`
	GroupName      string                  `json:"groupName,omitempty"`
	Id             int64                   `json:"id,omitempty"`
	Sequence       int                     `json:"sequence,omitempty"`
}

type ComboSaveRequestItems struct {
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

type ComboSaveRequestSellTimes struct {
	TimeList []ComboSaveRequestTimeList `json:"timeList,omitempty"`
	WeekDay  int                        `json:"weekDay,omitempty"`
}

type ComboSaveRequestTimeList struct {
	EndTime   string `json:"endTime,omitempty"`
	StartTime string `json:"startTime,omitempty"`
}
