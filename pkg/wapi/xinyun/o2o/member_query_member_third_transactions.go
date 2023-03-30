package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberQueryMemberThirdTransactions
// @id 2073
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=第三方查询会员余额交易记录)
func (client *Client) MemberQueryMemberThirdTransactions(request *MemberQueryMemberThirdTransactionsRequest) (response *MemberQueryMemberThirdTransactionsResponse, err error) {
	rpcResponse := CreateMemberQueryMemberThirdTransactionsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberQueryMemberThirdTransactionsRequest struct {
	*api.BaseRequest

	Wid        int64  `json:"wid,omitempty"`
	StartTime  int    `json:"startTime,omitempty"`
	EndTime    int    `json:"endTime,omitempty"`
	StoreId    int64  `json:"storeId,omitempty"`
	PayWay     int    `json:"payWay,omitempty"`
	ChangeType int    `json:"changeType,omitempty"`
	OrderNo    string `json:"orderNo,omitempty"`
	PageNum    int    `json:"pageNum,omitempty"`
	PageSize   int    `json:"pageSize,omitempty"`
}

type MemberQueryMemberThirdTransactionsResponse struct {
}

func CreateMemberQueryMemberThirdTransactionsRequest() (request *MemberQueryMemberThirdTransactionsRequest) {
	request = &MemberQueryMemberThirdTransactionsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "member/queryMemberThirdTransactions", "api")
	request.Method = api.POST
	return
}

func CreateMemberQueryMemberThirdTransactionsResponse() (response *api.BaseResponse[MemberQueryMemberThirdTransactionsResponse]) {
	response = api.CreateResponse[MemberQueryMemberThirdTransactionsResponse](&MemberQueryMemberThirdTransactionsResponse{})
	return
}
