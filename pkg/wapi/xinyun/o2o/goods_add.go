package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsAdd
// @id 63
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=添加菜品)
func (client *Client) GoodsAdd(request *GoodsAddRequest) (response *GoodsAddResponse, err error) {
	rpcResponse := CreateGoodsAddResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsAddRequest struct {
	*api.BaseRequest

	OrderInGoodsSku         GoodsAddRequestOrderInGoodsSku       `json:"orderInGoodsSku,omitempty"`
	OrderOutGoodsSku        GoodsAddRequestOrderOutGoodsSku      `json:"orderOutGoodsSku,omitempty"`
	GoodsSkuGroup           []GoodsAddRequestGoodsSkuGroup       `json:"goodsSkuGroup,omitempty"`
	OrderInGoodsTastes      []GoodsAddRequestOrderInGoodsTastes  `json:"orderInGoodsTastes,omitempty"`
	OrderOutGoodsTastes     []GoodsAddRequestOrderOutGoodsTastes `json:"orderOutGoodsTastes,omitempty"`
	OrderInSellTimes        []GoodsAddRequestOrderInSellTimes    `json:"orderInSellTimes,omitempty"`
	OrderOutSellTimes       []GoodsAddRequestOrderOutSellTimes   `json:"orderOutSellTimes,omitempty"`
	GoodsSkuPriceSales      []GoodsAddRequestGoodsSkuPriceSales  `json:"goodsSkuPriceSales,omitempty"`
	ThirdGoodsGroupIds      []string                             `json:"thirdGoodsGroupIds,omitempty"`
	GoodsName               string                               `json:"goodsName,omitempty"`
	PicUrl                  string                               `json:"picUrl,omitempty"`
	OrderInPrice            float64                              `json:"orderInPrice,omitempty"`
	OrderOutPrice           float64                              `json:"orderOutPrice,omitempty"`
	StockNum                int                                  `json:"stockNum,omitempty"`
	Status                  int                                  `json:"status,omitempty"`
	SuppOrderType           int                                  `json:"suppOrderType,omitempty"`
	GoodsDtl                string                               `json:"goodsDtl,omitempty"`
	RecmdPriceType          int                                  `json:"recmdPriceType,omitempty"`
	SkuGroupName            string                               `json:"skuGroupName,omitempty"`
	IsTastes                byte                                 `json:"isTastes,omitempty"`
	IsAdditional            byte                                 `json:"isAdditional,omitempty"`
	OrderOutGoodsAdditional string                               `json:"orderOutGoodsAdditional,omitempty"`
	OrderInGoodsAdditional  string                               `json:"orderInGoodsAdditional,omitempty"`
	ThirdGoodsId            string                               `json:"thirdGoodsId,omitempty"`
	BowlFee                 float64                              `json:"bowlFee,omitempty"`
	OrderInIncreaseUnit     int                                  `json:"orderInIncreaseUnit,omitempty"`
	OrderOutIncreaseUnit    int                                  `json:"orderOutIncreaseUnit,omitempty"`
	StoreId                 int64                                `json:"storeId,omitempty"`
}

type GoodsAddResponse struct {
}

func CreateGoodsAddRequest() (request *GoodsAddRequest) {
	request = &GoodsAddRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goods/add", "api")
	request.Method = api.POST
	return
}

func CreateGoodsAddResponse() (response *api.BaseResponse[GoodsAddResponse]) {
	response = api.CreateResponse[GoodsAddResponse](&GoodsAddResponse{})
	return
}

type GoodsAddRequestOrderInGoodsSku struct {
}

type GoodsAddRequestOrderOutGoodsSku struct {
}

type GoodsAddRequestGoodsSkuGroup struct {
	GoodsSkuGroupItem []GoodsAddRequestGoodsSkuGroupItem `json:"goodsSkuGroupItem,omitempty"`
	SkuGroupSequence  int                                `json:"skuGroupSequence,omitempty"`
}

type GoodsAddRequestGoodsSkuGroupItem struct {
	SkuName        string `json:"skuName,omitempty"`
	SkuUnitId      int64  `json:"skuUnitId,omitempty"`
	ThirdskuUnitId int64  `json:"thirdskuUnitId,omitempty"`
}

type GoodsAddRequestOrderInGoodsTastes struct {
	GoodsTastesGroupItem []GoodsAddRequestGoodsTastesGroupItem `json:"goodsTastesGroupItem,omitempty"`
	TastesGroupName      string                                `json:"tastesGroupName,omitempty"`
	TastesGroupSequence  int                                   `json:"tastesGroupSequence,omitempty"`
}

type GoodsAddRequestGoodsTastesGroupItem struct {
	GoodsTastesUnitId int64 `json:"goodsTastesUnitId,omitempty"`
	GoodsTastesName   int64 `json:"goodsTastesName,omitempty"`
}

type GoodsAddRequestOrderOutGoodsTastes struct {
	GoodsTastesGroupItem []GoodsAddRequestGoodsTastesGroupItem2 `json:"goodsTastesGroupItem,omitempty"`
	TastesGroupName      string                                 `json:"tastesGroupName,omitempty"`
	TastesGroupSequence  int                                    `json:"tastesGroupSequence,omitempty"`
}

type GoodsAddRequestGoodsTastesGroupItem2 struct {
	GoodsTastesUnitId int64 `json:"goodsTastesUnitId,omitempty"`
	GoodsTastesName   int64 `json:"goodsTastesName,omitempty"`
}

type GoodsAddRequestOrderInSellTimes struct {
	WeekDay  []map[string]any          `json:"weekDay,omitempty"`
	TimeList []GoodsAddRequestTimeList `json:"timeList,omitempty"`
}

type GoodsAddRequestTimeList struct {
	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
}

type GoodsAddRequestOrderOutSellTimes struct {
	WeekDay  []map[string]any           `json:"weekDay,omitempty"`
	TimeList []GoodsAddRequestTimeList2 `json:"timeList,omitempty"`
}

type GoodsAddRequestTimeList2 struct {
	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
}

type GoodsAddRequestGoodsSkuPriceSales struct {
	SalesPrice    float64 `json:"salesPrice,omitempty"`
	SkuItemsValue string  `json:"skuItemsValue,omitempty"`
}
