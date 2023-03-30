package kld

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// KLDMemberCardGetPointsLogPageListAndTotal
// @id 214
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取积分流水)
func (client *Client) KLDMemberCardGetPointsLogPageListAndTotal(request *KLDMemberCardGetPointsLogPageListAndTotalRequest) (response *KLDMemberCardGetPointsLogPageListAndTotalResponse, err error) {
	rpcResponse := CreateKLDMemberCardGetPointsLogPageListAndTotalResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type KLDMemberCardGetPointsLogPageListAndTotalRequest struct {
	*api.BaseRequest

	MemberCardNo    string `json:"memberCardNo,omitempty"`
	Name            string `json:"name,omitempty"`
	Phone           string `json:"phone,omitempty"`
	PointsPayType   int64  `json:"pointsPayType,omitempty"`
	Sort            int    `json:"sort,omitempty"`
	Begintime       int64  `json:"begintime,omitempty"`
	Endtime         int64  `json:"endtime,omitempty"`
	PageIndex       int    `json:"pageIndex,omitempty"`
	PageSize        int    `json:"pageSize,omitempty"`
	IsOnlyEffective bool   `json:"isOnlyEffective,omitempty"`
}

type KLDMemberCardGetPointsLogPageListAndTotalResponse struct {
}

func CreateKLDMemberCardGetPointsLogPageListAndTotalRequest() (request *KLDMemberCardGetPointsLogPageListAndTotalRequest) {
	request = &KLDMemberCardGetPointsLogPageListAndTotalRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("kld", "1_0", "KLDMemberCard/GetPointsLogPageListAndTotal", "api")
	request.Method = api.POST
	return
}

func CreateKLDMemberCardGetPointsLogPageListAndTotalResponse() (response *api.BaseResponse[KLDMemberCardGetPointsLogPageListAndTotalResponse]) {
	response = api.CreateResponse[KLDMemberCardGetPointsLogPageListAndTotalResponse](&KLDMemberCardGetPointsLogPageListAndTotalResponse{})
	return
}
