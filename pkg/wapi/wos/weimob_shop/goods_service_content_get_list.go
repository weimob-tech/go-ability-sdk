package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsServiceContentGetList
// @id 1655
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1655?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询服务值列表)
func (client *Client) GoodsServiceContentGetList(request *GoodsServiceContentGetListRequest) (response *GoodsServiceContentGetListResponse, err error) {
	rpcResponse := CreateGoodsServiceContentGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsServiceContentGetListRequest struct {
	*api.BaseRequest

	QueryParameter GoodsServiceContentGetListRequestQueryParameter `json:"queryParameter,omitempty"`
	BasicInfo      GoodsServiceContentGetListRequestBasicInfo      `json:"basicInfo,omitempty"`
	PageNum        int64                                           `json:"pageNum,omitempty"`
	PageSize       int64                                           `json:"pageSize,omitempty"`
}

type GoodsServiceContentGetListResponse struct {
	PageList   []GoodsServiceContentGetListResponsePageList `json:"pageList,omitempty"`
	PageSize   int64                                        `json:"pageSize,omitempty"`
	TotalCount int64                                        `json:"totalCount,omitempty"`
	PageNum    int64                                        `json:"pageNum,omitempty"`
}

func CreateGoodsServiceContentGetListRequest() (request *GoodsServiceContentGetListRequest) {
	request = &GoodsServiceContentGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/service/content/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsServiceContentGetListResponse() (response *api.BaseResponse[GoodsServiceContentGetListResponse]) {
	response = api.CreateResponse[GoodsServiceContentGetListResponse](&GoodsServiceContentGetListResponse{})
	return
}

type GoodsServiceContentGetListRequestQueryParameter struct {
	ServiceId int64  `json:"serviceId,omitempty"`
	Search    string `json:"search,omitempty"`
}

type GoodsServiceContentGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type GoodsServiceContentGetListResponsePageList struct {
	ServiceValueName string `json:"serviceValueName,omitempty"`
	ServiceValueId   int64  `json:"serviceValueId,omitempty"`
	ServiceId        int64  `json:"serviceId,omitempty"`
}
