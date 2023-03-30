package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsCategorySkuspecGetList
// @id 1521
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1521?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询规格列表)
func (client *Client) GoodsCategorySkuspecGetList(request *GoodsCategorySkuspecGetListRequest) (response *GoodsCategorySkuspecGetListResponse, err error) {
	rpcResponse := CreateGoodsCategorySkuspecGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsCategorySkuspecGetListRequest struct {
	*api.BaseRequest

	QueryParameter GoodsCategorySkuspecGetListRequestQueryParameter `json:"queryParameter,omitempty"`
	BasicInfo      GoodsCategorySkuspecGetListRequestBasicInfo      `json:"basicInfo,omitempty"`
	PageNum        int64                                            `json:"pageNum,omitempty"`
	PageSize       int64                                            `json:"pageSize,omitempty"`
}

type GoodsCategorySkuspecGetListResponse struct {
	PageList   []GoodsCategorySkuspecGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                         `json:"pageNum,omitempty"`
	PageSize   int64                                         `json:"pageSize,omitempty"`
	TotalCount int64                                         `json:"totalCount,omitempty"`
}

func CreateGoodsCategorySkuspecGetListRequest() (request *GoodsCategorySkuspecGetListRequest) {
	request = &GoodsCategorySkuspecGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/category/skuspec/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsCategorySkuspecGetListResponse() (response *api.BaseResponse[GoodsCategorySkuspecGetListResponse]) {
	response = api.CreateResponse[GoodsCategorySkuspecGetListResponse](&GoodsCategorySkuspecGetListResponse{})
	return
}

type GoodsCategorySkuspecGetListRequestQueryParameter struct {
	CategoryId int64  `json:"categoryId,omitempty"`
	Search     string `json:"search,omitempty"`
}

type GoodsCategorySkuspecGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type GoodsCategorySkuspecGetListResponsePageList struct {
	SpecId   int64  `json:"specId,omitempty"`
	SpecName string `json:"specName,omitempty"`
	Sort     int64  `json:"sort,omitempty"`
}
