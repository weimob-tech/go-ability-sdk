package weimob_aewc

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GroupGetList
// @id 2213
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2213?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=筛选客户群列表)
func (client *Client) GroupGetList(request *GroupGetListRequest) (response *GroupGetListResponse, err error) {
	rpcResponse := CreateGroupGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GroupGetListRequest struct {
	*api.BaseRequest

	Name            string  `json:"name,omitempty"`
	State           int64   `json:"state,omitempty"`
	CreateTimeStart int64   `json:"createTimeStart,omitempty"`
	CreateTimeEnd   int64   `json:"createTimeEnd,omitempty"`
	OwnerIdList     []int64 `json:"ownerIdList,omitempty"`
	PageNum         int64   `json:"pageNum,omitempty"`
	PageSize        int64   `json:"pageSize,omitempty"`
}

type GroupGetListResponse struct {
	Results    []GroupGetListResponseResults `json:"results,omitempty"`
	TotalCount int64                         `json:"totalCount,omitempty"`
	TotalPage  int64                         `json:"totalPage,omitempty"`
	PageNum    int64                         `json:"pageNum,omitempty"`
	PageSize   int64                         `json:"pageSize,omitempty"`
}

func CreateGroupGetListRequest() (request *GroupGetListRequest) {
	request = &GroupGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_aewc", "v2.0", "group/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGroupGetListResponse() (response *api.BaseResponse[GroupGetListResponse]) {
	response = api.CreateResponse[GroupGetListResponse](&GroupGetListResponse{})
	return
}

type GroupGetListResponseResults struct {
	Name            string `json:"name,omitempty"`
	Owner           string `json:"owner,omitempty"`
	GroupCreateTime int64  `json:"groupCreateTime,omitempty"`
	Num             int64  `json:"num,omitempty"`
	CustomerNum     int64  `json:"customerNum,omitempty"`
	JoinNum         int64  `json:"joinNum,omitempty"`
	QuitNum         int64  `json:"quitNum,omitempty"`
	ChatId          string `json:"chatId,omitempty"`
	State           int64  `json:"state,omitempty"`
}
