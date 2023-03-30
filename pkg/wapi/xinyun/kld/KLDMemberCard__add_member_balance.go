package kld

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// KLDMemberCardAddMemberBalance
// @id 234
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=增加余额)
func (client *Client) KLDMemberCardAddMemberBalance(request *KLDMemberCardAddMemberBalanceRequest) (response *KLDMemberCardAddMemberBalanceResponse, err error) {
	rpcResponse := CreateKLDMemberCardAddMemberBalanceResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type KLDMemberCardAddMemberBalanceRequest struct {
	*api.BaseRequest

	MemberCardNo string  `json:"MemberCardNo,omitempty"`
	Amount       float64 `json:"Amount,omitempty"`
	Remark       string  `json:"Remark,omitempty"`
	Title        string  `json:"Title,omitempty"`
	Operator     string  `json:"Operator,omitempty"`
}

type KLDMemberCardAddMemberBalanceResponse struct {
}

func CreateKLDMemberCardAddMemberBalanceRequest() (request *KLDMemberCardAddMemberBalanceRequest) {
	request = &KLDMemberCardAddMemberBalanceRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("kld", "1_0", "KLDMemberCard/AddMemberBalance", "api")
	request.Method = api.POST
	return
}

func CreateKLDMemberCardAddMemberBalanceResponse() (response *api.BaseResponse[KLDMemberCardAddMemberBalanceResponse]) {
	response = api.CreateResponse[KLDMemberCardAddMemberBalanceResponse](&KLDMemberCardAddMemberBalanceResponse{})
	return
}
