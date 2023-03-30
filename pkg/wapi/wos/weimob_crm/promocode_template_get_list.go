package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromocodeTemplateGetList
// @id 1599
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1599?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=批量获取优惠码模版列表)
func (client *Client) PromocodeTemplateGetList(request *PromocodeTemplateGetListRequest) (response *PromocodeTemplateGetListResponse, err error) {
	rpcResponse := CreatePromocodeTemplateGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromocodeTemplateGetListRequest struct {
	*api.BaseRequest

	Useful          bool    `json:"useful,omitempty"`
	CodeTemplateIds []int64 `json:"codeTemplateIds,omitempty"`
}

type PromocodeTemplateGetListResponse struct {
	AcceptTypeDTO          PromocodeTemplateGetListResponseAcceptTypeDTO `json:"acceptTypeDTO,omitempty"`
	CodeTemplateId         int64                                         `json:"codeTemplateId,omitempty"`
	LogoUrl                string                                        `json:"logoUrl,omitempty"`
	Status                 int64                                         `json:"status,omitempty"`
	IsDelete               int64                                         `json:"isDelete,omitempty"`
	CreateTime             string                                        `json:"createTime,omitempty"`
	CouponName             string                                        `json:"couponName,omitempty"`
	CouponType             int64                                         `json:"couponType,omitempty"`
	CouponAmount           int64                                         `json:"couponAmount,omitempty"`
	UseCondition           int64                                         `json:"useCondition,omitempty"`
	Desc                   string                                        `json:"desc,omitempty"`
	CodeType               int64                                         `json:"codeType,omitempty"`
	UseLimitNum            int64                                         `json:"useLimitNum,omitempty"`
	GenerateType           int64                                         `json:"generateType,omitempty"`
	Stock                  int64                                         `json:"stock,omitempty"`
	Code                   string                                        `json:"code,omitempty"`
	ValidStartTime         string                                        `json:"validStartTime,omitempty"`
	ValidEndTime           string                                        `json:"validEndTime,omitempty"`
	UseScenes              []int64                                       `json:"useScenes,omitempty"`
	BelongVid              int64                                         `json:"belongVid,omitempty"`
	Vid                    int64                                         `json:"vid,omitempty"`
	VidType                int64                                         `json:"vidType,omitempty"`
	SourceProdInstanceId   int64                                         `json:"sourceProdInstanceId,omitempty"`
	SourceProductId        int64                                         `json:"sourceProductId,omitempty"`
	SourceProductVersionId int64                                         `json:"sourceProductVersionId,omitempty"`
}

func CreatePromocodeTemplateGetListRequest() (request *PromocodeTemplateGetListRequest) {
	request = &PromocodeTemplateGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "promocode/template/getList", "apigw")
	request.Method = api.POST
	return
}

func CreatePromocodeTemplateGetListResponse() (response *api.BaseResponse[PromocodeTemplateGetListResponse]) {
	response = api.CreateResponse[PromocodeTemplateGetListResponse](&PromocodeTemplateGetListResponse{})
	return
}

type PromocodeTemplateGetListResponseAcceptTypeDTO struct {
	AcceptStoreType   int64   `json:"acceptStoreType,omitempty"`
	AcceptStoreIds    []int64 `json:"acceptStoreIds,omitempty"`
	AcceptGoodsType   int64   `json:"acceptGoodsType,omitempty"`
	AcceptGoodsIds    []int64 `json:"acceptGoodsIds,omitempty"`
	ExistExcludeGoods bool    `json:"existExcludeGoods,omitempty"`
	ExcludeGoodsIds   []int64 `json:"excludeGoodsIds,omitempty"`
}
