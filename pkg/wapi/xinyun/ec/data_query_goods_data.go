package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DataQueryGoodsData
// @id 1283
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询商品销量数据)
func (client *Client) DataQueryGoodsData(request *DataQueryGoodsDataRequest) (response *DataQueryGoodsDataResponse, err error) {
	rpcResponse := CreateDataQueryGoodsDataResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DataQueryGoodsDataRequest struct {
	*api.BaseRequest

	GoodsIdList  []int64 `json:"goodsIdList,omitempty"`
	StoreId      int64   `json:"storeId,omitempty"`
	QueryPidData bool    `json:"queryPidData,omitempty"`
	BeginDate    int     `json:"beginDate,omitempty"`
	EndDate      int     `json:"endDate,omitempty"`
}

type DataQueryGoodsDataResponse struct {
}

func CreateDataQueryGoodsDataRequest() (request *DataQueryGoodsDataRequest) {
	request = &DataQueryGoodsDataRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "data/queryGoodsData", "api")
	request.Method = api.POST
	return
}

func CreateDataQueryGoodsDataResponse() (response *api.BaseResponse[DataQueryGoodsDataResponse]) {
	response = api.CreateResponse[DataQueryGoodsDataResponse](&DataQueryGoodsDataResponse{})
	return
}
