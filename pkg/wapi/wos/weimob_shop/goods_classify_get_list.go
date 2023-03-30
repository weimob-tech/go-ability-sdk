package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsClassifyGetList
// @id 1645
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1645?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取商品分组列表)
func (client *Client) GoodsClassifyGetList(request *GoodsClassifyGetListRequest) (response *GoodsClassifyGetListResponse, err error) {
	rpcResponse := CreateGoodsClassifyGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsClassifyGetListRequest struct {
	*api.BaseRequest

	BasicInfo             GoodsClassifyGetListRequestBasicInfo `json:"basicInfo,omitempty"`
	PageNum               int64                                `json:"pageNum,omitempty"`
	PageSize              int64                                `json:"pageSize,omitempty"`
	ParentId              int64                                `json:"parentId,omitempty"`
	IsSystemClassifyShow  bool                                 `json:"isSystemClassifyShow,omitempty"`
	HiddenSysClassifyShow bool                                 `json:"hiddenSysClassifyShow,omitempty"`
}

type GoodsClassifyGetListResponse struct {
	GoodsClassifyList  []GoodsClassifyGetListResponseGoodsClassifyList  `json:"goodsClassifyList,omitempty"`
	SystemClassifyList []GoodsClassifyGetListResponseSystemClassifyList `json:"systemClassifyList,omitempty"`
	PageSize           int64                                            `json:"pageSize,omitempty"`
	TotalCount         int64                                            `json:"totalCount,omitempty"`
	PageNum            int64                                            `json:"pageNum,omitempty"`
}

func CreateGoodsClassifyGetListRequest() (request *GoodsClassifyGetListRequest) {
	request = &GoodsClassifyGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/classify/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsClassifyGetListResponse() (response *api.BaseResponse[GoodsClassifyGetListResponse]) {
	response = api.CreateResponse[GoodsClassifyGetListResponse](&GoodsClassifyGetListResponse{})
	return
}

type GoodsClassifyGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type GoodsClassifyGetListResponseGoodsClassifyList struct {
	ClassifyLevel int64  `json:"classifyLevel,omitempty"`
	ClassifyId    int64  `json:"classifyId,omitempty"`
	Name          string `json:"name,omitempty"`
	IsLeaf        bool   `json:"isLeaf,omitempty"`
	ParentId      int64  `json:"parentId,omitempty"`
	ImageUrl      string `json:"imageUrl,omitempty"`
	IsShow        bool   `json:"isShow,omitempty"`
	IsHot         bool   `json:"isHot,omitempty"`
	Sort          int64  `json:"sort,omitempty"`
}

type GoodsClassifyGetListResponseSystemClassifyList struct {
	ClassifyLevel int64  `json:"classifyLevel,omitempty"`
	ClassifyId    int64  `json:"classifyId,omitempty"`
	Name          string `json:"name,omitempty"`
}
