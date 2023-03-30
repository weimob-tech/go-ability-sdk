package weimob_shopexpress

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsUpdate
// @id 1958
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1958?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商品更新)
func (client *Client) GoodsUpdate(request *GoodsUpdateRequest) (response *GoodsUpdateResponse, err error) {
	rpcResponse := CreateGoodsUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsUpdateRequest struct {
	*api.BaseRequest

	Seo              GoodsUpdateRequestSeo               `json:"seo,omitempty"`
	BaseIn           GoodsUpdateRequestBaseIn            `json:"baseIn,omitempty"`
	FreightPlanList  []GoodsUpdateRequestFreightPlanList `json:"freightPlanList,omitempty"`
	MaterialList     []GoodsUpdateRequestMaterialList    `json:"materialList,omitempty"`
	SkuList          []GoodsUpdateRequestSkuList         `json:"skuList,omitempty"`
	Detail           string                              `json:"detail,omitempty"`
	GoodsTagCodeList []int64                             `json:"goodsTagCodeList,omitempty"`
	GroupCodeList    []int64                             `json:"groupCodeList,omitempty"`
	OperationName    string                              `json:"operationName,omitempty"`
}

type GoodsUpdateResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateGoodsUpdateRequest() (request *GoodsUpdateRequest) {
	request = &GoodsUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shopexpress", "v2.0", "goods/update", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsUpdateResponse() (response *api.BaseResponse[GoodsUpdateResponse]) {
	response = api.CreateResponse[GoodsUpdateResponse](&GoodsUpdateResponse{})
	return
}

type GoodsUpdateRequestSeo struct {
	IfOldRedirect bool   `json:"ifOldRedirect,omitempty"`
	SeoDesc       string `json:"seoDesc,omitempty"`
	SeoKeyWord    string `json:"seoKeyWord,omitempty"`
	SeoLink       string `json:"seoLink,omitempty"`
	SeoTitle      string `json:"seoTitle,omitempty"`
}

type GoodsUpdateRequestBaseIn struct {
	SpecList           []GoodsUpdateRequestSpecList `json:"specList,omitempty"`
	Subtitle           string                       `json:"subtitle,omitempty"`
	GoodsCode          string                       `json:"goodsCode,omitempty"`
	IfSingleSpec       bool                         `json:"ifSingleSpec,omitempty"`
	ImgUrl             string                       `json:"imgUrl,omitempty"`
	Name               string                       `json:"name,omitempty"`
	Sort               int64                        `json:"sort,omitempty"`
	Status             int64                        `json:"status,omitempty"`
	TimingOnShelfState bool                         `json:"timingOnShelfState,omitempty"`
	TimingOnShelfTime  string                       `json:"timingOnShelfTime,omitempty"`
}

type GoodsUpdateRequestSpecList struct {
	SpecValueList []GoodsUpdateRequestSpecValueList `json:"specValueList,omitempty"`
	SpecKeyName   string                            `json:"specKeyName,omitempty"`
}

type GoodsUpdateRequestSpecValueList struct {
	SpecValueName string `json:"specValueName,omitempty"`
}

type GoodsUpdateRequestFreightPlanList struct {
	FreightTemplateNo string `json:"freightTemplateNo,omitempty"`
}

type GoodsUpdateRequestMaterialList struct {
	Alt  string `json:"alt,omitempty"`
	Sort int64  `json:"sort,omitempty"`
	Type int64  `json:"type,omitempty"`
	Url  string `json:"url,omitempty"`
}

type GoodsUpdateRequestSkuList struct {
	GoodsSkuSpecList []GoodsUpdateRequestGoodsSkuSpecList `json:"goodsSkuSpecList,omitempty"`
	BarCode          string                               `json:"barCode,omitempty"`
	CustomSkuCode    string                               `json:"customSkuCode,omitempty"`
	IfTax            bool                                 `json:"ifTax,omitempty"`
	ImgUrl           string                               `json:"imgUrl,omitempty"`
	MarketPrice      int64                                `json:"marketPrice,omitempty"`
	Price            int64                                `json:"price,omitempty"`
	SkuCode          string                               `json:"skuCode,omitempty"`
	SkuName          string                               `json:"skuName,omitempty"`
	Sort             int64                                `json:"sort,omitempty"`
	Status           int64                                `json:"status,omitempty"`
	StockCount       int64                                `json:"stockCount,omitempty"`
	StockTrackState  bool                                 `json:"stockTrackState,omitempty"`
	WeightUnit       string                               `json:"weightUnit,omitempty"`
	WeightValue      int64                                `json:"weightValue,omitempty"`
}

type GoodsUpdateRequestGoodsSkuSpecList struct {
	SpecKeyName   string `json:"specKeyName,omitempty"`
	SpecValueName string `json:"specValueName,omitempty"`
}
