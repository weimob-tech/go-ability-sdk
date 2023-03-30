package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CanteenStaffQueryStaffList
// @id 1312
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询餐厅员工列表)
func (client *Client) CanteenStaffQueryStaffList(request *CanteenStaffQueryStaffListRequest) (response *CanteenStaffQueryStaffListResponse, err error) {
	rpcResponse := CreateCanteenStaffQueryStaffListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CanteenStaffQueryStaffListRequest struct {
	*api.BaseRequest

	Page    CanteenStaffQueryStaffListRequestPage `json:"page,omitempty"`
	Status  int                                   `json:"status,omitempty"`
	StoreId int64                                 `json:"storeId,omitempty"`
}

type CanteenStaffQueryStaffListResponse struct {
}

func CreateCanteenStaffQueryStaffListRequest() (request *CanteenStaffQueryStaffListRequest) {
	request = &CanteenStaffQueryStaffListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "canteenStaff/queryStaffList", "api")
	request.Method = api.POST
	return
}

func CreateCanteenStaffQueryStaffListResponse() (response *api.BaseResponse[CanteenStaffQueryStaffListResponse]) {
	response = api.CreateResponse[CanteenStaffQueryStaffListResponse](&CanteenStaffQueryStaffListResponse{})
	return
}

type CanteenStaffQueryStaffListRequestPage struct {
	PageIndex int `json:"pageIndex,omitempty"`
	PageSize  int `json:"pageSize,omitempty"`
}
