package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsEdit
// @id 64
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改菜品)
func (client *Client) GoodsEdit(request *GoodsEditRequest) (response *GoodsEditResponse, err error) {
	rpcResponse := CreateGoodsEditResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsEditRequest struct {
	*api.BaseRequest

	OrderInGoodsSku         GoodsEditRequestOrderInGoodsSku      `json:"orderInGoodsSku,omitempty"`
	OrderInSellTimes        []GoodsEditRequestOrderInSellTimes   `json:"orderInSellTimes,omitempty"`
	OrderOutSellTimes       []GoodsEditRequestOrderOutSellTimes  `json:"orderOutSellTimes,omitempty"`
	GoodsSkuPriceSales      []GoodsEditRequestGoodsSkuPriceSales `json:"goodsSkuPriceSales,omitempty"`
	GoodsSkuGroup           []GoodsEditRequestGoodsSkuGroup      `json:"goodsSkuGroup,omitempty"`
	ThirdGoodsGroupIds      []map[string]any                     `json:"thirdGoodsGroupIds,omitempty"`
	GoodsId                 int64                                `json:"goodsId,omitempty"`
	GoodsName               string                               `json:"goodsName,omitempty"`
	PicUrl                  string                               `json:"picUrl,omitempty"`
	OrderInPrice            int64                                `json:"orderInPrice,omitempty"`
	OrderOutPrice           int64                                `json:"orderOutPrice,omitempty"`
	StockNum                int                                  `json:"stockNum,omitempty"`
	Status                  int                                  `json:"status,omitempty"`
	SuppOrderType           int                                  `json:"suppOrderType,omitempty"`
	GoodsDtl                string                               `json:"goodsDtl,omitempty"`
	RecmdPriceType          int                                  `json:"recmdPriceType,omitempty"`
	OrderOutGoodsSku        string                               `json:"orderOutGoodsSku,omitempty"`
	IsTastes                byte                                 `json:"isTastes,omitempty"`
	IsAdditional            byte                                 `json:"isAdditional,omitempty"`
	OrderInGoodsAdditional  string                               `json:"orderInGoodsAdditional,omitempty"`
	OrderOutGoodsAdditional string                               `json:"orderOutGoodsAdditional,omitempty"`
	OrderInGoodsTastes      string                               `json:"orderInGoodsTastes,omitempty"`
	OrderOutGoodsTastes     string                               `json:"orderOutGoodsTastes,omitempty"`
	ThirdGoodsId            string                               `json:"thirdGoodsId,omitempty"`
	OrderInIncreaseUnit     int                                  `json:"orderInIncreaseUnit,omitempty"`
	OrderOutIncreaseUnit    int                                  `json:"orderOutIncreaseUnit,omitempty"`
}

type GoodsEditResponse struct {
}

func CreateGoodsEditRequest() (request *GoodsEditRequest) {
	request = &GoodsEditRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goods/edit", "api")
	request.Method = api.POST
	return
}

func CreateGoodsEditResponse() (response *api.BaseResponse[GoodsEditResponse]) {
	response = api.CreateResponse[GoodsEditResponse](&GoodsEditResponse{})
	return
}

type GoodsEditRequestOrderInGoodsSku struct {
}

type GoodsEditRequestOrderInSellTimes struct {
	WeekDay  []map[string]any           `json:"weekDay,omitempty"`
	TimeList []GoodsEditRequestTimeList `json:"timeList,omitempty"`
}

type GoodsEditRequestTimeList struct {
	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
}

type GoodsEditRequestOrderOutSellTimes struct {
	WeekDay  []map[string]any            `json:"weekDay,omitempty"`
	TimeList []GoodsEditRequestTimeList2 `json:"timeList,omitempty"`
}

type GoodsEditRequestTimeList2 struct {
	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
}

type GoodsEditRequestGoodsSkuPriceSales struct {
	SalesPrice    int64  `json:"salesPrice,omitempty"`
	SkuItemsValue string `json:"skuItemsValue,omitempty"`
}

type GoodsEditRequestGoodsSkuGroup struct {
	GoodsSkuGroupItem []GoodsEditRequestGoodsSkuGroupItem `json:"goodsSkuGroupItem,omitempty"`
	SkuGroupSequence  int                                 `json:"skuGroupSequence,omitempty"`
	SkuGroupName      string                              `json:"skuGroupName,omitempty"`
}

type GoodsEditRequestGoodsSkuGroupItem struct {
	SkuName        int64 `json:"skuName,omitempty"`
	SkuUnitId      int64 `json:"skuUnitId,omitempty"`
	ThirdskuUnitId int64 `json:"thirdskuUnitId,omitempty"`
}
