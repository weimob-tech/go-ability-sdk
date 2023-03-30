package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsPropertyContentGetList
// @id 1636
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1636?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询商品属性值列表)
func (client *Client) GoodsPropertyContentGetList(request *GoodsPropertyContentGetListRequest) (response *GoodsPropertyContentGetListResponse, err error) {
	rpcResponse := CreateGoodsPropertyContentGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsPropertyContentGetListRequest struct {
	*api.BaseRequest

	QueryParameter GoodsPropertyContentGetListRequestQueryParameter `json:"queryParameter,omitempty"`
	BasicInfo      GoodsPropertyContentGetListRequestBasicInfo      `json:"basicInfo,omitempty"`
	PageNum        int64                                            `json:"pageNum,omitempty"`
	PageSize       int64                                            `json:"pageSize,omitempty"`
}

type GoodsPropertyContentGetListResponse struct {
	PageList   []GoodsPropertyContentGetListResponsePageList `json:"pageList,omitempty"`
	PageSize   int64                                         `json:"pageSize,omitempty"`
	TotalCount int64                                         `json:"totalCount,omitempty"`
	PageNum    int64                                         `json:"pageNum,omitempty"`
}

func CreateGoodsPropertyContentGetListRequest() (request *GoodsPropertyContentGetListRequest) {
	request = &GoodsPropertyContentGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/property/content/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsPropertyContentGetListResponse() (response *api.BaseResponse[GoodsPropertyContentGetListResponse]) {
	response = api.CreateResponse[GoodsPropertyContentGetListResponse](&GoodsPropertyContentGetListResponse{})
	return
}

type GoodsPropertyContentGetListRequestQueryParameter struct {
	PropId int64  `json:"propId,omitempty"`
	Search string `json:"search,omitempty"`
}

type GoodsPropertyContentGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type GoodsPropertyContentGetListResponsePageList struct {
	PropId        int64  `json:"propId,omitempty"`
	PropValueId   int64  `json:"propValueId,omitempty"`
	PropValueName string `json:"propValueName,omitempty"`
	Sort          int64  `json:"sort,omitempty"`
}
