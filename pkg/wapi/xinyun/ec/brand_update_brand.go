package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// BrandUpdateBrand
// @id 1623
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改品牌)
func (client *Client) BrandUpdateBrand(request *BrandUpdateBrandRequest) (response *BrandUpdateBrandResponse, err error) {
	rpcResponse := CreateBrandUpdateBrandResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type BrandUpdateBrandRequest struct {
	*api.BaseRequest

	TrademarkAuthorizationPeriod     []int    `json:"trademarkAuthorizationPeriod,omitempty"`
	SaleAuthorization                []string `json:"saleAuthorization,omitempty"`
	Id                               int64    `json:"id,omitempty"`
	BrandManagementType              int      `json:"brandManagementType,omitempty"`
	BrandAuditType                   int      `json:"brandAuditType,omitempty"`
	StoreId                          int64    `json:"storeId,omitempty"`
	Title                            string   `json:"title,omitempty"`
	TrademarkRegistrant              string   `json:"trademarkRegistrant,omitempty"`
	TrademarkRegistrantNewNo         string   `json:"trademarkRegistrantNewNo,omitempty"`
	TrademarkType                    string   `json:"trademarkType,omitempty"`
	TrademarkApplicant               string   `json:"trademarkApplicant,omitempty"`
	TrademarkApplicationTime         int      `json:"trademarkApplicationTime,omitempty"`
	TrademarkRegistrationApplication string   `json:"trademarkRegistrationApplication,omitempty"`
	TrademarkChangeCertificate       string   `json:"trademarkChangeCertificate,omitempty"`
	ImportedGoodsForm                string   `json:"importedGoodsForm,omitempty"`
}

type BrandUpdateBrandResponse struct {
}

func CreateBrandUpdateBrandRequest() (request *BrandUpdateBrandRequest) {
	request = &BrandUpdateBrandRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "brand/updateBrand", "api")
	request.Method = api.POST
	return
}

func CreateBrandUpdateBrandResponse() (response *api.BaseResponse[BrandUpdateBrandResponse]) {
	response = api.CreateResponse[BrandUpdateBrandResponse](&BrandUpdateBrandResponse{})
	return
}
