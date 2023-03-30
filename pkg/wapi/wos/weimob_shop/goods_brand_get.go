package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsBrandGet
// @id 1866
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1866?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询商品品牌详情)
func (client *Client) GoodsBrandGet(request *GoodsBrandGetRequest) (response *GoodsBrandGetResponse, err error) {
	rpcResponse := CreateGoodsBrandGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsBrandGetRequest struct {
	*api.BaseRequest

	BasicInfo GoodsBrandGetRequestBasicInfo `json:"basicInfo,omitempty"`
	BrandId   int64                         `json:"brandId,omitempty"`
}

type GoodsBrandGetResponse struct {
	Id                                   int64   `json:"id,omitempty"`
	BrandManagementType                  int64   `json:"brandManagementType,omitempty"`
	BrandAuditType                       int64   `json:"brandAuditType,omitempty"`
	Title                                string  `json:"title,omitempty"`
	TrademarkRegistrant                  string  `json:"trademarkRegistrant,omitempty"`
	TrademarkRegistrantNo                string  `json:"trademarkRegistrantNo,omitempty"`
	TrademarkType                        string  `json:"trademarkType,omitempty"`
	TrademarkStartTime                   string  `json:"trademarkStartTime,omitempty"`
	TrademarkEndTime                     string  `json:"trademarkEndTime,omitempty"`
	TrademarkApplicant                   string  `json:"trademarkApplicant,omitempty"`
	TrademarkApplicationTime             string  `json:"trademarkApplicationTime,omitempty"`
	TrademarkRegistrationApplication     string  `json:"trademarkRegistrationApplication,omitempty"`
	TrademarkChangeCertificate           string  `json:"trademarkChangeCertificate,omitempty"`
	ImportedGoodsForm                    string  `json:"importedGoodsForm,omitempty"`
	SaleAuthorization                    []int64 `json:"saleAuthorization,omitempty"`
	TrademarkTypeNumber                  int64   `json:"trademarkTypeNumber,omitempty"`
	TrademarkForeverEffective            bool    `json:"trademarkForeverEffective,omitempty"`
	TrademarkRegistrationApplicationTime int64   `json:"trademarkRegistrationApplicationTime,omitempty"`
	AuthorizationLevel                   int64   `json:"authorizationLevel,omitempty"`
	AuthorizationStartTime               int64   `json:"authorizationStartTime,omitempty"`
	AuthorizationEndTime                 int64   `json:"authorizationEndTime,omitempty"`
	AuthorizationForeverEffective        bool    `json:"authorizationForeverEffective,omitempty"`
}

func CreateGoodsBrandGetRequest() (request *GoodsBrandGetRequest) {
	request = &GoodsBrandGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/brand/get", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsBrandGetResponse() (response *api.BaseResponse[GoodsBrandGetResponse]) {
	response = api.CreateResponse[GoodsBrandGetResponse](&GoodsBrandGetResponse{})
	return
}

type GoodsBrandGetRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
