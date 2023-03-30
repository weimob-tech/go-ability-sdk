package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FeicanStaffUpdateStaffStatusByStaffId
// @id 1320
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新员工状态)
func (client *Client) FeicanStaffUpdateStaffStatusByStaffId(request *FeicanStaffUpdateStaffStatusByStaffIdRequest) (response *FeicanStaffUpdateStaffStatusByStaffIdResponse, err error) {
	rpcResponse := CreateFeicanStaffUpdateStaffStatusByStaffIdResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FeicanStaffUpdateStaffStatusByStaffIdRequest struct {
	*api.BaseRequest

	StaffId int64 `json:"staffId,omitempty"`
	Status  int64 `json:"status,omitempty"`
}

type FeicanStaffUpdateStaffStatusByStaffIdResponse struct {
}

func CreateFeicanStaffUpdateStaffStatusByStaffIdRequest() (request *FeicanStaffUpdateStaffStatusByStaffIdRequest) {
	request = &FeicanStaffUpdateStaffStatusByStaffIdRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "feicanStaff/updateStaffStatusByStaffId", "api")
	request.Method = api.POST
	return
}

func CreateFeicanStaffUpdateStaffStatusByStaffIdResponse() (response *api.BaseResponse[FeicanStaffUpdateStaffStatusByStaffIdResponse]) {
	response = api.CreateResponse[FeicanStaffUpdateStaffStatusByStaffIdResponse](&FeicanStaffUpdateStaffStatusByStaffIdResponse{})
	return
}
