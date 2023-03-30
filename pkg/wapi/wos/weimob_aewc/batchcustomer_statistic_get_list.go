package weimob_aewc

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// BatchcustomerStatisticGetList
// @id 2225
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2225?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询客户添加统计数据)
func (client *Client) BatchcustomerStatisticGetList(request *BatchcustomerStatisticGetListRequest) (response *BatchcustomerStatisticGetListResponse, err error) {
	rpcResponse := CreateBatchcustomerStatisticGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type BatchcustomerStatisticGetListRequest struct {
	*api.BaseRequest

	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
	PageNum   int64  `json:"pageNum,omitempty"`
	PageSize  int64  `json:"pageSize,omitempty"`
}

type BatchcustomerStatisticGetListResponse struct {
	Result     []BatchcustomerStatisticGetListResponseResult `json:"result,omitempty"`
	PageNum    int64                                         `json:"pageNum,omitempty"`
	PageSize   int64                                         `json:"pageSize,omitempty"`
	TotalCount int64                                         `json:"totalCount,omitempty"`
}

func CreateBatchcustomerStatisticGetListRequest() (request *BatchcustomerStatisticGetListRequest) {
	request = &BatchcustomerStatisticGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_aewc", "v2.0", "batchcustomer/statistic/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateBatchcustomerStatisticGetListResponse() (response *api.BaseResponse[BatchcustomerStatisticGetListResponse]) {
	response = api.CreateResponse[BatchcustomerStatisticGetListResponse](&BatchcustomerStatisticGetListResponse{})
	return
}

type BatchcustomerStatisticGetListResponseResult struct {
	TaskCreateTime       string `json:"taskCreateTime,omitempty"`
	GetGuestSuccessRate  int64  `json:"getGuestSuccessRate,omitempty"`
	AddedCustomerNum     int64  `json:"addedCustomerNum,omitempty"`
	TaskId               string `json:"taskId,omitempty"`
	ShouldExeUserNum     int64  `json:"shouldExeUserNum,omitempty"`
	ExecutedUserNum      int64  `json:"executedUserNum,omitempty"`
	ShouldAddCustomerNum int64  `json:"shouldAddCustomerNum,omitempty"`
	AssignedCustomerNum  int64  `json:"assignedCustomerNum,omitempty"`
	TaskName             string `json:"taskName,omitempty"`
	Deleted              bool   `json:"deleted,omitempty"`
	TaskExecutedRate     int64  `json:"taskExecutedRate,omitempty"`
}
