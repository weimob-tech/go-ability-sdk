package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// BrandAddBrand
// @id 1622
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增品牌)
func (client *Client) BrandAddBrand(request *BrandAddBrandRequest) (response *BrandAddBrandResponse, err error) {
	rpcResponse := CreateBrandAddBrandResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type BrandAddBrandRequest struct {
	*api.BaseRequest

	TrademarkAuthorizationPeriod     []int    `json:"trademarkAuthorizationPeriod,omitempty"`
	SaleAuthorization                []string `json:"saleAuthorization,omitempty"`
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

type BrandAddBrandResponse struct {
}

func CreateBrandAddBrandRequest() (request *BrandAddBrandRequest) {
	request = &BrandAddBrandRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "brand/addBrand", "api")
	request.Method = api.POST
	return
}

func CreateBrandAddBrandResponse() (response *api.BaseResponse[BrandAddBrandResponse]) {
	response = api.CreateResponse[BrandAddBrandResponse](&BrandAddBrandResponse{})
	return
}
