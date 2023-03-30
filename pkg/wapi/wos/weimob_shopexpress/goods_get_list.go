package weimob_shopexpress

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsGetList
// @id 1960
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1960?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商品列表)
func (client *Client) GoodsGetList(request *GoodsGetListRequest) (response *GoodsGetListResponse, err error) {
	rpcResponse := CreateGoodsGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsGetListRequest struct {
	*api.BaseRequest

	GoodsCodeList []int64 `json:"goodsCodeList,omitempty"`
	GoodsName     string  `json:"goodsName,omitempty"`
	PageIndex     int64   `json:"pageIndex,omitempty"`
	PageSize      int64   `json:"pageSize,omitempty"`
	Status        int64   `json:"status,omitempty"`
}

type GoodsGetListResponse struct {
	Items      []GoodsGetListResponseItems `json:"items,omitempty"`
	TotalCount int64                       `json:"totalCount,omitempty"`
}

func CreateGoodsGetListRequest() (request *GoodsGetListRequest) {
	request = &GoodsGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shopexpress", "v2.0", "goods/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsGetListResponse() (response *api.BaseResponse[GoodsGetListResponse]) {
	response = api.CreateResponse[GoodsGetListResponse](&GoodsGetListResponse{})
	return
}

type GoodsGetListResponseItems struct {
	GoodsCode    string `json:"goodsCode,omitempty"`
	GoodsName    string `json:"goodsName,omitempty"`
	GoodsStatus  int64  `json:"goodsStatus,omitempty"`
	Price        string `json:"price,omitempty"`
	CurrentStock int64  `json:"currentStock,omitempty"`
	SortNum      int64  `json:"sortNum,omitempty"`
	ImgUrl       string `json:"imgUrl,omitempty"`
	IfSingleSpec bool   `json:"ifSingleSpec,omitempty"`
}
