package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FeicanStaffDeleteStaffByStaffId
// @id 1323
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除员工)
func (client *Client) FeicanStaffDeleteStaffByStaffId(request *FeicanStaffDeleteStaffByStaffIdRequest) (response *FeicanStaffDeleteStaffByStaffIdResponse, err error) {
	rpcResponse := CreateFeicanStaffDeleteStaffByStaffIdResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FeicanStaffDeleteStaffByStaffIdRequest struct {
	*api.BaseRequest

	StaffId int64 `json:"staffId,omitempty"`
}

type FeicanStaffDeleteStaffByStaffIdResponse struct {
}

func CreateFeicanStaffDeleteStaffByStaffIdRequest() (request *FeicanStaffDeleteStaffByStaffIdRequest) {
	request = &FeicanStaffDeleteStaffByStaffIdRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "feicanStaff/deleteStaffByStaffId", "api")
	request.Method = api.POST
	return
}

func CreateFeicanStaffDeleteStaffByStaffIdResponse() (response *api.BaseResponse[FeicanStaffDeleteStaffByStaffIdResponse]) {
	response = api.CreateResponse[FeicanStaffDeleteStaffByStaffIdResponse](&FeicanStaffDeleteStaffByStaffIdResponse{})
	return
}
