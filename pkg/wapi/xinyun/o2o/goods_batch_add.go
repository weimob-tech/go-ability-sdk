package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsBatchAdd
// @id 1459
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量新增菜品)
func (client *Client) GoodsBatchAdd(request *GoodsBatchAddRequest) (response *GoodsBatchAddResponse, err error) {
	rpcResponse := CreateGoodsBatchAddResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsBatchAddRequest struct {
	*api.BaseRequest
}

type GoodsBatchAddResponse struct {
}

func CreateGoodsBatchAddRequest() (request *GoodsBatchAddRequest) {
	request = &GoodsBatchAddRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goods/batchAdd", "api")
	request.Method = api.POST
	return
}

func CreateGoodsBatchAddResponse() (response *api.BaseResponse[GoodsBatchAddResponse]) {
	response = api.CreateResponse[GoodsBatchAddResponse](&GoodsBatchAddResponse{})
	return
}
