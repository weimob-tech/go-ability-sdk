package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserAccountMerge
// @id 2103
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2103?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=账号合并)
func (client *Client) UserAccountMerge(request *UserAccountMergeRequest) (response *UserAccountMergeResponse, err error) {
	rpcResponse := CreateUserAccountMergeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserAccountMergeRequest struct {
	*api.BaseRequest

	SourceInfo UserAccountMergeRequestSourceInfo `json:"sourceInfo,omitempty"`
	TargetInfo UserAccountMergeRequestTargetInfo `json:"targetInfo,omitempty"`
	BosId      int64                             `json:"bosId,omitempty"`
	MergeType  int64                             `json:"mergeType,omitempty"`
}

type UserAccountMergeResponse struct {
	MergeResult bool  `json:"mergeResult,omitempty"`
	SuperWid    int64 `json:"superWid,omitempty"`
}

func CreateUserAccountMergeRequest() (request *UserAccountMergeRequest) {
	request = &UserAccountMergeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "user/account/merge", "apigw")
	request.Method = api.POST
	return
}

func CreateUserAccountMergeResponse() (response *api.BaseResponse[UserAccountMergeResponse]) {
	response = api.CreateResponse[UserAccountMergeResponse](&UserAccountMergeResponse{})
	return
}

type UserAccountMergeRequestSourceInfo struct {
	SourceWid  int64  `json:"sourceWid,omitempty"`
	AppId      string `json:"appId,omitempty"`
	OriginalId string `json:"originalId,omitempty"`
}

type UserAccountMergeRequestTargetInfo struct {
	TargetWid int64 `json:"targetWid,omitempty"`
}
