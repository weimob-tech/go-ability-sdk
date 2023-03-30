package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NewmembercardGetRechargeLogPageListAndTotal
// @id 125
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取充值流水)
func (client *Client) NewmembercardGetRechargeLogPageListAndTotal(request *NewmembercardGetRechargeLogPageListAndTotalRequest) (response *NewmembercardGetRechargeLogPageListAndTotalResponse, err error) {
	rpcResponse := CreateNewmembercardGetRechargeLogPageListAndTotalResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NewmembercardGetRechargeLogPageListAndTotalRequest struct {
	*api.BaseRequest

	MemberCardNo string  `json:"memberCardNo,omitempty"`
	Name         string  `json:"name,omitempty"`
	Phone        string  `json:"phone,omitempty"`
	AmountMin    float64 `json:"amountMin,omitempty"`
	AmountMax    float64 `json:"amountMax,omitempty"`
	Sort         int     `json:"sort,omitempty"`
	Begintime    int64   `json:"begintime,omitempty"`
	Endtime      int64   `json:"endtime,omitempty"`
	PageIndex    int     `json:"pageIndex,omitempty"`
	PageSize     int     `json:"pageSize,omitempty"`
	Operator     string  `json:"operator,omitempty"`
}

type NewmembercardGetRechargeLogPageListAndTotalResponse struct {
}

func CreateNewmembercardGetRechargeLogPageListAndTotalRequest() (request *NewmembercardGetRechargeLogPageListAndTotalRequest) {
	request = &NewmembercardGetRechargeLogPageListAndTotalRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "newmembercard/GetRechargeLogPageListAndTotal", "api")
	request.Method = api.POST
	return
}

func CreateNewmembercardGetRechargeLogPageListAndTotalResponse() (response *api.BaseResponse[NewmembercardGetRechargeLogPageListAndTotalResponse]) {
	response = api.CreateResponse[NewmembercardGetRechargeLogPageListAndTotalResponse](&NewmembercardGetRechargeLogPageListAndTotalResponse{})
	return
}
