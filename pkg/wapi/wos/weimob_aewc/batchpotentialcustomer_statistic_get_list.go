package weimob_aewc

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// BatchpotentialcustomerStatisticGetList
// @id 2219
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2219?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询潜在客户列表)
func (client *Client) BatchpotentialcustomerStatisticGetList(request *BatchpotentialcustomerStatisticGetListRequest) (response *BatchpotentialcustomerStatisticGetListResponse, err error) {
	rpcResponse := CreateBatchpotentialcustomerStatisticGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type BatchpotentialcustomerStatisticGetListRequest struct {
	*api.BaseRequest

	BatchId  string `json:"batchId,omitempty"`
	PageSize int64  `json:"pageSize,omitempty"`
	PageNo   int64  `json:"pageNo,omitempty"`
}

type BatchpotentialcustomerStatisticGetListResponse struct {
	Results    []BatchpotentialcustomerStatisticGetListResponseResults `json:"results,omitempty"`
	TotalCount int64                                                   `json:"totalCount,omitempty"`
}

func CreateBatchpotentialcustomerStatisticGetListRequest() (request *BatchpotentialcustomerStatisticGetListRequest) {
	request = &BatchpotentialcustomerStatisticGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_aewc", "v2.0", "batchpotentialcustomer/statistic/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateBatchpotentialcustomerStatisticGetListResponse() (response *api.BaseResponse[BatchpotentialcustomerStatisticGetListResponse]) {
	response = api.CreateResponse[BatchpotentialcustomerStatisticGetListResponse](&BatchpotentialcustomerStatisticGetListResponse{})
	return
}

type BatchpotentialcustomerStatisticGetListResponseResults struct {
	AddedTime       BatchpotentialcustomerStatisticGetListResponseAddedTime       `json:"addedTime,omitempty"`
	TagInfoList     []BatchpotentialcustomerStatisticGetListResponseTagInfoList   `json:"tagInfoList,omitempty"`
	SimpleUserModel BatchpotentialcustomerStatisticGetListResponseSimpleUserModel `json:"simpleUserModel,omitempty"`
	CustomerId      string                                                        `json:"customerId,omitempty"`
	BatchId         string                                                        `json:"batchId,omitempty"`
	Phone           string                                                        `json:"phone,omitempty"`
	Remark          string                                                        `json:"remark,omitempty"`
	AllocateTimes   int64                                                         `json:"allocateTimes,omitempty"`
	AllocatedTime   string                                                        `json:"allocatedTime,omitempty"`
	UserName        string                                                        `json:"userName,omitempty"`
	Stat            int64                                                         `json:"stat,omitempty"`
}

type BatchpotentialcustomerStatisticGetListResponseAddedTime struct {
}

type BatchpotentialcustomerStatisticGetListResponseTagInfoList struct {
	List []BatchpotentialcustomerStatisticGetListResponselist `json:"list,omitempty"`
	Id   int64                                                `json:"id,omitempty"`
	Name string                                               `json:"name,omitempty"`
}

type BatchpotentialcustomerStatisticGetListResponselist struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type BatchpotentialcustomerStatisticGetListResponseSimpleUserModel struct {
	UserName     string `json:"userName,omitempty"`
	OrgUserId    string `json:"orgUserId,omitempty"`
	DepartmentId string `json:"departmentId,omitempty"`
}
