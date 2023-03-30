package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsCategorySkuspecContentGetList
// @id 1520
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1520?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询规格值列表)
func (client *Client) GoodsCategorySkuspecContentGetList(request *GoodsCategorySkuspecContentGetListRequest) (response *GoodsCategorySkuspecContentGetListResponse, err error) {
	rpcResponse := CreateGoodsCategorySkuspecContentGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsCategorySkuspecContentGetListRequest struct {
	*api.BaseRequest

	QueryParameter GoodsCategorySkuspecContentGetListRequestQueryParameter `json:"queryParameter,omitempty"`
	BasicInfo      GoodsCategorySkuspecContentGetListRequestBasicInfo      `json:"basicInfo,omitempty"`
	PageNum        int64                                                   `json:"pageNum,omitempty"`
	PageSize       int64                                                   `json:"pageSize,omitempty"`
}

type GoodsCategorySkuspecContentGetListResponse struct {
	PageList   []GoodsCategorySkuspecContentGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                                `json:"pageNum,omitempty"`
	PageSize   int64                                                `json:"pageSize,omitempty"`
	TotalCount int64                                                `json:"totalCount,omitempty"`
}

func CreateGoodsCategorySkuspecContentGetListRequest() (request *GoodsCategorySkuspecContentGetListRequest) {
	request = &GoodsCategorySkuspecContentGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/category/skuspec/content/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsCategorySkuspecContentGetListResponse() (response *api.BaseResponse[GoodsCategorySkuspecContentGetListResponse]) {
	response = api.CreateResponse[GoodsCategorySkuspecContentGetListResponse](&GoodsCategorySkuspecContentGetListResponse{})
	return
}

type GoodsCategorySkuspecContentGetListRequestQueryParameter struct {
	SpecId int64  `json:"specId,omitempty"`
	Search string `json:"search,omitempty"`
}

type GoodsCategorySkuspecContentGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type GoodsCategorySkuspecContentGetListResponsePageList struct {
	SpecValueId   int64  `json:"specValueId,omitempty"`
	SpecValueName string `json:"specValueName,omitempty"`
}
