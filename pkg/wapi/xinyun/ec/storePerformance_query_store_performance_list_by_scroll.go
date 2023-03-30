package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StorePerformanceQueryStorePerformanceListByScroll
// @id 1588
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取门店业绩明细)
func (client *Client) StorePerformanceQueryStorePerformanceListByScroll(request *StorePerformanceQueryStorePerformanceListByScrollRequest) (response *StorePerformanceQueryStorePerformanceListByScrollResponse, err error) {
	rpcResponse := CreateStorePerformanceQueryStorePerformanceListByScrollResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StorePerformanceQueryStorePerformanceListByScrollRequest struct {
	*api.BaseRequest

	StartId         int64  `json:"startId,omitempty"`
	PageSize        int    `json:"pageSize,omitempty"`
	PerformanceType int    `json:"performanceType,omitempty"`
	StoreId         int64  `json:"storeId,omitempty"`
	StartTime       string `json:"startTime,omitempty"`
	EndTime         string `json:"endTime,omitempty"`
}

type StorePerformanceQueryStorePerformanceListByScrollResponse struct {
}

func CreateStorePerformanceQueryStorePerformanceListByScrollRequest() (request *StorePerformanceQueryStorePerformanceListByScrollRequest) {
	request = &StorePerformanceQueryStorePerformanceListByScrollRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "storePerformance/queryStorePerformanceListByScroll", "api")
	request.Method = api.POST
	return
}

func CreateStorePerformanceQueryStorePerformanceListByScrollResponse() (response *api.BaseResponse[StorePerformanceQueryStorePerformanceListByScrollResponse]) {
	response = api.CreateResponse[StorePerformanceQueryStorePerformanceListByScrollResponse](&StorePerformanceQueryStorePerformanceListByScrollResponse{})
	return
}
