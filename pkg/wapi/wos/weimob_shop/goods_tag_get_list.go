package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsTagGetList
// @id 1850
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1850?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询商品标签列表)
func (client *Client) GoodsTagGetList(request *GoodsTagGetListRequest) (response *GoodsTagGetListResponse, err error) {
	rpcResponse := CreateGoodsTagGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsTagGetListRequest struct {
	*api.BaseRequest

	QueryParameter GoodsTagGetListRequestQueryParameter `json:"queryParameter,omitempty"`
	BasicInfo      GoodsTagGetListRequestBasicInfo      `json:"basicInfo,omitempty"`
	PageNum        int64                                `json:"pageNum,omitempty"`
	PageSize       int64                                `json:"pageSize,omitempty"`
}

type GoodsTagGetListResponse struct {
	PageList   []GoodsTagGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                             `json:"pageNum,omitempty"`
	PageSize   int64                             `json:"pageSize,omitempty"`
	TotalCount int64                             `json:"totalCount,omitempty"`
}

func CreateGoodsTagGetListRequest() (request *GoodsTagGetListRequest) {
	request = &GoodsTagGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/tag/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsTagGetListResponse() (response *api.BaseResponse[GoodsTagGetListResponse]) {
	response = api.CreateResponse[GoodsTagGetListResponse](&GoodsTagGetListResponse{})
	return
}

type GoodsTagGetListRequestQueryParameter struct {
	Name string `json:"name,omitempty"`
}

type GoodsTagGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type GoodsTagGetListResponsePageList struct {
	TagId int64  `json:"tagId,omitempty"`
	Name  string `json:"name,omitempty"`
}
