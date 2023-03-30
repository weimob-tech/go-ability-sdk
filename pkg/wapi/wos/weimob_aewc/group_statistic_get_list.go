package weimob_aewc

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GroupStatisticGetList
// @id 2217
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2217?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=分页查询客户群数据统计列表)
func (client *Client) GroupStatisticGetList(request *GroupStatisticGetListRequest) (response *GroupStatisticGetListResponse, err error) {
	rpcResponse := CreateGroupStatisticGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GroupStatisticGetListRequest struct {
	*api.BaseRequest

	StatisticType int64 `json:"statisticType,omitempty"`
	DateType      int64 `json:"dateType,omitempty"`
	StartTime     int64 `json:"startTime,omitempty"`
	EndTime       int64 `json:"endTime,omitempty"`
	PageSize      int64 `json:"pageSize,omitempty"`
	PageNum       int64 `json:"pageNum,omitempty"`
}

type GroupStatisticGetListResponse struct {
	Results    []GroupStatisticGetListResponseResults `json:"results,omitempty"`
	TotalCount int64                                  `json:"totalCount,omitempty"`
	TotalPage  int64                                  `json:"totalPage,omitempty"`
	PageSize   int64                                  `json:"pageSize,omitempty"`
	PageNum    int64                                  `json:"pageNum,omitempty"`
}

func CreateGroupStatisticGetListRequest() (request *GroupStatisticGetListRequest) {
	request = &GroupStatisticGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_aewc", "v2.0", "group/statistic/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGroupStatisticGetListResponse() (response *api.BaseResponse[GroupStatisticGetListResponse]) {
	response = api.CreateResponse[GroupStatisticGetListResponse](&GroupStatisticGetListResponse{})
	return
}

type GroupStatisticGetListResponseResults struct {
	CorpId                string `json:"corpId,omitempty"`
	StatisticType         int64  `json:"statisticType,omitempty"`
	DateType              int64  `json:"dateType,omitempty"`
	StatisticDate         string `json:"statisticDate,omitempty"`
	EmployeeName          string `json:"employeeName,omitempty"`
	DepartmentName        string `json:"departmentName,omitempty"`
	GroupNum              int64  `json:"groupNum,omitempty"`
	IncreaseGroupNum      int64  `json:"increaseGroupNum,omitempty"`
	GroupTotalMemberNum   int64  `json:"groupTotalMemberNum,omitempty"`
	GroupCustomerNum      int64  `json:"groupCustomerNum,omitempty"`
	JoinGroupCustomerNum  int64  `json:"joinGroupCustomerNum,omitempty"`
	QuitGroupCustomerNum  int64  `json:"quitGroupCustomerNum,omitempty"`
	SendMsgGroupNum       int64  `json:"sendMsgGroupNum,omitempty"`
	SendMsgGroupMemberNum int64  `json:"sendMsgGroupMemberNum,omitempty"`
	GroupMsgTotalNum      int64  `json:"groupMsgTotalNum,omitempty"`
	GroupName             string `json:"groupName,omitempty"`
}
