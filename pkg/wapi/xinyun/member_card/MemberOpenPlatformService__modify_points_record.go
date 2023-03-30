package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberOpenPlatformServiceModifyPointsRecord
// @id 171
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改会员积分)
func (client *Client) MemberOpenPlatformServiceModifyPointsRecord(request *MemberOpenPlatformServiceModifyPointsRecordRequest) (response *MemberOpenPlatformServiceModifyPointsRecordResponse, err error) {
	rpcResponse := CreateMemberOpenPlatformServiceModifyPointsRecordResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberOpenPlatformServiceModifyPointsRecordRequest struct {
	*api.BaseRequest

	CustomerNos     []string `json:"customer_nos,omitempty"`
	Mulshopid       int64    `json:"mulshopid,omitempty"`
	Points          int      `json:"points,omitempty"`
	TransactionMode int      `json:"transaction_mode,omitempty"`
	Opreater        string   `json:"opreater,omitempty"`
	Remark          string   `json:"remark,omitempty"`
	VerifyId        string   `json:"verify_id,omitempty"`
}

type MemberOpenPlatformServiceModifyPointsRecordResponse struct {
}

func CreateMemberOpenPlatformServiceModifyPointsRecordRequest() (request *MemberOpenPlatformServiceModifyPointsRecordRequest) {
	request = &MemberOpenPlatformServiceModifyPointsRecordRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "MemberOpenPlatformService/ModifyPointsRecord", "api")
	request.Method = api.POST
	return
}

func CreateMemberOpenPlatformServiceModifyPointsRecordResponse() (response *api.BaseResponse[MemberOpenPlatformServiceModifyPointsRecordResponse]) {
	response = api.CreateResponse[MemberOpenPlatformServiceModifyPointsRecordResponse](&MemberOpenPlatformServiceModifyPointsRecordResponse{})
	return
}
