package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsImageUpdate
// @id 2327
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2327?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新商品图片)
func (client *Client) GoodsImageUpdate(request *GoodsImageUpdateRequest) (response *GoodsImageUpdateResponse, err error) {
	rpcResponse := CreateGoodsImageUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsImageUpdateRequest struct {
	*api.BaseRequest

	ImageList []GoodsImageUpdateRequestImageList `json:"imageList,omitempty"`
	GoodsId   int64                              `json:"goodsId,omitempty"`
}

type GoodsImageUpdateResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateGoodsImageUpdateRequest() (request *GoodsImageUpdateRequest) {
	request = &GoodsImageUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/image/update", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsImageUpdateResponse() (response *api.BaseResponse[GoodsImageUpdateResponse]) {
	response = api.CreateResponse[GoodsImageUpdateResponse](&GoodsImageUpdateResponse{})
	return
}

type GoodsImageUpdateRequestImageList struct {
}
