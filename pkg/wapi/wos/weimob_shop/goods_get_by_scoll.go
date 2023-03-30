package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsGetByScoll
// @id 2032
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2032?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=滚动查询商品简要信息)
func (client *Client) GoodsGetByScoll(request *GoodsGetByScollRequest) (response *GoodsGetByScollResponse, err error) {
	rpcResponse := CreateGoodsGetByScollResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsGetByScollRequest struct {
	*api.BaseRequest

	BasicInfo    GoodsGetByScollRequestBasicInfo `json:"basicInfo,omitempty"`
	PageSize     int64                           `json:"pageSize,omitempty"`
	StartGoodsId int64                           `json:"startGoodsId,omitempty"`
}

type GoodsGetByScollResponse struct {
	GoodInfoList []GoodsGetByScollResponseGoodInfoList `json:"goodInfoList,omitempty"`
	LastGoodsId  int64                                 `json:"lastGoodsId,omitempty"`
}

func CreateGoodsGetByScollRequest() (request *GoodsGetByScollRequest) {
	request = &GoodsGetByScollRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/getByScoll", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsGetByScollResponse() (response *api.BaseResponse[GoodsGetByScollResponse]) {
	response = api.CreateResponse[GoodsGetByScollResponse](&GoodsGetByScollResponse{})
	return
}

type GoodsGetByScollRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type GoodsGetByScollResponseGoodInfoList struct {
	SkuList         []GoodsGetByScollResponseSkuList `json:"skuList,omitempty"`
	GoodsId         int64                            `json:"goodsId,omitempty"`
	GoodsType       int64                            `json:"goodsType,omitempty"`
	SubGoodsType    int64                            `json:"subGoodsType,omitempty"`
	Title           string                           `json:"title,omitempty"`
	DefaultImageUrl string                           `json:"defaultImageUrl,omitempty"`
	OuterGoodsCode  string                           `json:"outerGoodsCode,omitempty"`
	SaleChannelType int64                            `json:"saleChannelType,omitempty"`
	SoldType        int64                            `json:"soldType,omitempty"`
	IsCanSell       bool                             `json:"isCanSell,omitempty"`
	IsOnline        bool                             `json:"isOnline,omitempty"`
	OnlineTime      int64                            `json:"onlineTime,omitempty"`
}

type GoodsGetByScollResponseSkuList struct {
	GoodsId      int64  `json:"goodsId,omitempty"`
	SkuId        int64  `json:"skuId,omitempty"`
	OuterSkuCode string `json:"outerSkuCode,omitempty"`
	ImageUrl     string `json:"imageUrl,omitempty"`
	IsDisabled   bool   `json:"isDisabled,omitempty"`
	SkuBarCode   string `json:"skuBarCode,omitempty"`
}
