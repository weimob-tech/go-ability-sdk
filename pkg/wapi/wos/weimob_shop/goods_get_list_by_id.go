package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsGetListById
// @id 1642
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1642?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=根据商品ID列表查询商品信息)
func (client *Client) GoodsGetListById(request *GoodsGetListByIdRequest) (response *GoodsGetListByIdResponse, err error) {
	rpcResponse := CreateGoodsGetListByIdResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsGetListByIdRequest struct {
	*api.BaseRequest

	BasicInfo   GoodsGetListByIdRequestBasicInfo `json:"basicInfo,omitempty"`
	GoodsIdList []int64                          `json:"goodsIdList,omitempty"`
}

type GoodsGetListByIdResponse struct {
	GoodsStock      GoodsGetListByIdResponseGoodsStock `json:"goodsStock,omitempty"`
	GoodsPrice      GoodsGetListByIdResponseGoodsPrice `json:"goodsPrice,omitempty"`
	GoodsId         int64                              `json:"goodsId,omitempty"`
	Title           string                             `json:"title,omitempty"`
	SubTitle        string                             `json:"subTitle,omitempty"`
	CreateVid       int64                              `json:"createVid,omitempty"`
	DefaultImageUrl string                             `json:"defaultImageUrl,omitempty"`
	PayGoodsNum     int64                              `json:"payGoodsNum,omitempty"`
	IsMultiSku      bool                               `json:"isMultiSku,omitempty"`
	IsOnline        bool                               `json:"isOnline,omitempty"`
	IsCanSell       bool                               `json:"isCanSell,omitempty"`
	GoodsType       int64                              `json:"goodsType,omitempty"`
	SubGoodsType    int64                              `json:"subGoodsType,omitempty"`
	OuterGoodsCode  string                             `json:"outerGoodsCode,omitempty"`
}

func CreateGoodsGetListByIdRequest() (request *GoodsGetListByIdRequest) {
	request = &GoodsGetListByIdRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/getListById", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsGetListByIdResponse() (response *api.BaseResponse[GoodsGetListByIdResponse]) {
	response = api.CreateResponse[GoodsGetListByIdResponse](&GoodsGetListByIdResponse{})
	return
}

type GoodsGetListByIdRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type GoodsGetListByIdResponseGoodsStock struct {
	GoodsStockNum int64 `json:"goodsStockNum,omitempty"`
}

type GoodsGetListByIdResponseGoodsPrice struct {
	MinSalePrice int64 `json:"minSalePrice,omitempty"`
	MaxSalePrice int64 `json:"maxSalePrice,omitempty"`
}
