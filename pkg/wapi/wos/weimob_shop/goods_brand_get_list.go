package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsBrandGetList
// @id 1830
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1830?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取商品品牌列表)
func (client *Client) GoodsBrandGetList(request *GoodsBrandGetListRequest) (response *GoodsBrandGetListResponse, err error) {
	rpcResponse := CreateGoodsBrandGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsBrandGetListRequest struct {
	*api.BaseRequest

	BasicInfo      GoodsBrandGetListRequestBasicInfo      `json:"basicInfo,omitempty"`
	QueryParameter GoodsBrandGetListRequestQueryParameter `json:"queryParameter,omitempty"`
	PageNum        int64                                  `json:"pageNum,omitempty"`
	PageSize       int64                                  `json:"pageSize,omitempty"`
}

type GoodsBrandGetListResponse struct {
	PageList   []GoodsBrandGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                               `json:"pageNum,omitempty"`
	PageSize   int64                               `json:"pageSize,omitempty"`
	TotalCount int64                               `json:"totalCount,omitempty"`
}

func CreateGoodsBrandGetListRequest() (request *GoodsBrandGetListRequest) {
	request = &GoodsBrandGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/brand/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsBrandGetListResponse() (response *api.BaseResponse[GoodsBrandGetListResponse]) {
	response = api.CreateResponse[GoodsBrandGetListResponse](&GoodsBrandGetListResponse{})
	return
}

type GoodsBrandGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type GoodsBrandGetListRequestQueryParameter struct {
	Title string `json:"title,omitempty"`
}

type GoodsBrandGetListResponsePageList struct {
	Id                  int64  `json:"id,omitempty"`
	BrandManagementType int64  `json:"brandManagementType,omitempty"`
	BrandAuditType      int64  `json:"brandAuditType,omitempty"`
	Title               string `json:"title,omitempty"`
}
