package weimob_shopexpress

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsAdd
// @id 1959
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1959?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商品添加)
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

	BaseIn           GoodsAddRequestBaseIn            `json:"baseIn,omitempty"`
	FreightPlanList  []GoodsAddRequestFreightPlanList `json:"freightPlanList,omitempty"`
	MaterialList     []GoodsAddRequestMaterialList    `json:"materialList,omitempty"`
	SkuList          []GoodsAddRequestSkuList         `json:"skuList,omitempty"`
	Seo              GoodsAddRequestSeo               `json:"seo,omitempty"`
	Detail           string                           `json:"detail,omitempty"`
	GoodsTagCodeList []int64                          `json:"goodsTagCodeList,omitempty"`
	GroupCodeList    []int64                          `json:"groupCodeList,omitempty"`
}

type GoodsAddResponse struct {
	GoodsCode string `json:"goodsCode,omitempty"`
}

func CreateGoodsAddRequest() (request *GoodsAddRequest) {
	request = &GoodsAddRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shopexpress", "v2.0", "goods/add", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsAddResponse() (response *api.BaseResponse[GoodsAddResponse]) {
	response = api.CreateResponse[GoodsAddResponse](&GoodsAddResponse{})
	return
}

type GoodsAddRequestBaseIn struct {
	SpecList           []GoodsAddRequestSpecList `json:"specList,omitempty"`
	Subtitle           string                    `json:"subtitle,omitempty"`
	IfSingleSpec       bool                      `json:"ifSingleSpec,omitempty"`
	ImgUrl             string                    `json:"imgUrl,omitempty"`
	Name               string                    `json:"name,omitempty"`
	Sort               int64                     `json:"sort,omitempty"`
	Status             int64                     `json:"status,omitempty"`
	TimingOnShelfState bool                      `json:"timingOnShelfState,omitempty"`
	TimingOnShelfTime  string                    `json:"timingOnShelfTime,omitempty"`
}

type GoodsAddRequestSpecList struct {
	SpecValueList []GoodsAddRequestSpecValueList `json:"specValueList,omitempty"`
	SpecKeyName   string                         `json:"specKeyName,omitempty"`
}

type GoodsAddRequestSpecValueList struct {
	SpecValueName string `json:"specValueName,omitempty"`
}

type GoodsAddRequestFreightPlanList struct {
	FreightTemplateNo string `json:"freightTemplateNo,omitempty"`
}

type GoodsAddRequestMaterialList struct {
	Alt  string `json:"alt,omitempty"`
	Sort int64  `json:"sort,omitempty"`
	Type int64  `json:"type,omitempty"`
	Url  string `json:"url,omitempty"`
}

type GoodsAddRequestSkuList struct {
	GoodsSkuSpecList []GoodsAddRequestGoodsSkuSpecList `json:"goodsSkuSpecList,omitempty"`
	BarCode          string                            `json:"barCode,omitempty"`
	CustomSkuCode    string                            `json:"customSkuCode,omitempty"`
	IfTax            bool                              `json:"ifTax,omitempty"`
	ImgUrl           string                            `json:"imgUrl,omitempty"`
	MarketPrice      int64                             `json:"marketPrice,omitempty"`
	Price            int64                             `json:"price,omitempty"`
	SkuName          string                            `json:"skuName,omitempty"`
	Sort             int64                             `json:"sort,omitempty"`
	Status           int64                             `json:"status,omitempty"`
	StockCount       int64                             `json:"stockCount,omitempty"`
	StockTrackState  bool                              `json:"stockTrackState,omitempty"`
	WeightUnit       string                            `json:"weightUnit,omitempty"`
	WeightValue      int64                             `json:"weightValue,omitempty"`
}

type GoodsAddRequestGoodsSkuSpecList struct {
	SpecKeyName   string `json:"specKeyName,omitempty"`
	SpecValueName string `json:"specValueName,omitempty"`
}

type GoodsAddRequestSeo struct {
	SeoTitle      string `json:"seoTitle,omitempty"`
	SeoDesc       string `json:"seoDesc,omitempty"`
	SeoKeyWord    string `json:"seoKeyWord,omitempty"`
	SeoLink       string `json:"seoLink,omitempty"`
	IfOldRedirect bool   `json:"ifOldRedirect,omitempty"`
}
