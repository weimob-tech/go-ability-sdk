package weimob_aewc

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SoptaskFetchSopStatisticsPage
// @id 2224
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2224?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=SOP执行统计接口)
func (client *Client) SoptaskFetchSopStatisticsPage(request *SoptaskFetchSopStatisticsPageRequest) (response *SoptaskFetchSopStatisticsPageResponse, err error) {
	rpcResponse := CreateSoptaskFetchSopStatisticsPageResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SoptaskFetchSopStatisticsPageRequest struct {
	*api.BaseRequest

	BasicInfo           SoptaskFetchSopStatisticsPageRequestBasicInfo `json:"basicInfo,omitempty"`
	EnableTransformData bool                                          `json:"enableTransformData,omitempty"`
	StartTime           string                                        `json:"startTime,omitempty"`
	EndTime             string                                        `json:"endTime,omitempty"`
	TaskPackageType     int64                                         `json:"taskPackageType,omitempty"`
	PageNum             int64                                         `json:"pageNum,omitempty"`
	PageSize            int64                                         `json:"pageSize,omitempty"`
}

type SoptaskFetchSopStatisticsPageResponse struct {
	Code         SoptaskFetchSopStatisticsPageResponseCode `json:"code,omitempty"`
	Data         SoptaskFetchSopStatisticsPageResponseData `json:"data,omitempty"`
	GlobalTicket string                                    `json:"globalTicket,omitempty"`
}

func CreateSoptaskFetchSopStatisticsPageRequest() (request *SoptaskFetchSopStatisticsPageRequest) {
	request = &SoptaskFetchSopStatisticsPageRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_aewc", "v2.0", "soptask/fetchSopStatisticsPage", "apigw")
	request.Method = api.POST
	return
}

func CreateSoptaskFetchSopStatisticsPageResponse() (response *api.BaseResponse[SoptaskFetchSopStatisticsPageResponse]) {
	response = api.CreateResponse[SoptaskFetchSopStatisticsPageResponse](&SoptaskFetchSopStatisticsPageResponse{})
	return
}

type SoptaskFetchSopStatisticsPageRequestBasicInfo struct {
}

type SoptaskFetchSopStatisticsPageResponseCode struct {
	Errcode string `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

type SoptaskFetchSopStatisticsPageResponseData struct {
	Data       SoptaskFetchSopStatisticsPageResponseData2     `json:"data,omitempty"`
	Results    []SoptaskFetchSopStatisticsPageResponseResults `json:"results,omitempty"`
	TotalPage  int64                                          `json:"totalPage,omitempty"`
	PageSize   int64                                          `json:"pageSize,omitempty"`
	TotalCount int64                                          `json:"totalCount,omitempty"`
	PageNum    int64                                          `json:"pageNum,omitempty"`
}

type SoptaskFetchSopStatisticsPageResponseData2 struct {
}

type SoptaskFetchSopStatisticsPageResponseResults struct {
	PushTime               SoptaskFetchSopStatisticsPageResponsePushTime `json:"pushTime,omitempty"`
	TaskPackageId          string                                        `json:"taskPackageId,omitempty"`
	TransactionOrderNum    int64                                         `json:"transactionOrderNum,omitempty"`
	ExpectTouchNum         int64                                         `json:"expectTouchNum,omitempty"`
	PurchaseCustomerNum    int64                                         `json:"purchaseCustomerNum,omitempty"`
	ExecutedUserNum        int64                                         `json:"executedUserNum,omitempty"`
	OrderedCustomerNum     int64                                         `json:"orderedCustomerNum,omitempty"`
	TaskExecutedRate       string                                        `json:"taskExecutedRate,omitempty"`
	TaskState              int64                                         `json:"taskState,omitempty"`
	Deleted                bool                                          `json:"deleted,omitempty"`
	TransactionCustomerNum int64                                         `json:"transactionCustomerNum,omitempty"`
	TransactionAmount      string                                        `json:"transactionAmount,omitempty"`
	SendType               int64                                         `json:"sendType,omitempty"`
	ActualTouchNum         int64                                         `json:"actualTouchNum,omitempty"`
	TaskCreateTime         string                                        `json:"taskCreateTime,omitempty"`
	TaskName               string                                        `json:"taskName,omitempty"`
	TouchSuccessRate       string                                        `json:"touchSuccessRate,omitempty"`
	ShouldExeUserNum       int64                                         `json:"shouldExeUserNum,omitempty"`
	Class                  string                                        `json:"_class,omitempty"`
	VisitCustomerNum       int64                                         `json:"visitCustomerNum,omitempty"`
}

type SoptaskFetchSopStatisticsPageResponsePushTime struct {
}
