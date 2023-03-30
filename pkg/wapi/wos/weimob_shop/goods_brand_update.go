package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsBrandUpdate
// @id 1832
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1832?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=修改商品品牌)
func (client *Client) GoodsBrandUpdate(request *GoodsBrandUpdateRequest) (response *GoodsBrandUpdateResponse, err error) {
	rpcResponse := CreateGoodsBrandUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsBrandUpdateRequest struct {
	*api.BaseRequest

	BasicInfo                            GoodsBrandUpdateRequestBasicInfo `json:"basicInfo,omitempty"`
	BrandId                              int64                            `json:"brandId,omitempty"`
	BrandManagementType                  string                           `json:"brandManagementType,omitempty"`
	BrandAuditType                       string                           `json:"brandAuditType,omitempty"`
	Title                                string                           `json:"title,omitempty"`
	TrademarkRegistrant                  string                           `json:"trademarkRegistrant,omitempty"`
	TrademarkRegistrantNo                string                           `json:"trademarkRegistrantNo,omitempty"`
	TrademarkType                        string                           `json:"trademarkType,omitempty"`
	TrademarkStartTime                   int64                            `json:"trademarkStartTime,omitempty"`
	TrademarkEndTime                     int64                            `json:"trademarkEndTime,omitempty"`
	TrademarkApplicant                   string                           `json:"trademarkApplicant,omitempty"`
	TrademarkApplicationTime             int64                            `json:"trademarkApplicationTime,omitempty"`
	TrademarkRegistrationApplication     string                           `json:"trademarkRegistrationApplication,omitempty"`
	TrademarkChangeCertificate           string                           `json:"trademarkChangeCertificate,omitempty"`
	ImportedGoodsForm                    string                           `json:"importedGoodsForm,omitempty"`
	SaleAuthorization                    string                           `json:"saleAuthorization,omitempty"`
	TrademarkTypeNumber                  int64                            `json:"trademarkTypeNumber,omitempty"`
	TrademarkForeverEffective            bool                             `json:"trademarkForeverEffective,omitempty"`
	TrademarkRegistrationApplicationTime int64                            `json:"trademarkRegistrationApplicationTime,omitempty"`
	AuthorizationLevel                   int64                            `json:"authorizationLevel,omitempty"`
	AuthorizationStartTime               int64                            `json:"authorizationStartTime,omitempty"`
	AuthorizationEndTime                 int64                            `json:"authorizationEndTime,omitempty"`
	AuthorizationForeverEffective        bool                             `json:"authorizationForeverEffective,omitempty"`
}

type GoodsBrandUpdateResponse struct {
	ReturnResult bool `json:"returnResult,omitempty"`
}

func CreateGoodsBrandUpdateRequest() (request *GoodsBrandUpdateRequest) {
	request = &GoodsBrandUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/brand/update", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsBrandUpdateResponse() (response *api.BaseResponse[GoodsBrandUpdateResponse]) {
	response = api.CreateResponse[GoodsBrandUpdateResponse](&GoodsBrandUpdateResponse{})
	return
}

type GoodsBrandUpdateRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
