package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// EmployeeRelationGet
// @id 2082
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2082?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询员工的组织关系)
func (client *Client) EmployeeRelationGet(request *EmployeeRelationGetRequest) (response *EmployeeRelationGetResponse, err error) {
	rpcResponse := CreateEmployeeRelationGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type EmployeeRelationGetRequest struct {
	*api.BaseRequest

	Wid      int64 `json:"wid,omitempty"`
	PageNum  int64 `json:"pageNum,omitempty"`
	PageSize int64 `json:"pageSize,omitempty"`
}

type EmployeeRelationGetResponse struct {
	Data      []EmployeeRelationGetResponseData `json:"data,omitempty"`
	PageSize  int64                             `json:"pageSize,omitempty"`
	PageNum   int64                             `json:"pageNum,omitempty"`
	TotalSize int64                             `json:"totalSize,omitempty"`
}

func CreateEmployeeRelationGetRequest() (request *EmployeeRelationGetRequest) {
	request = &EmployeeRelationGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "employee/relation/get", "apigw")
	request.Method = api.POST
	return
}

func CreateEmployeeRelationGetResponse() (response *api.BaseResponse[EmployeeRelationGetResponse]) {
	response = api.CreateResponse[EmployeeRelationGetResponse](&EmployeeRelationGetResponse{})
	return
}

type EmployeeRelationGetResponseData struct {
	VidType int64  `json:"vidType,omitempty"`
	VidName string `json:"vidName,omitempty"`
	Vid     int64  `json:"vid,omitempty"`
}
