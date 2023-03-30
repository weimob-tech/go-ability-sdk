package weimob_shopexpress

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsGet
// @id 1961
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1961?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商品详情)
func (client *Client) GoodsGet(request *GoodsGetRequest) (response *GoodsGetResponse, err error) {
	rpcResponse := CreateGoodsGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsGetRequest struct {
	*api.BaseRequest

	GoodsCode string `json:"goodsCode,omitempty"`
}

type GoodsGetResponse struct {
	MaterialList       []GoodsGetResponseMaterialList    `json:"materialList,omitempty"`
	SkuList            []GoodsGetResponseSkuList         `json:"skuList,omitempty"`
	SpecList           []GoodsGetResponseSpecList        `json:"specList,omitempty"`
	Seo                GoodsGetResponseSeo               `json:"seo,omitempty"`
	FreightPlanList    []GoodsGetResponseFreightPlanList `json:"freightPlanList,omitempty"`
	GoodsCode          string                            `json:"goodsCode,omitempty"`
	Name               string                            `json:"name,omitempty"`
	ImgUrl             string                            `json:"imgUrl,omitempty"`
	Status             int64                             `json:"status,omitempty"`
	IfSingleSpec       bool                              `json:"ifSingleSpec,omitempty"`
	Sort               int64                             `json:"sort,omitempty"`
	Detail             string                            `json:"detail,omitempty"`
	GroupCodeList      []int64                           `json:"groupCodeList,omitempty"`
	GoodsTagCodeList   []int64                           `json:"goodsTagCodeList,omitempty"`
	TimingOnShelfState bool                              `json:"timingOnShelfState,omitempty"`
	TimingOnShelfTime  string                            `json:"timingOnShelfTime,omitempty"`
	Subtitle           string                            `json:"subtitle,omitempty"`
}

func CreateGoodsGetRequest() (request *GoodsGetRequest) {
	request = &GoodsGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shopexpress", "v2.0", "goods/get", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsGetResponse() (response *api.BaseResponse[GoodsGetResponse]) {
	response = api.CreateResponse[GoodsGetResponse](&GoodsGetResponse{})
	return
}

type GoodsGetResponseMaterialList struct {
	Type int64  `json:"type,omitempty"`
	Url  string `json:"url,omitempty"`
	Sort int64  `json:"sort,omitempty"`
	Alt  string `json:"alt,omitempty"`
}

type GoodsGetResponseSkuList struct {
	GoodsSkuSpecList []GoodsGetResponseGoodsSkuSpecList `json:"goodsSkuSpecList,omitempty"`
	SkuCode          string                             `json:"skuCode,omitempty"`
	CustomSkuCode    string                             `json:"customSkuCode,omitempty"`
	SkuName          string                             `json:"skuName,omitempty"`
	Status           int64                              `json:"status,omitempty"`
	ImgUrl           string                             `json:"imgUrl,omitempty"`
	Price            string                             `json:"price,omitempty"`
	MarketPrice      string                             `json:"marketPrice,omitempty"`
	StockCount       int64                              `json:"stockCount,omitempty"`
	IfTax            bool                               `json:"ifTax,omitempty"`
	StockTrackState  bool                               `json:"stockTrackState,omitempty"`
	WeightValue      string                             `json:"weightValue,omitempty"`
	WeightUnit       string                             `json:"weightUnit,omitempty"`
	BarCode          string                             `json:"barCode,omitempty"`
}

type GoodsGetResponseGoodsSkuSpecList struct {
	SpecKey   string `json:"specKey,omitempty"`
	SpecValue string `json:"specValue,omitempty"`
}

type GoodsGetResponseSpecList struct {
	SpecValueList []GoodsGetResponseSpecValueList `json:"specValueList,omitempty"`
	SpecKeyName   string                          `json:"specKeyName,omitempty"`
}

type GoodsGetResponseSpecValueList struct {
	SpecValueName string `json:"specValueName,omitempty"`
}

type GoodsGetResponseSeo struct {
	SeoTitle      string `json:"seoTitle,omitempty"`
	SeoDesc       string `json:"seoDesc,omitempty"`
	SeoLink       string `json:"seoLink,omitempty"`
	IfOldRedirect bool   `json:"ifOldRedirect,omitempty"`
}

type GoodsGetResponseFreightPlanList struct {
	FreightTemplateNo     string `json:"freightTemplateNo,omitempty"`
	FreightTemplateRuleNo string `json:"freightTemplateRuleNo,omitempty"`
}
