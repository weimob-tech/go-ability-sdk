package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromocodeLock
// @id 1597
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1597?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=优惠码锁定)
func (client *Client) PromocodeLock(request *PromocodeLockRequest) (response *PromocodeLockResponse, err error) {
	rpcResponse := CreatePromocodeLockResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromocodeLockRequest struct {
	*api.BaseRequest

	StoreGoodsInfoDTOS []PromocodeLockRequestStoreGoodsInfoDTOS `json:"storeGoodsInfoDTOS,omitempty"`
	CodeList           []PromocodeLockRequestCodeList           `json:"codeList,omitempty"`
	UseScene           int64                                    `json:"useScene,omitempty"`
	Wid                int64                                    `json:"wid,omitempty"`
	VidType            int64                                    `json:"vidType,omitempty"`
	Vid                int64                                    `json:"vid,omitempty"`
}

type PromocodeLockResponse struct {
}

func CreatePromocodeLockRequest() (request *PromocodeLockRequest) {
	request = &PromocodeLockRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "promocode/lock", "apigw")
	request.Method = api.POST
	return
}

func CreatePromocodeLockResponse() (response *api.BaseResponse[PromocodeLockResponse]) {
	response = api.CreateResponse[PromocodeLockResponse](&PromocodeLockResponse{})
	return
}

type PromocodeLockRequestStoreGoodsInfoDTOS struct {
	VidNodeList PromocodeLockRequestVidNodeList  `json:"vidNodeList,omitempty"`
	GoodsInfos  []PromocodeLockRequestGoodsInfos `json:"goodsInfos,omitempty"`
}

type PromocodeLockRequestVidNodeList struct {
	BosId   int64 `json:"bosId,omitempty"`
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type PromocodeLockRequestGoodsInfos struct {
	Skus        []PromocodeLockRequestSkus `json:"skus,omitempty"`
	Id          int64                      `json:"id,omitempty"`
	CategoryIds []int64                    `json:"categoryIds,omitempty"`
	Tags        []int64                    `json:"tags,omitempty"`
	GroupIds    []int64                    `json:"groupIds,omitempty"`
}

type PromocodeLockRequestSkus struct {
	Id       int64 `json:"id,omitempty"`
	Price    int64 `json:"price,omitempty"`
	Num      int64 `json:"num,omitempty"`
	Selected bool  `json:"selected,omitempty"`
}

type PromocodeLockRequestCodeList struct {
	CodeTemplateId int64  `json:"codeTemplateId,omitempty"`
	OrderNo        string `json:"orderNo,omitempty"`
	Code           string `json:"code,omitempty"`
	Amount         int64  `json:"amount,omitempty"`
}
