package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuideGetGuiderTradeList
// @id 1133
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询导购业绩明细)
func (client *Client) GuideGetGuiderTradeList(request *GuideGetGuiderTradeListRequest) (response *GuideGetGuiderTradeListResponse, err error) {
	rpcResponse := CreateGuideGetGuiderTradeListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuideGetGuiderTradeListRequest struct {
	*api.BaseRequest

	PageNum     int    `json:"pageNum,omitempty"`
	PageSize    int    `json:"pageSize,omitempty"`
	GuiderWid   int64  `json:"guiderWid,omitempty"`
	CreateTime  string `json:"createTime,omitempty"`
	EndTime     string `json:"endTime,omitempty"`
	GuiderPhone string `json:"guiderPhone,omitempty"`
	JobNumber   string `json:"jobNumber,omitempty"`
}

type GuideGetGuiderTradeListResponse struct {
}

func CreateGuideGetGuiderTradeListRequest() (request *GuideGetGuiderTradeListRequest) {
	request = &GuideGetGuiderTradeListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "guide/getGuiderTradeList", "api")
	request.Method = api.POST
	return
}

func CreateGuideGetGuiderTradeListResponse() (response *api.BaseResponse[GuideGetGuiderTradeListResponse]) {
	response = api.CreateResponse[GuideGetGuiderTradeListResponse](&GuideGetGuiderTradeListResponse{})
	return
}
