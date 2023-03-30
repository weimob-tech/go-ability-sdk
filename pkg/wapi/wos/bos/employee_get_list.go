package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// EmployeeGetList
// @id 1365
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1365?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=分页查询员工列表)
func (client *Client) EmployeeGetList(request *EmployeeGetListRequest) (response *EmployeeGetListResponse, err error) {
	rpcResponse := CreateEmployeeGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type EmployeeGetListRequest struct {
	*api.BaseRequest

	PageNum  int64 `json:"pageNum,omitempty"`
	PageSize int64 `json:"pageSize,omitempty"`
}

type EmployeeGetListResponse struct {
	Data      []EmployeeGetListResponseData `json:"data,omitempty"`
	PageSize  int64                         `json:"pageSize,omitempty"`
	PageNum   int64                         `json:"pageNum,omitempty"`
	TotalSize int64                         `json:"totalSize,omitempty"`
}

func CreateEmployeeGetListRequest() (request *EmployeeGetListRequest) {
	request = &EmployeeGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "employee/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateEmployeeGetListResponse() (response *api.BaseResponse[EmployeeGetListResponse]) {
	response = api.CreateResponse[EmployeeGetListResponse](&EmployeeGetListResponse{})
	return
}

type EmployeeGetListResponseData struct {
	Phone      string `json:"phone,omitempty"`
	EmployeeNo string `json:"employeeNo,omitempty"`
	Name       string `json:"name,omitempty"`
	Status     int64  `json:"status,omitempty"`
	Wid        int64  `json:"wid,omitempty"`
}
