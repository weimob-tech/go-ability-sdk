package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// InsuranceSave
// @id 1025
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增/编辑保险产品)
func (client *Client) InsuranceSave(request *InsuranceSaveRequest) (response *InsuranceSaveResponse, err error) {
	rpcResponse := CreateInsuranceSaveResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type InsuranceSaveRequest struct {
	*api.BaseRequest

	InsuranceCode string  `json:"insuranceCode,omitempty"`
	InsureHints   string  `json:"insureHints,omitempty"`
	Name          string  `json:"name,omitempty"`
	Price         float64 `json:"price,omitempty"`
	Remark        string  `json:"remark,omitempty"`
	TypeIds       int     `json:"typeIds,omitempty"`
}

type InsuranceSaveResponse struct {
}

func CreateInsuranceSaveRequest() (request *InsuranceSaveRequest) {
	request = &InsuranceSaveRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "insurance/save", "api")
	request.Method = api.POST
	return
}

func CreateInsuranceSaveResponse() (response *api.BaseResponse[InsuranceSaveResponse]) {
	response = api.CreateResponse[InsuranceSaveResponse](&InsuranceSaveResponse{})
	return
}
