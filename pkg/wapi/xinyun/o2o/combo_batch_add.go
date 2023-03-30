package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ComboBatchAdd
// @id 1461
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量新增套餐)
func (client *Client) ComboBatchAdd(request *ComboBatchAddRequest) (response *ComboBatchAddResponse, err error) {
	rpcResponse := CreateComboBatchAddResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ComboBatchAddRequest struct {
	*api.BaseRequest
}

type ComboBatchAddResponse struct {
}

func CreateComboBatchAddRequest() (request *ComboBatchAddRequest) {
	request = &ComboBatchAddRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "combo/batchAdd", "api")
	request.Method = api.POST
	return
}

func CreateComboBatchAddResponse() (response *api.BaseResponse[ComboBatchAddResponse]) {
	response = api.CreateResponse[ComboBatchAddResponse](&ComboBatchAddResponse{})
	return
}
