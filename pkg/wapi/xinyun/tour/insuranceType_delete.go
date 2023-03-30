package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// InsuranceTypeDelete
// @id 1022
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除保险险种)
func (client *Client) InsuranceTypeDelete(request *InsuranceTypeDeleteRequest) (response *InsuranceTypeDeleteResponse, err error) {
	rpcResponse := CreateInsuranceTypeDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type InsuranceTypeDeleteRequest struct {
	*api.BaseRequest

	Id int `json:"id,omitempty"`
}

type InsuranceTypeDeleteResponse struct {
}

func CreateInsuranceTypeDeleteRequest() (request *InsuranceTypeDeleteRequest) {
	request = &InsuranceTypeDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "insuranceType/delete", "api")
	request.Method = api.POST
	return
}

func CreateInsuranceTypeDeleteResponse() (response *api.BaseResponse[InsuranceTypeDeleteResponse]) {
	response = api.CreateResponse[InsuranceTypeDeleteResponse](&InsuranceTypeDeleteResponse{})
	return
}
