package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserSearch
// @id 2227
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2227?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=客户搜索)
func (client *Client) UserSearchV2_1(request *UserSearchRequest) (response *UserSearchResponse, err error) {
	rpcResponse := CreateUserSearchResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserSearchRequest struct {
	*api.BaseRequest

	Query              []UserSearchRequestQuery `json:"query,omitempty"`
	Vid                int64                    `json:"vid,omitempty"`
	VidType            int64                    `json:"vidType,omitempty"`
	Fields             []int64                  `json:"fields,omitempty"`
	PageSize           int64                    `json:"pageSize,omitempty"`
	PageNum            int64                    `json:"pageNum,omitempty"`
	IsReturnPageResult int64                    `json:"isReturnPageResult,omitempty"`
}

type UserSearchResponse struct {
	Result    string `json:"result,omitempty"`
	TotalSize int64  `json:"totalSize,omitempty"`
}

func CreateUserSearchRequest() (request *UserSearchRequest) {
	request = &UserSearchRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.1", "user/search", "apigw")
	request.Method = api.POST
	return
}

func CreateUserSearchResponse() (response *api.BaseResponse[UserSearchResponse]) {
	response = api.CreateResponse[UserSearchResponse](&UserSearchResponse{})
	return
}

type UserSearchRequestQuery struct {
	Value UserSearchRequestValue `json:"value,omitempty"`
	Field string                 `json:"field,omitempty"`
}

type UserSearchRequestValue struct {
}
