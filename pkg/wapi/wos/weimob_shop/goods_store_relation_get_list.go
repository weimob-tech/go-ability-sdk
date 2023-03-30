package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsStoreRelationGetList
// @id 1526
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1526?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询商品关联门店ID列表)
func (client *Client) GoodsStoreRelationGetList(request *GoodsStoreRelationGetListRequest) (response *GoodsStoreRelationGetListResponse, err error) {
	rpcResponse := CreateGoodsStoreRelationGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsStoreRelationGetListRequest struct {
	*api.BaseRequest

	BasicInfo GoodsStoreRelationGetListRequestBasicInfo `json:"basicInfo,omitempty"`
	GoodsId   int64                                     `json:"goodsId,omitempty"`
}

type GoodsStoreRelationGetListResponse struct {
	VidList []int64 `json:"vidList,omitempty"`
}

func CreateGoodsStoreRelationGetListRequest() (request *GoodsStoreRelationGetListRequest) {
	request = &GoodsStoreRelationGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/store/relation/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsStoreRelationGetListResponse() (response *api.BaseResponse[GoodsStoreRelationGetListResponse]) {
	response = api.CreateResponse[GoodsStoreRelationGetListResponse](&GoodsStoreRelationGetListResponse{})
	return
}

type GoodsStoreRelationGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
