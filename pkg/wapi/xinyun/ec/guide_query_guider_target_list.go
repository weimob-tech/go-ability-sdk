package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuideQueryGuiderTargetList
// @id 1255
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询导购目标列表)
func (client *Client) GuideQueryGuiderTargetList(request *GuideQueryGuiderTargetListRequest) (response *GuideQueryGuiderTargetListResponse, err error) {
	rpcResponse := CreateGuideQueryGuiderTargetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuideQueryGuiderTargetListRequest struct {
	*api.BaseRequest

	FiscalYear  int    `json:"fiscalYear,omitempty"`
	StoreId     int64  `json:"storeId,omitempty"`
	GuiderWid   int64  `json:"guiderWid,omitempty"`
	GuiderPhone string `json:"guiderPhone,omitempty"`
	JobNumber   string `json:"jobNumber,omitempty"`
	PageNum     int    `json:"pageNum,omitempty"`
	PageSize    int    `json:"pageSize,omitempty"`
}

type GuideQueryGuiderTargetListResponse struct {
}

func CreateGuideQueryGuiderTargetListRequest() (request *GuideQueryGuiderTargetListRequest) {
	request = &GuideQueryGuiderTargetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "guide/queryGuiderTargetList", "api")
	request.Method = api.POST
	return
}

func CreateGuideQueryGuiderTargetListResponse() (response *api.BaseResponse[GuideQueryGuiderTargetListResponse]) {
	response = api.CreateResponse[GuideQueryGuiderTargetListResponse](&GuideQueryGuiderTargetListResponse{})
	return
}
