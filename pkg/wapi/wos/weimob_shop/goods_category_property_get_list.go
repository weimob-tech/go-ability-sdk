package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsCategoryPropertyGetList
// @id 1586
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1586?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取类目下的商品参数)
func (client *Client) GoodsCategoryPropertyGetList(request *GoodsCategoryPropertyGetListRequest) (response *GoodsCategoryPropertyGetListResponse, err error) {
	rpcResponse := CreateGoodsCategoryPropertyGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsCategoryPropertyGetListRequest struct {
	*api.BaseRequest

	QueryParameter GoodsCategoryPropertyGetListRequestQueryParameter `json:"queryParameter,omitempty"`
	BasicInfo      GoodsCategoryPropertyGetListRequestBasicInfo      `json:"basicInfo,omitempty"`
	PageNum        int64                                             `json:"pageNum,omitempty"`
	PageSize       int64                                             `json:"pageSize,omitempty"`
}

type GoodsCategoryPropertyGetListResponse struct {
	PageList   []GoodsCategoryPropertyGetListResponsePageList `json:"pageList,omitempty"`
	PageSize   int64                                          `json:"pageSize,omitempty"`
	TotalCount int64                                          `json:"totalCount,omitempty"`
	PageNum    int64                                          `json:"pageNum,omitempty"`
}

func CreateGoodsCategoryPropertyGetListRequest() (request *GoodsCategoryPropertyGetListRequest) {
	request = &GoodsCategoryPropertyGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/category/property/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsCategoryPropertyGetListResponse() (response *api.BaseResponse[GoodsCategoryPropertyGetListResponse]) {
	response = api.CreateResponse[GoodsCategoryPropertyGetListResponse](&GoodsCategoryPropertyGetListResponse{})
	return
}

type GoodsCategoryPropertyGetListRequestQueryParameter struct {
	Search     string `json:"search,omitempty"`
	CategoryId int64  `json:"categoryId,omitempty"`
}

type GoodsCategoryPropertyGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type GoodsCategoryPropertyGetListResponsePageList struct {
	CategoryId       int64  `json:"categoryId,omitempty"`
	CategoryPropId   int64  `json:"categoryPropId,omitempty"`
	CategoryPropName string `json:"categoryPropName,omitempty"`
	CategoryPropType int64  `json:"categoryPropType,omitempty"`
}
