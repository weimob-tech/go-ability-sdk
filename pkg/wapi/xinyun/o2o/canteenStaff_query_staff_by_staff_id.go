package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CanteenStaffQueryStaffByStaffId
// @id 1313
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询餐厅员工详情)
func (client *Client) CanteenStaffQueryStaffByStaffId(request *CanteenStaffQueryStaffByStaffIdRequest) (response *CanteenStaffQueryStaffByStaffIdResponse, err error) {
	rpcResponse := CreateCanteenStaffQueryStaffByStaffIdResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CanteenStaffQueryStaffByStaffIdRequest struct {
	*api.BaseRequest

	StaffId int64 `json:"staffId,omitempty"`
}

type CanteenStaffQueryStaffByStaffIdResponse struct {
}

func CreateCanteenStaffQueryStaffByStaffIdRequest() (request *CanteenStaffQueryStaffByStaffIdRequest) {
	request = &CanteenStaffQueryStaffByStaffIdRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "canteenStaff/queryStaffByStaffId", "api")
	request.Method = api.POST
	return
}

func CreateCanteenStaffQueryStaffByStaffIdResponse() (response *api.BaseResponse[CanteenStaffQueryStaffByStaffIdResponse]) {
	response = api.CreateResponse[CanteenStaffQueryStaffByStaffIdResponse](&CanteenStaffQueryStaffByStaffIdResponse{})
	return
}
