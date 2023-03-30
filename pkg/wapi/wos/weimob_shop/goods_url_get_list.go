package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsUrlGetList
// @id 2033
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2033?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取商品分享链接)
func (client *Client) GoodsUrlGetList(request *GoodsUrlGetListRequest) (response *GoodsUrlGetListResponse, err error) {
	rpcResponse := CreateGoodsUrlGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsUrlGetListRequest struct {
	*api.BaseRequest

	BasicInfo    GoodsUrlGetListRequestBasicInfo `json:"basicInfo,omitempty"`
	GoodsIdList  []int64                         `json:"goodsIdList,omitempty"`
	IsQueryScene bool                            `json:"isQueryScene,omitempty"`
}

type GoodsUrlGetListResponse struct {
	GoodsUrlOutputList []GoodsUrlGetListResponseGoodsUrlOutputList `json:"goodsUrlOutputList,omitempty"`
}

func CreateGoodsUrlGetListRequest() (request *GoodsUrlGetListRequest) {
	request = &GoodsUrlGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/url/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsUrlGetListResponse() (response *api.BaseResponse[GoodsUrlGetListResponse]) {
	response = api.CreateResponse[GoodsUrlGetListResponse](&GoodsUrlGetListResponse{})
	return
}

type GoodsUrlGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type GoodsUrlGetListResponseGoodsUrlOutputList struct {
	GoodsId int64  `json:"goodsId,omitempty"`
	H5Url   string `json:"h5Url,omitempty"`
	MiniUrl string `json:"miniUrl,omitempty"`
	Scene   string `json:"scene,omitempty"`
}
