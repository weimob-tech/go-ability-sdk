package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantprivilegeQueryAccountDetailInfo
// @id 1594
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询账号详情信息)
func (client *Client) MerchantprivilegeQueryAccountDetailInfo(request *MerchantprivilegeQueryAccountDetailInfoRequest) (response *MerchantprivilegeQueryAccountDetailInfoResponse, err error) {
	rpcResponse := CreateMerchantprivilegeQueryAccountDetailInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantprivilegeQueryAccountDetailInfoRequest struct {
	*api.BaseRequest

	AccountWid int64  `json:"accountWid,omitempty"`
	LoginName  string `json:"loginName,omitempty"`
	AccountNum string `json:"accountNum,omitempty"`
}

type MerchantprivilegeQueryAccountDetailInfoResponse struct {
	AccountDepartmentDetailVoList []MerchantprivilegeQueryAccountDetailInfoResponseAccountDepartmentDetailVoList `json:"accountDepartmentDetailVoList,omitempty"`
	AccountRoleDetailVoList       []MerchantprivilegeQueryAccountDetailInfoResponseAccountRoleDetailVoList       `json:"accountRoleDetailVoList,omitempty"`
	Pid                           int64                                                                          `json:"pid,omitempty"`
	Wid                           int64                                                                          `json:"wid,omitempty"`
	LoginName                     string                                                                         `json:"loginName,omitempty"`
	Zone                          string                                                                         `json:"zone,omitempty"`
	AccountName                   string                                                                         `json:"accountName,omitempty"`
	AccountNum                    string                                                                         `json:"accountNum,omitempty"`
}

func CreateMerchantprivilegeQueryAccountDetailInfoRequest() (request *MerchantprivilegeQueryAccountDetailInfoRequest) {
	request = &MerchantprivilegeQueryAccountDetailInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchantprivilege/queryAccountDetailInfo", "api")
	request.Method = api.POST
	return
}

func CreateMerchantprivilegeQueryAccountDetailInfoResponse() (response *api.BaseResponse[MerchantprivilegeQueryAccountDetailInfoResponse]) {
	response = api.CreateResponse[MerchantprivilegeQueryAccountDetailInfoResponse](&MerchantprivilegeQueryAccountDetailInfoResponse{})
	return
}

type MerchantprivilegeQueryAccountDetailInfoResponseAccountDepartmentDetailVoList struct {
	SceneType int    `json:"sceneType,omitempty"`
	Id        int64  `json:"id,omitempty"`
	Code      string `json:"code,omitempty"`
}

type MerchantprivilegeQueryAccountDetailInfoResponseAccountRoleDetailVoList struct {
	MerchantRoleType int    `json:"merchantRoleType,omitempty"`
	RoleName         string `json:"roleName,omitempty"`
}
