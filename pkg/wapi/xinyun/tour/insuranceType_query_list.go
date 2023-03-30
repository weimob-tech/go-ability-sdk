package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// InsuranceTypeQueryList
// @id 1024
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询保险险种列表)
func (client *Client) InsuranceTypeQueryList(request *InsuranceTypeQueryListRequest) (response *InsuranceTypeQueryListResponse, err error) {
	rpcResponse := CreateInsuranceTypeQueryListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type InsuranceTypeQueryListRequest struct {
	*api.BaseRequest
}

type InsuranceTypeQueryListResponse struct {
}

func CreateInsuranceTypeQueryListRequest() (request *InsuranceTypeQueryListRequest) {
	request = &InsuranceTypeQueryListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "insuranceType/queryList", "api")
	request.Method = api.POST
	return
}

func CreateInsuranceTypeQueryListResponse() (response *api.BaseResponse[InsuranceTypeQueryListResponse]) {
	response = api.CreateResponse[InsuranceTypeQueryListResponse](&InsuranceTypeQueryListResponse{})
	return
}
