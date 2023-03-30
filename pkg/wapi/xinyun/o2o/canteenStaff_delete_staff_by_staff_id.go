package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CanteenStaffDeleteStaffByStaffId
// @id 1317
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除餐厅员工)
func (client *Client) CanteenStaffDeleteStaffByStaffId(request *CanteenStaffDeleteStaffByStaffIdRequest) (response *CanteenStaffDeleteStaffByStaffIdResponse, err error) {
	rpcResponse := CreateCanteenStaffDeleteStaffByStaffIdResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CanteenStaffDeleteStaffByStaffIdRequest struct {
	*api.BaseRequest

	StaffId int64 `json:"staffId,omitempty"`
}

type CanteenStaffDeleteStaffByStaffIdResponse struct {
}

func CreateCanteenStaffDeleteStaffByStaffIdRequest() (request *CanteenStaffDeleteStaffByStaffIdRequest) {
	request = &CanteenStaffDeleteStaffByStaffIdRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "canteenStaff/deleteStaffByStaffId", "api")
	request.Method = api.POST
	return
}

func CreateCanteenStaffDeleteStaffByStaffIdResponse() (response *api.BaseResponse[CanteenStaffDeleteStaffByStaffIdResponse]) {
	response = api.CreateResponse[CanteenStaffDeleteStaffByStaffIdResponse](&CanteenStaffDeleteStaffByStaffIdResponse{})
	return
}
