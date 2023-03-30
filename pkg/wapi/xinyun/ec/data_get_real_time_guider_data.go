package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DataGetRealTimeGuiderData
// @id 1590
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询导购实时UV数据)
func (client *Client) DataGetRealTimeGuiderData(request *DataGetRealTimeGuiderDataRequest) (response *DataGetRealTimeGuiderDataResponse, err error) {
	rpcResponse := CreateDataGetRealTimeGuiderDataResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DataGetRealTimeGuiderDataRequest struct {
	*api.BaseRequest

	GuiderId     int64 `json:"guiderId,omitempty"`
	EcBizStoreId int64 `json:"ecBizStoreId,omitempty"`
	PageNum      int   `json:"pageNum,omitempty"`
	PageSize     int   `json:"pageSize,omitempty"`
}

type DataGetRealTimeGuiderDataResponse struct {
}

func CreateDataGetRealTimeGuiderDataRequest() (request *DataGetRealTimeGuiderDataRequest) {
	request = &DataGetRealTimeGuiderDataRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "data/getRealTimeGuiderData", "api")
	request.Method = api.POST
	return
}

func CreateDataGetRealTimeGuiderDataResponse() (response *api.BaseResponse[DataGetRealTimeGuiderDataResponse]) {
	response = api.CreateResponse[DataGetRealTimeGuiderDataResponse](&DataGetRealTimeGuiderDataResponse{})
	return
}
