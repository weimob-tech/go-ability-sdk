package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DataGetRealTimeStoreData
// @id 1589
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询门店实时uv数据)
func (client *Client) DataGetRealTimeStoreData(request *DataGetRealTimeStoreDataRequest) (response *DataGetRealTimeStoreDataResponse, err error) {
	rpcResponse := CreateDataGetRealTimeStoreDataResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DataGetRealTimeStoreDataRequest struct {
	*api.BaseRequest

	EcBizStoreId int `json:"ecBizStoreId,omitempty"`
	PageNum      int `json:"pageNum,omitempty"`
	PageSize     int `json:"pageSize,omitempty"`
}

type DataGetRealTimeStoreDataResponse struct {
}

func CreateDataGetRealTimeStoreDataRequest() (request *DataGetRealTimeStoreDataRequest) {
	request = &DataGetRealTimeStoreDataRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "data/getRealTimeStoreData", "api")
	request.Method = api.POST
	return
}

func CreateDataGetRealTimeStoreDataResponse() (response *api.BaseResponse[DataGetRealTimeStoreDataResponse]) {
	response = api.CreateResponse[DataGetRealTimeStoreDataResponse](&DataGetRealTimeStoreDataResponse{})
	return
}
