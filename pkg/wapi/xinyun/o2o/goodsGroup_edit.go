package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsGroupEdit
// @id 59
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改菜品分组)
func (client *Client) GoodsGroupEdit(request *GoodsGroupEditRequest) (response *GoodsGroupEditResponse, err error) {
	rpcResponse := CreateGoodsGroupEditResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsGroupEditRequest struct {
	*api.BaseRequest

	SellTimes         []GoodsGroupEditRequestSellTimes `json:"sellTimes,omitempty"`
	GoodsGroupId      int64                            `json:"goodsGroupId,omitempty"`
	GroupName         string                           `json:"groupName,omitempty"`
	SortNo            int                              `json:"sortNo,omitempty"`
	StoreId           int64                            `json:"storeId,omitempty"`
	MerchantId        int64                            `json:"merchantId,omitempty"`
	RequireChoose     int                              `json:"requireChoose,omitempty"`
	RequireLeastCount int                              `json:"requireLeastCount,omitempty"`
}

type GoodsGroupEditResponse struct {
}

func CreateGoodsGroupEditRequest() (request *GoodsGroupEditRequest) {
	request = &GoodsGroupEditRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goodsGroup/edit", "api")
	request.Method = api.POST
	return
}

func CreateGoodsGroupEditResponse() (response *api.BaseResponse[GoodsGroupEditResponse]) {
	response = api.CreateResponse[GoodsGroupEditResponse](&GoodsGroupEditResponse{})
	return
}

type GoodsGroupEditRequestSellTimes struct {
	WeekDay  []map[string]any                `json:"weekDay,omitempty"`
	TimeList []GoodsGroupEditRequestTimeList `json:"timeList,omitempty"`
}

type GoodsGroupEditRequestTimeList struct {
	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
}
