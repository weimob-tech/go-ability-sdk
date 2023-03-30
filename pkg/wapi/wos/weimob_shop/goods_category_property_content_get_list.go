package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsCategoryPropertyContentGetList
// @id 1660
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1660?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询参数值列表)
func (client *Client) GoodsCategoryPropertyContentGetList(request *GoodsCategoryPropertyContentGetListRequest) (response *GoodsCategoryPropertyContentGetListResponse, err error) {
	rpcResponse := CreateGoodsCategoryPropertyContentGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsCategoryPropertyContentGetListRequest struct {
	*api.BaseRequest

	QueryParameter GoodsCategoryPropertyContentGetListRequestQueryParameter `json:"queryParameter,omitempty"`
	BasicInfo      GoodsCategoryPropertyContentGetListRequestBasicInfo      `json:"basicInfo,omitempty"`
	PageNum        int64                                                    `json:"pageNum,omitempty"`
	PageSize       int64                                                    `json:"pageSize,omitempty"`
}

type GoodsCategoryPropertyContentGetListResponse struct {
	PageList   []GoodsCategoryPropertyContentGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                                 `json:"pageNum,omitempty"`
	PageSize   int64                                                 `json:"pageSize,omitempty"`
	TotalCount int64                                                 `json:"totalCount,omitempty"`
}

func CreateGoodsCategoryPropertyContentGetListRequest() (request *GoodsCategoryPropertyContentGetListRequest) {
	request = &GoodsCategoryPropertyContentGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/category/property/content/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsCategoryPropertyContentGetListResponse() (response *api.BaseResponse[GoodsCategoryPropertyContentGetListResponse]) {
	response = api.CreateResponse[GoodsCategoryPropertyContentGetListResponse](&GoodsCategoryPropertyContentGetListResponse{})
	return
}

type GoodsCategoryPropertyContentGetListRequestQueryParameter struct {
	CategoryPropId int64  `json:"categoryPropId,omitempty"`
	Search         string `json:"search,omitempty"`
}

type GoodsCategoryPropertyContentGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type GoodsCategoryPropertyContentGetListResponsePageList struct {
	CategoryId            int64  `json:"categoryId,omitempty"`
	CategoryPropId        int64  `json:"categoryPropId,omitempty"`
	CategoryPropValueId   int64  `json:"categoryPropValueId,omitempty"`
	CategoryPropValueName string `json:"categoryPropValueName,omitempty"`
}
