package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsGetList
// @id 1628
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1628?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询商品列表)
func (client *Client) GoodsGetList(request *GoodsGetListRequest) (response *GoodsGetListResponse, err error) {
	rpcResponse := CreateGoodsGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsGetListRequest struct {
	*api.BaseRequest

	QueryParameter GoodsGetListRequestQueryParameter `json:"queryParameter,omitempty"`
	BasicInfo      GoodsGetListRequestBasicInfo      `json:"basicInfo,omitempty"`
	PageNum        int64                             `json:"pageNum,omitempty"`
	PageSize       int64                             `json:"pageSize,omitempty"`
}

type GoodsGetListResponse struct {
	PageList   []GoodsGetListResponsePageList `json:"pageList,omitempty"`
	PageSize   int64                          `json:"pageSize,omitempty"`
	TotalCount int64                          `json:"totalCount,omitempty"`
	PageNum    int64                          `json:"pageNum,omitempty"`
}

func CreateGoodsGetListRequest() (request *GoodsGetListRequest) {
	request = &GoodsGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsGetListResponse() (response *api.BaseResponse[GoodsGetListResponse]) {
	response = api.CreateResponse[GoodsGetListResponse](&GoodsGetListResponse{})
	return
}

type GoodsGetListRequestQueryParameter struct {
	GoodsStatus      int64   `json:"goodsStatus,omitempty"`
	ClassifyId       int64   `json:"classifyId,omitempty"`
	Search           string  `json:"search,omitempty"`
	Sort             int64   `json:"sort,omitempty"`
	SearchType       int64   `json:"searchType,omitempty"`
	MinSalePrice     int64   `json:"minSalePrice,omitempty"`
	MaxSalePrice     int64   `json:"maxSalePrice,omitempty"`
	SearchOptionType int64   `json:"searchOptionType,omitempty"`
	GoodsTagIdList   []int64 `json:"goodsTagIdList,omitempty"`
	StartUpdateTime  int64   `json:"startUpdateTime,omitempty"`
	EndUpdateTime    int64   `json:"endUpdateTime,omitempty"`
}

type GoodsGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type GoodsGetListResponsePageList struct {
	GoodsStock      GoodsGetListResponseGoodsStock `json:"goodsStock,omitempty"`
	GoodsPrice      GoodsGetListResponseGoodsPrice `json:"goodsPrice,omitempty"`
	GoodsId         int64                          `json:"goodsId,omitempty"`
	Title           string                         `json:"title,omitempty"`
	SubTitle        string                         `json:"subTitle,omitempty"`
	CreateVid       int64                          `json:"createVid,omitempty"`
	DefaultImageUrl string                         `json:"defaultImageUrl,omitempty"`
	OuterGoodsCode  string                         `json:"outerGoodsCode,omitempty"`
	RealSaleNum     int64                          `json:"realSaleNum,omitempty"`
	SoldType        int64                          `json:"soldType,omitempty"`
	GoodsType       int64                          `json:"goodsType,omitempty"`
	SubGoodsType    int64                          `json:"subGoodsType,omitempty"`
	IsMultiSku      bool                           `json:"isMultiSku,omitempty"`
	IsOnline        bool                           `json:"isOnline,omitempty"`
	IsCanSell       bool                           `json:"isCanSell,omitempty"`
	Sort            int64                          `json:"sort,omitempty"`
	OnlineTime      int64                          `json:"onlineTime,omitempty"`
}

type GoodsGetListResponseGoodsStock struct {
	GoodsStockNum int64 `json:"goodsStockNum,omitempty"`
}

type GoodsGetListResponseGoodsPrice struct {
	MinSalePrice int64 `json:"minSalePrice,omitempty"`
	MaxSalePrice int64 `json:"maxSalePrice,omitempty"`
}
