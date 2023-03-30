package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NewmembercardGetPointsLogPageListAndTotal
// @id 122
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取积分流水)
func (client *Client) NewmembercardGetPointsLogPageListAndTotal(request *NewmembercardGetPointsLogPageListAndTotalRequest) (response *NewmembercardGetPointsLogPageListAndTotalResponse, err error) {
	rpcResponse := CreateNewmembercardGetPointsLogPageListAndTotalResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NewmembercardGetPointsLogPageListAndTotalRequest struct {
	*api.BaseRequest

	MemberCardNo  string `json:"memberCardNo,omitempty"`
	Name          string `json:"name,omitempty"`
	Phone         string `json:"phone,omitempty"`
	PointsPayType int64  `json:"pointsPayType,omitempty"`
	Sort          int    `json:"sort,omitempty"`
	Begintime     int64  `json:"begintime,omitempty"`
	Endtime       int64  `json:"endtime,omitempty"`
	PageIndex     int    `json:"pageIndex,omitempty"`
	PageSize      int    `json:"pageSize,omitempty"`
}

type NewmembercardGetPointsLogPageListAndTotalResponse struct {
}

func CreateNewmembercardGetPointsLogPageListAndTotalRequest() (request *NewmembercardGetPointsLogPageListAndTotalRequest) {
	request = &NewmembercardGetPointsLogPageListAndTotalRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "newmembercard/GetPointsLogPageListAndTotal", "api")
	request.Method = api.POST
	return
}

func CreateNewmembercardGetPointsLogPageListAndTotalResponse() (response *api.BaseResponse[NewmembercardGetPointsLogPageListAndTotalResponse]) {
	response = api.CreateResponse[NewmembercardGetPointsLogPageListAndTotalResponse](&NewmembercardGetPointsLogPageListAndTotalResponse{})
	return
}
