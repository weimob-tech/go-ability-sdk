package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembershipImportUser
// @id 1114
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=客户/会员导入)
func (client *Client) MembershipImportUser(request *MembershipImportUserRequest) (response *MembershipImportUserResponse, err error) {
	rpcResponse := CreateMembershipImportUserResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembershipImportUserRequest struct {
	*api.BaseRequest

	UserList                 []MembershipImportUserRequestUserList `json:"userList,omitempty"`
	ImportType               int                                   `json:"importType,omitempty"`
	CardType                 int                                   `json:"cardType,omitempty"`
	MembershipCardTemplateId int64                                 `json:"membershipCardTemplateId,omitempty"`
	NeedGetCardGift          int                                   `json:"needGetCardGift,omitempty"`
}

type MembershipImportUserResponse struct {
}

func CreateMembershipImportUserRequest() (request *MembershipImportUserRequest) {
	request = &MembershipImportUserRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "membership/importUser", "api")
	request.Method = api.POST
	return
}

func CreateMembershipImportUserResponse() (response *api.BaseResponse[MembershipImportUserResponse]) {
	response = api.CreateResponse[MembershipImportUserResponse](&MembershipImportUserResponse{})
	return
}

type MembershipImportUserRequestUserList struct {
	ExtMap               MembershipImportUserRequestExtMap `json:"extMap,omitempty"`
	Phone                string                            `json:"phone,omitempty"`
	GetChannel           int                               `json:"getChannel,omitempty"`
	UserName             string                            `json:"userName,omitempty"`
	Sex                  int                               `json:"sex,omitempty"`
	Birthday             int                               `json:"birthday,omitempty"`
	Province             string                            `json:"province,omitempty"`
	City                 string                            `json:"city,omitempty"`
	County               string                            `json:"county,omitempty"`
	Address              string                            `json:"address,omitempty"`
	AppChannel           int                               `json:"appChannel,omitempty"`
	GuideMobileNo        string                            `json:"guideMobileNo,omitempty"`
	Growth               int                               `json:"growth,omitempty"`
	Balance              float64                           `json:"balance,omitempty"`
	TotalBalance         float64                           `json:"totalBalance,omitempty"`
	Point                int                               `json:"point,omitempty"`
	TotalPoint           float64                           `json:"totalPoint,omitempty"`
	Tag                  string                            `json:"tag,omitempty"`
	UnionId              string                            `json:"unionId,omitempty"`
	OpenId               string                            `json:"openId,omitempty"`
	AppId                string                            `json:"appId,omitempty"`
	HomeStoreId          int64                             `json:"homeStoreId,omitempty"`
	GetMemberCardStoreId int64                             `json:"getMemberCardStoreId,omitempty"`
	BecomeStore          string                            `json:"becomeStore,omitempty"`
	RegionCode           string                            `json:"regionCode,omitempty"`
	BizUserId            string                            `json:"bizUserId,omitempty"`
	RankName             string                            `json:"rankName,omitempty"`
	Industry             int                               `json:"industry,omitempty"`
	Income               int                               `json:"income,omitempty"`
	Education            int                               `json:"education,omitempty"`
	IdCardNo             string                            `json:"idCardNo,omitempty"`
	CardCode             string                            `json:"cardCode,omitempty"`
	GetCardTime          int                               `json:"getCardTime,omitempty"`
}

type MembershipImportUserRequestExtMap struct {
}
