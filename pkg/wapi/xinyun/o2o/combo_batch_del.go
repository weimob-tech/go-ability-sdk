package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ComboBatchDel
// @id 1462
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量删除套餐)
func (client *Client) ComboBatchDel(request *ComboBatchDelRequest) (response *ComboBatchDelResponse, err error) {
	rpcResponse := CreateComboBatchDelResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ComboBatchDelRequest struct {
	*api.BaseRequest
}

type ComboBatchDelResponse struct {
}

func CreateComboBatchDelRequest() (request *ComboBatchDelRequest) {
	request = &ComboBatchDelRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "combo/batchDel", "api")
	request.Method = api.POST
	return
}

func CreateComboBatchDelResponse() (response *api.BaseResponse[ComboBatchDelResponse]) {
	response = api.CreateResponse[ComboBatchDelResponse](&ComboBatchDelResponse{})
	return
}
