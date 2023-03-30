package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsCategoryGetList
// @id 1584
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1584?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取一级/二级类目)
func (client *Client) GoodsCategoryGetList(request *GoodsCategoryGetListRequest) (response *GoodsCategoryGetListResponse, err error) {
	rpcResponse := CreateGoodsCategoryGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsCategoryGetListRequest struct {
	*api.BaseRequest

	BasicInfo        GoodsCategoryGetListRequestBasicInfo `json:"basicInfo,omitempty"`
	ParentCategoryId int64                                `json:"parentCategoryId,omitempty"`
}

type GoodsCategoryGetListResponse struct {
	GoodsCategoryInfoList []GoodsCategoryGetListResponseGoodsCategoryInfoList `json:"goodsCategoryInfoList,omitempty"`
}

func CreateGoodsCategoryGetListRequest() (request *GoodsCategoryGetListRequest) {
	request = &GoodsCategoryGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/category/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsCategoryGetListResponse() (response *api.BaseResponse[GoodsCategoryGetListResponse]) {
	response = api.CreateResponse[GoodsCategoryGetListResponse](&GoodsCategoryGetListResponse{})
	return
}

type GoodsCategoryGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type GoodsCategoryGetListResponseGoodsCategoryInfoList struct {
	CategoryId       int64  `json:"categoryId,omitempty"`
	CategoryLevel    int64  `json:"categoryLevel,omitempty"`
	CategoryName     string `json:"categoryName,omitempty"`
	IsLeaf           bool   `json:"isLeaf,omitempty"`
	ParentCategoryId int64  `json:"parentCategoryId,omitempty"`
}
