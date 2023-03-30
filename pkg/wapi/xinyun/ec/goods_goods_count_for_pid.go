package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsGoodsCountForPid
// @id 1445
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询商家商品总数与上架中商品数量)
func (client *Client) GoodsGoodsCountForPid(request *GoodsGoodsCountForPidRequest) (response *GoodsGoodsCountForPidResponse, err error) {
	rpcResponse := CreateGoodsGoodsCountForPidResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsGoodsCountForPidRequest struct {
	*api.BaseRequest
}

type GoodsGoodsCountForPidResponse struct {
}

func CreateGoodsGoodsCountForPidRequest() (request *GoodsGoodsCountForPidRequest) {
	request = &GoodsGoodsCountForPidRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goods/goodsCountForPid", "api")
	request.Method = api.POST
	return
}

func CreateGoodsGoodsCountForPidResponse() (response *api.BaseResponse[GoodsGoodsCountForPidResponse]) {
	response = api.CreateResponse[GoodsGoodsCountForPidResponse](&GoodsGoodsCountForPidResponse{})
	return
}
