package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsBrandInsert
// @id 1831
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1831?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=添加品牌)
func (client *Client) GoodsBrandInsert(request *GoodsBrandInsertRequest) (response *GoodsBrandInsertResponse, err error) {
	rpcResponse := CreateGoodsBrandInsertResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsBrandInsertRequest struct {
	*api.BaseRequest

	BasicInfo                            GoodsBrandInsertRequestBasicInfo `json:"basicInfo,omitempty"`
	BrandManagementType                  int64                            `json:"brandManagementType,omitempty"`
	BrandAuditType                       int64                            `json:"brandAuditType,omitempty"`
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
	SaleAuthorization                    []int64                          `json:"saleAuthorization,omitempty"`
	TrademarkTypeNumber                  int64                            `json:"trademarkTypeNumber,omitempty"`
	TrademarkForeverEffective            bool                             `json:"trademarkForeverEffective,omitempty"`
	TrademarkRegistrationApplicationTime int64                            `json:"trademarkRegistrationApplicationTime,omitempty"`
	AuthorizationLevel                   int64                            `json:"authorizationLevel,omitempty"`
	AuthorizationStartTime               int64                            `json:"authorizationStartTime,omitempty"`
	AuthorizationEndTime                 int64                            `json:"authorizationEndTime,omitempty"`
	AuthorizationForeverEffective        bool                             `json:"authorizationForeverEffective,omitempty"`
}

type GoodsBrandInsertResponse struct {
	Id int64 `json:"id,omitempty"`
}

func CreateGoodsBrandInsertRequest() (request *GoodsBrandInsertRequest) {
	request = &GoodsBrandInsertRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/brand/insert", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsBrandInsertResponse() (response *api.BaseResponse[GoodsBrandInsertResponse]) {
	response = api.CreateResponse[GoodsBrandInsertResponse](&GoodsBrandInsertResponse{})
	return
}

type GoodsBrandInsertRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
