package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FeicanStaffQueryStaffByStaffId
// @id 1319
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询员工详情)
func (client *Client) FeicanStaffQueryStaffByStaffId(request *FeicanStaffQueryStaffByStaffIdRequest) (response *FeicanStaffQueryStaffByStaffIdResponse, err error) {
	rpcResponse := CreateFeicanStaffQueryStaffByStaffIdResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FeicanStaffQueryStaffByStaffIdRequest struct {
	*api.BaseRequest

	StaffId int64 `json:"staffId,omitempty"`
}

type FeicanStaffQueryStaffByStaffIdResponse struct {
}

func CreateFeicanStaffQueryStaffByStaffIdRequest() (request *FeicanStaffQueryStaffByStaffIdRequest) {
	request = &FeicanStaffQueryStaffByStaffIdRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "feicanStaff/queryStaffByStaffId", "api")
	request.Method = api.POST
	return
}

func CreateFeicanStaffQueryStaffByStaffIdResponse() (response *api.BaseResponse[FeicanStaffQueryStaffByStaffIdResponse]) {
	response = api.CreateResponse[FeicanStaffQueryStaffByStaffIdResponse](&FeicanStaffQueryStaffByStaffIdResponse{})
	return
}
