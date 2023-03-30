package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserChannelGetList
// @id 2065
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2065?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取用户渠道信息列表)
func (client *Client) UserChannelGetList(request *UserChannelGetListRequest) (response *UserChannelGetListResponse, err error) {
	rpcResponse := CreateUserChannelGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserChannelGetListRequest struct {
	*api.BaseRequest

	BosId           int64   `json:"bosId,omitempty"`
	Wid             int64   `json:"wid,omitempty"`
	ChannelTypeList []int64 `json:"channelTypeList,omitempty"`
	AppId           string  `json:"appId,omitempty"`
	PageSize        int64   `json:"pageSize,omitempty"`
	PageNum         int64   `json:"pageNum,omitempty"`
	QueryType       int64   `json:"queryType,omitempty"`
}

type UserChannelGetListResponse struct {
	UserChannelInfoDtoList []UserChannelGetListResponseUserChannelInfoDtoList `json:"userChannelInfoDtoList,omitempty"`
	SuperWid               int64                                              `json:"superWid,omitempty"`
}

func CreateUserChannelGetListRequest() (request *UserChannelGetListRequest) {
	request = &UserChannelGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "user/channel/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateUserChannelGetListResponse() (response *api.BaseResponse[UserChannelGetListResponse]) {
	response = api.CreateResponse[UserChannelGetListResponse](&UserChannelGetListResponse{})
	return
}

type UserChannelGetListResponseUserChannelInfoDtoList struct {
	BasicChannelInfo  UserChannelGetListResponseBasicChannelInfo  `json:"basicChannelInfo,omitempty"`
	ExtendChannelInfo UserChannelGetListResponseExtendChannelInfo `json:"extendChannelInfo,omitempty"`
	ChannelType       int64                                       `json:"channelType,omitempty"`
	AppId             string                                      `json:"appId,omitempty"`
	OpenId            string                                      `json:"openId,omitempty"`
}

type UserChannelGetListResponseBasicChannelInfo struct {
	Phone     string `json:"phone,omitempty"`
	NickName  string `json:"nickName,omitempty"`
	Gender    int64  `json:"gender,omitempty"`
	AvatarUrl string `json:"avatarUrl,omitempty"`
	Country   string `json:"country,omitempty"`
	Province  string `json:"province,omitempty"`
	City      string `json:"city	,omitempty"`
	Language  string `json:"language	,omitempty"`
	Unionid   string `json:"unionid,omitempty"`
	Remark    string `json:"remark,omitempty"`
}

type UserChannelGetListResponseExtendChannelInfo struct {
}
