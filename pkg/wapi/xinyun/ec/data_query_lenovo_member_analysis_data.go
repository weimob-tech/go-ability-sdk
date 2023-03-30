package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DataQueryLenovoMemberAnalysisData
// @id 2719
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=会员分析)
func (client *Client) DataQueryLenovoMemberAnalysisData(request *DataQueryLenovoMemberAnalysisDataRequest) (response *DataQueryLenovoMemberAnalysisDataResponse, err error) {
	rpcResponse := CreateDataQueryLenovoMemberAnalysisDataResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DataQueryLenovoMemberAnalysisDataRequest struct {
	*api.BaseRequest
}

type DataQueryLenovoMemberAnalysisDataResponse struct {
}

func CreateDataQueryLenovoMemberAnalysisDataRequest() (request *DataQueryLenovoMemberAnalysisDataRequest) {
	request = &DataQueryLenovoMemberAnalysisDataRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "data/queryLenovoMemberAnalysisData", "api")
	request.Method = api.POST
	return
}

func CreateDataQueryLenovoMemberAnalysisDataResponse() (response *api.BaseResponse[DataQueryLenovoMemberAnalysisDataResponse]) {
	response = api.CreateResponse[DataQueryLenovoMemberAnalysisDataResponse](&DataQueryLenovoMemberAnalysisDataResponse{})
	return
}
