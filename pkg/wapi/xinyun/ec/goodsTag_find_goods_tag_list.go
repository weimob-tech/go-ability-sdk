package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsTagFindGoodsTagList
// @id 360
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询商品标签列表)
func (client *Client) GoodsTagFindGoodsTagList(request *GoodsTagFindGoodsTagListRequest) (response *GoodsTagFindGoodsTagListResponse, err error) {
	rpcResponse := CreateGoodsTagFindGoodsTagListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsTagFindGoodsTagListRequest struct {
	*api.BaseRequest
}

type GoodsTagFindGoodsTagListResponse struct {
}

func CreateGoodsTagFindGoodsTagListRequest() (request *GoodsTagFindGoodsTagListRequest) {
	request = &GoodsTagFindGoodsTagListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goodsTag/findGoodsTagList", "api")
	request.Method = api.POST
	return
}

func CreateGoodsTagFindGoodsTagListResponse() (response *api.BaseResponse[GoodsTagFindGoodsTagListResponse]) {
	response = api.CreateResponse[GoodsTagFindGoodsTagListResponse](&GoodsTagFindGoodsTagListResponse{})
	return
}
