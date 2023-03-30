package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NewmembercardRechargeForAdmin
// @id 116
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=会员充值接口——后台接口)
func (client *Client) NewmembercardRechargeForAdmin(request *NewmembercardRechargeForAdminRequest) (response *NewmembercardRechargeForAdminResponse, err error) {
	rpcResponse := CreateNewmembercardRechargeForAdminResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NewmembercardRechargeForAdminRequest struct {
	*api.BaseRequest

	MemberCardNo string  `json:"memberCardNo,omitempty"`
	Amount       float64 `json:"amount,omitempty"`
	Remark       string  `json:"remark,omitempty"`
	Operator     string  `json:"operator,omitempty"`
	PassWord     string  `json:"passWord,omitempty"`
	Payment      string  `json:"payment,omitempty"`
}

type NewmembercardRechargeForAdminResponse struct {
}

func CreateNewmembercardRechargeForAdminRequest() (request *NewmembercardRechargeForAdminRequest) {
	request = &NewmembercardRechargeForAdminRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "newmembercard/RechargeForAdmin", "api")
	request.Method = api.POST
	return
}

func CreateNewmembercardRechargeForAdminResponse() (response *api.BaseResponse[NewmembercardRechargeForAdminResponse]) {
	response = api.CreateResponse[NewmembercardRechargeForAdminResponse](&NewmembercardRechargeForAdminResponse{})
	return
}
