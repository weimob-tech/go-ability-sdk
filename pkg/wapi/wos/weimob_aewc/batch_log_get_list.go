package weimob_aewc

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// BatchLogGetList
// @id 2220
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2220?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=添加记录列表)
func (client *Client) BatchLogGetList(request *BatchLogGetListRequest) (response *BatchLogGetListResponse, err error) {
	rpcResponse := CreateBatchLogGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type BatchLogGetListRequest struct {
	*api.BaseRequest

	BatchId  string `json:"batchId,omitempty"`
	PageSize int64  `json:"pageSize,omitempty"`
	PageNo   int64  `json:"pageNo,omitempty"`
}

type BatchLogGetListResponse struct {
	Results    []BatchLogGetListResponseResults `json:"results,omitempty"`
	TotalCount int64                            `json:"totalCount,omitempty"`
}

func CreateBatchLogGetListRequest() (request *BatchLogGetListRequest) {
	request = &BatchLogGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_aewc", "v2.0", "batch/log/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateBatchLogGetListResponse() (response *api.BaseResponse[BatchLogGetListResponse]) {
	response = api.CreateResponse[BatchLogGetListResponse](&BatchLogGetListResponse{})
	return
}

type BatchLogGetListResponseResults struct {
	SimpleUserModel BatchLogGetListResponseSimpleUserModel `json:"simpleUserModel,omitempty"`
	Tags            []BatchLogGetListResponseTags          `json:"tags,omitempty"`
	AddNum          int64                                  `json:"addNum,omitempty"`
	AddType         int64                                  `json:"addType,omitempty"`
	AddedNum        int64                                  `json:"addedNum,omitempty"`
	AllocatedTime   int64                                  `json:"allocatedTime,omitempty"`
	CreateTime      int64                                  `json:"createTime,omitempty"`
	Deleted         bool                                   `json:"deleted,omitempty"`
	Follower        string                                 `json:"follower,omitempty"`
	FollowerNums    int64                                  `json:"followerNums,omitempty"`
	Id              string                                 `json:"id,omitempty"`
	Name            string                                 `json:"name,omitempty"`
	Source          string                                 `json:"source,omitempty"`
	UpdateTime      int64                                  `json:"updateTime,omitempty"`
	Wid             int64                                  `json:"wid,omitempty"`
}

type BatchLogGetListResponseSimpleUserModel struct {
	DepartmentId int64  `json:"departmentId,omitempty"`
	OrgUserId    string `json:"orgUserId,omitempty"`
	UserId       string `json:"userId,omitempty"`
	UserName     string `json:"userName,omitempty"`
}

type BatchLogGetListResponseTags struct {
	List []BatchLogGetListResponselist `json:"list,omitempty"`
	Id   string                        `json:"id,omitempty"`
	Name string                        `json:"name,omitempty"`
}

type BatchLogGetListResponselist struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
