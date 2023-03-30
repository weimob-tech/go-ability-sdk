package weimob_shopexpress

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsTagGetList
// @id 1968
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1968?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商品标签列表)
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

	PageIndex int64  `json:"pageIndex,omitempty"`
	PageSize  int64  `json:"pageSize,omitempty"`
	TagCode   string `json:"tagCode,omitempty"`
	TagName   string `json:"tagName,omitempty"`
}

type GoodsTagGetListResponse struct {
	Items      []GoodsTagGetListResponseItems `json:"items,omitempty"`
	TotalCount int64                          `json:"totalCount,omitempty"`
}

func CreateGoodsTagGetListRequest() (request *GoodsTagGetListRequest) {
	request = &GoodsTagGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shopexpress", "v2.0", "goods/tag/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsTagGetListResponse() (response *api.BaseResponse[GoodsTagGetListResponse]) {
	response = api.CreateResponse[GoodsTagGetListResponse](&GoodsTagGetListResponse{})
	return
}

type GoodsTagGetListResponseItems struct {
	TagCode string `json:"tagCode,omitempty"`
	TagName string `json:"tagName,omitempty"`
}
