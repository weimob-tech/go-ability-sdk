package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NewmembercardGetConsumingLogPageListAndTotal
// @id 124
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取消费流水)
func (client *Client) NewmembercardGetConsumingLogPageListAndTotal(request *NewmembercardGetConsumingLogPageListAndTotalRequest) (response *NewmembercardGetConsumingLogPageListAndTotalResponse, err error) {
	rpcResponse := CreateNewmembercardGetConsumingLogPageListAndTotalResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NewmembercardGetConsumingLogPageListAndTotalRequest struct {
	*api.BaseRequest

	MemberCardNo string  `json:"memberCardNo,omitempty"`
	Name         string  `json:"name,omitempty"`
	CouponType   int64   `json:"couponType,omitempty"`
	Phone        string  `json:"phone,omitempty"`
	AmountMin    float64 `json:"amountMin,omitempty"`
	AmountMax    float64 `json:"amountMax,omitempty"`
	Sort         int     `json:"sort,omitempty"`
	Begintime    int64   `json:"begintime,omitempty"`
	Endtime      int64   `json:"endtime,omitempty"`
	PageIndex    int     `json:"pageIndex,omitempty"`
	PageSize     int     `json:"pageSize,omitempty"`
}

type NewmembercardGetConsumingLogPageListAndTotalResponse struct {
}

func CreateNewmembercardGetConsumingLogPageListAndTotalRequest() (request *NewmembercardGetConsumingLogPageListAndTotalRequest) {
	request = &NewmembercardGetConsumingLogPageListAndTotalRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "newmembercard/GetConsumingLogPageListAndTotal", "api")
	request.Method = api.POST
	return
}

func CreateNewmembercardGetConsumingLogPageListAndTotalResponse() (response *api.BaseResponse[NewmembercardGetConsumingLogPageListAndTotalResponse]) {
	response = api.CreateResponse[NewmembercardGetConsumingLogPageListAndTotalResponse](&NewmembercardGetConsumingLogPageListAndTotalResponse{})
	return
}
