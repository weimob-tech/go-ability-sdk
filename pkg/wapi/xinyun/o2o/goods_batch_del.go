package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsBatchDel
// @id 1460
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量删除菜品)
func (client *Client) GoodsBatchDel(request *GoodsBatchDelRequest) (response *GoodsBatchDelResponse, err error) {
	rpcResponse := CreateGoodsBatchDelResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsBatchDelRequest struct {
	*api.BaseRequest
}

type GoodsBatchDelResponse struct {
}

func CreateGoodsBatchDelRequest() (request *GoodsBatchDelRequest) {
	request = &GoodsBatchDelRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goods/batchDel", "api")
	request.Method = api.POST
	return
}

func CreateGoodsBatchDelResponse() (response *api.BaseResponse[GoodsBatchDelResponse]) {
	response = api.CreateResponse[GoodsBatchDelResponse](&GoodsBatchDelResponse{})
	return
}
