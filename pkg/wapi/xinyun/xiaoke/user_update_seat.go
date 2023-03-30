package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserUpdateSeat
// @id 1636
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=开通/关闭员工销氪坐席)
func (client *Client) UserUpdateSeat(request *UserUpdateSeatRequest) (response *UserUpdateSeatResponse, err error) {
	rpcResponse := CreateUserUpdateSeatResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserUpdateSeatRequest struct {
	*api.BaseRequest

	WidList    []int `json:"widList,omitempty"`
	SeatStatus int   `json:"seatStatus,omitempty"`
}

type UserUpdateSeatResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateUserUpdateSeatRequest() (request *UserUpdateSeatRequest) {
	request = &UserUpdateSeatRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "user/updateSeat", "api")
	request.Method = api.POST
	return
}

func CreateUserUpdateSeatResponse() (response *api.BaseResponse[UserUpdateSeatResponse]) {
	response = api.CreateResponse[UserUpdateSeatResponse](&UserUpdateSeatResponse{})
	return
}
