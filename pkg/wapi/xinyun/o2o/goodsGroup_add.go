package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsGroupAdd
// @id 58
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=添加菜品分组)
func (client *Client) GoodsGroupAdd(request *GoodsGroupAddRequest) (response *GoodsGroupAddResponse, err error) {
	rpcResponse := CreateGoodsGroupAddResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsGroupAddRequest struct {
	*api.BaseRequest

	SellTimes         []GoodsGroupAddRequestSellTimes `json:"sellTimes,omitempty"`
	GroupName         string                          `json:"groupName,omitempty"`
	SortNo            int                             `json:"sortNo,omitempty"`
	StoreId           int64                           `json:"storeId,omitempty"`
	MerchantId        int64                           `json:"merchantId,omitempty"`
	SuppOrderType     int                             `json:"suppOrderType,omitempty"`
	RequireChoose     int                             `json:"requireChoose,omitempty"`
	RequireLeastCount int                             `json:"requireLeastCount,omitempty"`
	ThirdGroupId      string                          `json:"thirdGroupId,omitempty"`
}

type GoodsGroupAddResponse struct {
}

func CreateGoodsGroupAddRequest() (request *GoodsGroupAddRequest) {
	request = &GoodsGroupAddRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goodsGroup/add", "api")
	request.Method = api.POST
	return
}

func CreateGoodsGroupAddResponse() (response *api.BaseResponse[GoodsGroupAddResponse]) {
	response = api.CreateResponse[GoodsGroupAddResponse](&GoodsGroupAddResponse{})
	return
}

type GoodsGroupAddRequestSellTimes struct {
	WeekDay  []map[string]any               `json:"weekDay,omitempty"`
	TimeList []GoodsGroupAddRequestTimeList `json:"timeList,omitempty"`
}

type GoodsGroupAddRequestTimeList struct {
	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
}
