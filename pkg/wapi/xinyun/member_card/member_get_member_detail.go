package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberGetMemberDetail
// @id 652
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询客户详情)
func (client *Client) MemberGetMemberDetail(request *MemberGetMemberDetailRequest) (response *MemberGetMemberDetailResponse, err error) {
	rpcResponse := CreateMemberGetMemberDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberGetMemberDetailRequest struct {
	*api.BaseRequest

	Type           int    `json:"type,omitempty"`
	Code           string `json:"code,omitempty"`
	Phone          string `json:"phone,omitempty"`
	Wid            string `json:"wid,omitempty"`
	IsNeedTagsInfo bool   `json:"isNeedTagsInfo,omitempty"`
	Mid            string `json:"mid,omitempty"`
	IsNeedExtInfo  bool   `json:"isNeedExtInfo,omitempty"`
}

type MemberGetMemberDetailResponse struct {
	Wid                    int64   `json:"wid,omitempty"`
	Phone                  string  `json:"phone,omitempty"`
	Mid                    int64   `json:"mid,omitempty"`
	GmtCreate              int64   `json:"gmtCreate,omitempty"`
	CardPublishTime        int64   `json:"cardPublishTime,omitempty"`
	Name                   string  `json:"name,omitempty"`
	Birthday               int64   `json:"birthday,omitempty"`
	IdCard                 string  `json:"idCard,omitempty"`
	Address                string  `json:"address,omitempty"`
	Mail                   string  `json:"mail,omitempty"`
	Sex                    int64   `json:"sex,omitempty"`
	Education              string  `json:"education,omitempty"`
	Hobby                  string  `json:"hobby,omitempty"`
	Income                 string  `json:"income,omitempty"`
	Industry               string  `json:"industry,omitempty"`
	WechatStatus           int64   `json:"wechatStatus,omitempty"`
	WechatCode             string  `json:"wechatCode,omitempty"`
	TotalPoint             int64   `json:"totalPoint,omitempty"`
	CurrentPoint           int64   `json:"currentPoint,omitempty"`
	CurrentAmount          int64   `json:"currentAmount,omitempty"`
	CurrentGrowth          int64   `json:"currentGrowth,omitempty"`
	OrderConsumeAmount     int64   `json:"orderConsumeAmount,omitempty"`
	OrderConsumeCount      int64   `json:"orderConsumeCount,omitempty"`
	AvgOrderConsumeAmount  int64   `json:"avgOrderConsumeAmount,omitempty"`
	LastTradeTime          int64   `json:"lastTradeTime,omitempty"`
	TagList                []int64 `json:"tagList,omitempty"`
	RankName               string  `json:"rankName,omitempty"`
	RankId                 string  `json:"rankId,omitempty"`
	Status                 int64   `json:"status,omitempty"`
	ExtInfoList            []int64 `json:"extInfoList,omitempty"`
	StoreId                int64   `json:"storeId,omitempty"`
	HomeStoreId            int64   `json:"homeStoreId,omitempty"`
	AvailableDepositAmount int64   `json:"availableDepositAmount,omitempty"`
	AvailableGiveAmount    int64   `json:"availableGiveAmount,omitempty"`
	Code                   string  `json:"code,omitempty"`
}

func CreateMemberGetMemberDetailRequest() (request *MemberGetMemberDetailRequest) {
	request = &MemberGetMemberDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "member/getMemberDetail", "api")
	request.Method = api.POST
	return
}

func CreateMemberGetMemberDetailResponse() (response *api.BaseResponse[MemberGetMemberDetailResponse]) {
	response = api.CreateResponse[MemberGetMemberDetailResponse](&MemberGetMemberDetailResponse{})
	return
}
