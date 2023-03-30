package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NewmembercardConsumeForAdmin
// @id 115
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=会员消费接口——后台接口)
func (client *Client) NewmembercardConsumeForAdmin(request *NewmembercardConsumeForAdminRequest) (response *NewmembercardConsumeForAdminResponse, err error) {
	rpcResponse := CreateNewmembercardConsumeForAdminResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NewmembercardConsumeForAdminRequest struct {
	*api.BaseRequest

	MemberCardNo   string  `json:"memberCardNo,omitempty"`
	TotalAmount    float64 `json:"totalAmount,omitempty"`
	RealAmount     float64 `json:"realAmount,omitempty"`
	BlancheAmount  float64 `json:"blancheAmount,omitempty"`
	Points         int     `json:"points,omitempty"`
	CashAmount     float64 `json:"cashAmount,omitempty"`
	BackAmount     float64 `json:"backAmount,omitempty"`
	SnNo           string  `json:"snNo,omitempty"`
	Remark         string  `json:"remark,omitempty"`
	Operator       string  `json:"operator,omitempty"`
	AdminPassWord  string  `json:"adminPassWord,omitempty"`
	MemberPassWord string  `json:"memberPassWord,omitempty"`
}

type NewmembercardConsumeForAdminResponse struct {
}

func CreateNewmembercardConsumeForAdminRequest() (request *NewmembercardConsumeForAdminRequest) {
	request = &NewmembercardConsumeForAdminRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "newmembercard/ConsumeForAdmin", "api")
	request.Method = api.POST
	return
}

func CreateNewmembercardConsumeForAdminResponse() (response *api.BaseResponse[NewmembercardConsumeForAdminResponse]) {
	response = api.CreateResponse[NewmembercardConsumeForAdminResponse](&NewmembercardConsumeForAdminResponse{})
	return
}
