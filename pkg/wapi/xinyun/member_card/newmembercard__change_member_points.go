package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NewmembercardChangeMemberPoints
// @id 113
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=增加减少积分)
func (client *Client) NewmembercardChangeMemberPoints(request *NewmembercardChangeMemberPointsRequest) (response *NewmembercardChangeMemberPointsResponse, err error) {
	rpcResponse := CreateNewmembercardChangeMemberPointsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NewmembercardChangeMemberPointsRequest struct {
	*api.BaseRequest

	MemberCardNo       string `json:"memberCardNo,omitempty"`
	Points             int    `json:"points,omitempty"`
	Title              string `json:"title,omitempty"`
	Remark             string `json:"remark,omitempty"`
	IsAboutGrowthValue bool   `json:"isAboutGrowthValue,omitempty"`
	Operator           string `json:"operator,omitempty"`
	PointsPayType      int    `json:"pointsPayType,omitempty"`
}

type NewmembercardChangeMemberPointsResponse struct {
}

func CreateNewmembercardChangeMemberPointsRequest() (request *NewmembercardChangeMemberPointsRequest) {
	request = &NewmembercardChangeMemberPointsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "newmembercard/ChangeMemberPoints", "api")
	request.Method = api.POST
	return
}

func CreateNewmembercardChangeMemberPointsResponse() (response *api.BaseResponse[NewmembercardChangeMemberPointsResponse]) {
	response = api.CreateResponse[NewmembercardChangeMemberPointsResponse](&NewmembercardChangeMemberPointsResponse{})
	return
}
