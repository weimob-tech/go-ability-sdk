package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsUnitGetList
// @id 1627
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1627?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询单位列表)
func (client *Client) GoodsUnitGetList(request *GoodsUnitGetListRequest) (response *GoodsUnitGetListResponse, err error) {
	rpcResponse := CreateGoodsUnitGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsUnitGetListRequest struct {
	*api.BaseRequest

	BasicInfo GoodsUnitGetListRequestBasicInfo `json:"basicInfo,omitempty"`
	PageNum   int64                            `json:"pageNum,omitempty"`
	PageSize  int64                            `json:"pageSize,omitempty"`
}

type GoodsUnitGetListResponse struct {
	PageList   []GoodsUnitGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                              `json:"pageNum,omitempty"`
	PageSize   int64                              `json:"pageSize,omitempty"`
	TotalCount int64                              `json:"totalCount,omitempty"`
}

func CreateGoodsUnitGetListRequest() (request *GoodsUnitGetListRequest) {
	request = &GoodsUnitGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/unit/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsUnitGetListResponse() (response *api.BaseResponse[GoodsUnitGetListResponse]) {
	response = api.CreateResponse[GoodsUnitGetListResponse](&GoodsUnitGetListResponse{})
	return
}

type GoodsUnitGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type GoodsUnitGetListResponsePageList struct {
	UnitId    int64  `json:"unitId,omitempty"`
	UnitTitle string `json:"unitTitle,omitempty"`
}
