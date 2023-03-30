package weimob_shopexpress

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsGroupGetList
// @id 1971
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1971?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商品分组列表)
func (client *Client) GoodsGroupGetList(request *GoodsGroupGetListRequest) (response *GoodsGroupGetListResponse, err error) {
	rpcResponse := CreateGoodsGroupGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsGroupGetListRequest struct {
	*api.BaseRequest

	GroupType string `json:"groupType,omitempty"`
	Name      string `json:"name,omitempty"`
	PageIndex int64  `json:"pageIndex,omitempty"`
	PageSize  int64  `json:"pageSize,omitempty"`
}

type GoodsGroupGetListResponse struct {
	Items      []GoodsGroupGetListResponseItems `json:"items,omitempty"`
	TotalCount int64                            `json:"totalCount,omitempty"`
}

func CreateGoodsGroupGetListRequest() (request *GoodsGroupGetListRequest) {
	request = &GoodsGroupGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shopexpress", "v2.0", "goods/group/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsGroupGetListResponse() (response *api.BaseResponse[GoodsGroupGetListResponse]) {
	response = api.CreateResponse[GoodsGroupGetListResponse](&GoodsGroupGetListResponse{})
	return
}

type GoodsGroupGetListResponseItems struct {
	GroupCode  string `json:"groupCode,omitempty"`
	Name       string `json:"name,omitempty"`
	Sort       int64  `json:"sort,omitempty"`
	GroupType  int64  `json:"groupType,omitempty"`
	ShowStatus int64  `json:"showStatus,omitempty"`
}
