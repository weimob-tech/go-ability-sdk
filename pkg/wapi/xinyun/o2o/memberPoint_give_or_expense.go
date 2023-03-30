package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberPointGiveOrExpense
// @id 423
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=为单个会员调整积分)
func (client *Client) MemberPointGiveOrExpense(request *MemberPointGiveOrExpenseRequest) (response *MemberPointGiveOrExpenseResponse, err error) {
	rpcResponse := CreateMemberPointGiveOrExpenseResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberPointGiveOrExpenseRequest struct {
	*api.BaseRequest

	MerchantId      int64  `json:"merchantId,omitempty"`
	CustomerId      int64  `json:"customerId,omitempty"`
	Point           int64  `json:"point,omitempty"`
	OperationType   int    `json:"operationType,omitempty"`
	Reason          string `json:"reason,omitempty"`
	ThirdBusinessId string `json:"thirdBusinessId,omitempty"`
}

type MemberPointGiveOrExpenseResponse struct {
}

func CreateMemberPointGiveOrExpenseRequest() (request *MemberPointGiveOrExpenseRequest) {
	request = &MemberPointGiveOrExpenseRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "memberPoint/giveOrExpense", "api")
	request.Method = api.POST
	return
}

func CreateMemberPointGiveOrExpenseResponse() (response *api.BaseResponse[MemberPointGiveOrExpenseResponse]) {
	response = api.CreateResponse[MemberPointGiveOrExpenseResponse](&MemberPointGiveOrExpenseResponse{})
	return
}
