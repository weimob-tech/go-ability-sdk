package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsPropertyGetList
// @id 1632
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1632?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询商品属性列表)
func (client *Client) GoodsPropertyGetList(request *GoodsPropertyGetListRequest) (response *GoodsPropertyGetListResponse, err error) {
	rpcResponse := CreateGoodsPropertyGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsPropertyGetListRequest struct {
	*api.BaseRequest

	QueryParameter GoodsPropertyGetListRequestQueryParameter `json:"queryParameter,omitempty"`
	BasicInfo      GoodsPropertyGetListRequestBasicInfo      `json:"basicInfo,omitempty"`
	PageNum        int64                                     `json:"pageNum,omitempty"`
	PageSize       int64                                     `json:"pageSize,omitempty"`
}

type GoodsPropertyGetListResponse struct {
	PageList   []GoodsPropertyGetListResponsePageList `json:"pageList,omitempty"`
	PageSize   int64                                  `json:"pageSize,omitempty"`
	TotalCount int64                                  `json:"totalCount,omitempty"`
	PageNum    int64                                  `json:"pageNum,omitempty"`
}

func CreateGoodsPropertyGetListRequest() (request *GoodsPropertyGetListRequest) {
	request = &GoodsPropertyGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/property/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsPropertyGetListResponse() (response *api.BaseResponse[GoodsPropertyGetListResponse]) {
	response = api.CreateResponse[GoodsPropertyGetListResponse](&GoodsPropertyGetListResponse{})
	return
}

type GoodsPropertyGetListRequestQueryParameter struct {
	Search   string `json:"search,omitempty"`
	PropType int64  `json:"propType,omitempty"`
}

type GoodsPropertyGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type GoodsPropertyGetListResponsePageList struct {
	PropId       int64  `json:"propId,omitempty"`
	Sort         int64  `json:"sort,omitempty"`
	PropName     string `json:"propName,omitempty"`
	IsSearchable bool   `json:"isSearchable,omitempty"`
}
