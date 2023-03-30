package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CanteenStaffUpdateStaffStatusByStaffId
// @id 1314
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新餐厅员工状态)
func (client *Client) CanteenStaffUpdateStaffStatusByStaffId(request *CanteenStaffUpdateStaffStatusByStaffIdRequest) (response *CanteenStaffUpdateStaffStatusByStaffIdResponse, err error) {
	rpcResponse := CreateCanteenStaffUpdateStaffStatusByStaffIdResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CanteenStaffUpdateStaffStatusByStaffIdRequest struct {
	*api.BaseRequest

	StaffId int64 `json:"staffId,omitempty"`
	Status  int   `json:"status,omitempty"`
}

type CanteenStaffUpdateStaffStatusByStaffIdResponse struct {
}

func CreateCanteenStaffUpdateStaffStatusByStaffIdRequest() (request *CanteenStaffUpdateStaffStatusByStaffIdRequest) {
	request = &CanteenStaffUpdateStaffStatusByStaffIdRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "canteenStaff/updateStaffStatusByStaffId", "api")
	request.Method = api.POST
	return
}

func CreateCanteenStaffUpdateStaffStatusByStaffIdResponse() (response *api.BaseResponse[CanteenStaffUpdateStaffStatusByStaffIdResponse]) {
	response = api.CreateResponse[CanteenStaffUpdateStaffStatusByStaffIdResponse](&CanteenStaffUpdateStaffStatusByStaffIdResponse{})
	return
}
