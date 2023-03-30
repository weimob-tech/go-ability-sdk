package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// InsuranceDelete
// @id 1026
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除保险产品)
func (client *Client) InsuranceDelete(request *InsuranceDeleteRequest) (response *InsuranceDeleteResponse, err error) {
	rpcResponse := CreateInsuranceDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type InsuranceDeleteRequest struct {
	*api.BaseRequest

	InsuranceCode string `json:"insuranceCode,omitempty"`
}

type InsuranceDeleteResponse struct {
}

func CreateInsuranceDeleteRequest() (request *InsuranceDeleteRequest) {
	request = &InsuranceDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "insurance/delete", "api")
	request.Method = api.POST
	return
}

func CreateInsuranceDeleteResponse() (response *api.BaseResponse[InsuranceDeleteResponse]) {
	response = api.CreateResponse[InsuranceDeleteResponse](&InsuranceDeleteResponse{})
	return
}
