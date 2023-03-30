package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DataQueryMerchantRealTimeData
// @id 1522
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询店铺实时概览数据)
func (client *Client) DataQueryMerchantRealTimeData(request *DataQueryMerchantRealTimeDataRequest) (response *DataQueryMerchantRealTimeDataResponse, err error) {
	rpcResponse := CreateDataQueryMerchantRealTimeDataResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DataQueryMerchantRealTimeDataRequest struct {
	*api.BaseRequest
}

type DataQueryMerchantRealTimeDataResponse struct {
}

func CreateDataQueryMerchantRealTimeDataRequest() (request *DataQueryMerchantRealTimeDataRequest) {
	request = &DataQueryMerchantRealTimeDataRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "data/queryMerchantRealTimeData", "api")
	request.Method = api.POST
	return
}

func CreateDataQueryMerchantRealTimeDataResponse() (response *api.BaseResponse[DataQueryMerchantRealTimeDataResponse]) {
	response = api.CreateResponse[DataQueryMerchantRealTimeDataResponse](&DataQueryMerchantRealTimeDataResponse{})
	return
}
