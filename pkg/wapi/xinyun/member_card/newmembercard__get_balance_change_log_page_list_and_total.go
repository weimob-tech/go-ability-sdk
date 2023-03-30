package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NewmembercardGetBalanceChangeLogPageListAndTotal
// @id 123
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取余额流水)
func (client *Client) NewmembercardGetBalanceChangeLogPageListAndTotal(request *NewmembercardGetBalanceChangeLogPageListAndTotalRequest) (response *NewmembercardGetBalanceChangeLogPageListAndTotalResponse, err error) {
	rpcResponse := CreateNewmembercardGetBalanceChangeLogPageListAndTotalResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NewmembercardGetBalanceChangeLogPageListAndTotalRequest struct {
	*api.BaseRequest

	MemberCardNo string `json:"memberCardNo,omitempty"`
	Sort         int    `json:"sort,omitempty"`
	BeginTime    int64  `json:"beginTime,omitempty"`
	EndTime      int64  `json:"endTime,omitempty"`
	PageIndex    int    `json:"pageIndex,omitempty"`
	PageSize     int    `json:"pageSize,omitempty"`
}

type NewmembercardGetBalanceChangeLogPageListAndTotalResponse struct {
}

func CreateNewmembercardGetBalanceChangeLogPageListAndTotalRequest() (request *NewmembercardGetBalanceChangeLogPageListAndTotalRequest) {
	request = &NewmembercardGetBalanceChangeLogPageListAndTotalRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "newmembercard/GetBalanceChangeLogPageListAndTotal", "api")
	request.Method = api.POST
	return
}

func CreateNewmembercardGetBalanceChangeLogPageListAndTotalResponse() (response *api.BaseResponse[NewmembercardGetBalanceChangeLogPageListAndTotalResponse]) {
	response = api.CreateResponse[NewmembercardGetBalanceChangeLogPageListAndTotalResponse](&NewmembercardGetBalanceChangeLogPageListAndTotalResponse{})
	return
}
