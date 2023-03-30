package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NewmembercardGetMemberInfoPagelistAndTotal
// @id 109
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=会员列表查询)
func (client *Client) NewmembercardGetMemberInfoPagelistAndTotal(request *NewmembercardGetMemberInfoPagelistAndTotalRequest) (response *NewmembercardGetMemberInfoPagelistAndTotalResponse, err error) {
	rpcResponse := CreateNewmembercardGetMemberInfoPagelistAndTotalResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NewmembercardGetMemberInfoPagelistAndTotalRequest struct {
	*api.BaseRequest

	Info                  string  `json:"info,omitempty"`
	ListlevelNo           []int   `json:"listlevelNo,omitempty"`
	ListTagId             []int64 `json:"listTagId,omitempty"`
	Sex                   int64   `json:"sex,omitempty"`
	AgeMin                int     `json:"ageMin,omitempty"`
	AgeMax                int     `json:"ageMax,omitempty"`
	ProvinceId            string  `json:"provinceId,omitempty"`
	CityId                string  `json:"cityId,omitempty"`
	LastConsumingTimeMin  int64   `json:"lastConsumingTimeMin,omitempty"`
	LastConsumingTimeMax  int64   `json:"lastConsumingTimeMax,omitempty"`
	ConsumingCountMin     int     `json:"consumingCountMin,omitempty"`
	ConsumingCountMax     int     `json:"consumingCountMax,omitempty"`
	AllConsumingAmountMin float64 `json:"allConsumingAmountMin,omitempty"`
	AllConsumingAmountMax float64 `json:"allConsumingAmountMax,omitempty"`
	IsNeedConditiontags   bool    `json:"isNeedConditiontags,omitempty"`
	SortBy                int     `json:"sortBy,omitempty"`
	AscOrDesc             int     `json:"ascOrDesc,omitempty"`
	PageIndex             int     `json:"pageIndex,omitempty"`
	PageSize              int     `json:"pageSize,omitempty"`
}

type NewmembercardGetMemberInfoPagelistAndTotalResponse struct {
}

func CreateNewmembercardGetMemberInfoPagelistAndTotalRequest() (request *NewmembercardGetMemberInfoPagelistAndTotalRequest) {
	request = &NewmembercardGetMemberInfoPagelistAndTotalRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "newmembercard/GetMemberInfoPagelistAndTotal", "api")
	request.Method = api.POST
	return
}

func CreateNewmembercardGetMemberInfoPagelistAndTotalResponse() (response *api.BaseResponse[NewmembercardGetMemberInfoPagelistAndTotalResponse]) {
	response = api.CreateResponse[NewmembercardGetMemberInfoPagelistAndTotalResponse](&NewmembercardGetMemberInfoPagelistAndTotalResponse{})
	return
}
