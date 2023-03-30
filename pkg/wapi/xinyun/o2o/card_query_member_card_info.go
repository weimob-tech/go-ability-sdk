package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CardQueryMemberCardInfo
// @id 237
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询会员已领取卡券)
func (client *Client) CardQueryMemberCardInfo(request *CardQueryMemberCardInfoRequest) (response *CardQueryMemberCardInfoResponse, err error) {
	rpcResponse := CreateCardQueryMemberCardInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CardQueryMemberCardInfoRequest struct {
	*api.BaseRequest

	PageNum    int    `json:"pageNum,omitempty"`
	PageSize   int    `json:"pageSize,omitempty"`
	OpenId     string `json:"openId,omitempty"`
	Status     int    `json:"status,omitempty"`
	CustomerId int64  `json:"customerId,omitempty"`
	CardCode   string `json:"cardCode,omitempty"`
	Mobile     string `json:"mobile,omitempty"`
	Wid        string `json:"wid,omitempty"`
}

type CardQueryMemberCardInfoResponse struct {
}

func CreateCardQueryMemberCardInfoRequest() (request *CardQueryMemberCardInfoRequest) {
	request = &CardQueryMemberCardInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "card/queryMemberCardInfo", "api")
	request.Method = api.POST
	return
}

func CreateCardQueryMemberCardInfoResponse() (response *api.BaseResponse[CardQueryMemberCardInfoResponse]) {
	response = api.CreateResponse[CardQueryMemberCardInfoResponse](&CardQueryMemberCardInfoResponse{})
	return
}
