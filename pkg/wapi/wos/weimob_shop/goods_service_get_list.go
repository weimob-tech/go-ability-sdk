package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsServiceGetList
// @id 1654
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1654?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取商品服务列表)
func (client *Client) GoodsServiceGetList(request *GoodsServiceGetListRequest) (response *GoodsServiceGetListResponse, err error) {
	rpcResponse := CreateGoodsServiceGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsServiceGetListRequest struct {
	*api.BaseRequest

	QueryParameter GoodsServiceGetListRequestQueryParameter `json:"queryParameter,omitempty"`
	BasicInfo      GoodsServiceGetListRequestBasicInfo      `json:"basicInfo,omitempty"`
	PageNum        int64                                    `json:"pageNum,omitempty"`
	PageSize       int64                                    `json:"pageSize,omitempty"`
}

type GoodsServiceGetListResponse struct {
	PageList   []GoodsServiceGetListResponsePageList `json:"pageList,omitempty"`
	PageSize   int64                                 `json:"pageSize,omitempty"`
	TotalCount int64                                 `json:"totalCount,omitempty"`
	PageNum    int64                                 `json:"pageNum,omitempty"`
}

func CreateGoodsServiceGetListRequest() (request *GoodsServiceGetListRequest) {
	request = &GoodsServiceGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/service/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsServiceGetListResponse() (response *api.BaseResponse[GoodsServiceGetListResponse]) {
	response = api.CreateResponse[GoodsServiceGetListResponse](&GoodsServiceGetListResponse{})
	return
}

type GoodsServiceGetListRequestQueryParameter struct {
	Search string `json:"search,omitempty"`
}

type GoodsServiceGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type GoodsServiceGetListResponsePageList struct {
	Sort        int64  `json:"sort,omitempty"`
	ServiceId   int64  `json:"serviceId,omitempty"`
	ServiceName string `json:"serviceName,omitempty"`
}
