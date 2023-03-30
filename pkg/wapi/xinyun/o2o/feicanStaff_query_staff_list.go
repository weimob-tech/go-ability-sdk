package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FeicanStaffQueryStaffList
// @id 1318
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询员工列表)
func (client *Client) FeicanStaffQueryStaffList(request *FeicanStaffQueryStaffListRequest) (response *FeicanStaffQueryStaffListResponse, err error) {
	rpcResponse := CreateFeicanStaffQueryStaffListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FeicanStaffQueryStaffListRequest struct {
	*api.BaseRequest

	Page     FeicanStaffQueryStaffListRequestPage `json:"page,omitempty"`
	Identity int64                                `json:"identity,omitempty"`
	Status   int64                                `json:"status,omitempty"`
	StoreId  int64                                `json:"storeId,omitempty"`
}

type FeicanStaffQueryStaffListResponse struct {
}

func CreateFeicanStaffQueryStaffListRequest() (request *FeicanStaffQueryStaffListRequest) {
	request = &FeicanStaffQueryStaffListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "feicanStaff/queryStaffList", "api")
	request.Method = api.POST
	return
}

func CreateFeicanStaffQueryStaffListResponse() (response *api.BaseResponse[FeicanStaffQueryStaffListResponse]) {
	response = api.CreateResponse[FeicanStaffQueryStaffListResponse](&FeicanStaffQueryStaffListResponse{})
	return
}

type FeicanStaffQueryStaffListRequestPage struct {
	PageIndex int64 `json:"pageIndex,omitempty"`
	PageSize  int64 `json:"pageSize,omitempty"`
}
