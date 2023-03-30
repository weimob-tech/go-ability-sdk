package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// InsuranceTypeSave
// @id 1023
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增/编辑保险险种)
func (client *Client) InsuranceTypeSave(request *InsuranceTypeSaveRequest) (response *InsuranceTypeSaveResponse, err error) {
	rpcResponse := CreateInsuranceTypeSaveResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type InsuranceTypeSaveRequest struct {
	*api.BaseRequest

	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type InsuranceTypeSaveResponse struct {
}

func CreateInsuranceTypeSaveRequest() (request *InsuranceTypeSaveRequest) {
	request = &InsuranceTypeSaveRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "insuranceType/save", "api")
	request.Method = api.POST
	return
}

func CreateInsuranceTypeSaveResponse() (response *api.BaseResponse[InsuranceTypeSaveResponse]) {
	response = api.CreateResponse[InsuranceTypeSaveResponse](&InsuranceTypeSaveResponse{})
	return
}
