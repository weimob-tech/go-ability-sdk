package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DataQueryLenovoTrafficAnalysisData
// @id 2720
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=流量分析)
func (client *Client) DataQueryLenovoTrafficAnalysisData(request *DataQueryLenovoTrafficAnalysisDataRequest) (response *DataQueryLenovoTrafficAnalysisDataResponse, err error) {
	rpcResponse := CreateDataQueryLenovoTrafficAnalysisDataResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DataQueryLenovoTrafficAnalysisDataRequest struct {
	*api.BaseRequest
}

type DataQueryLenovoTrafficAnalysisDataResponse struct {
}

func CreateDataQueryLenovoTrafficAnalysisDataRequest() (request *DataQueryLenovoTrafficAnalysisDataRequest) {
	request = &DataQueryLenovoTrafficAnalysisDataRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "data/queryLenovoTrafficAnalysisData", "api")
	request.Method = api.POST
	return
}

func CreateDataQueryLenovoTrafficAnalysisDataResponse() (response *api.BaseResponse[DataQueryLenovoTrafficAnalysisDataResponse]) {
	response = api.CreateResponse[DataQueryLenovoTrafficAnalysisDataResponse](&DataQueryLenovoTrafficAnalysisDataResponse{})
	return
}
