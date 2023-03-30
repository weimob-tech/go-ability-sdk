package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuideQueryStoreTargetList
// @id 1256
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询门店目标列表)
func (client *Client) GuideQueryStoreTargetList(request *GuideQueryStoreTargetListRequest) (response *GuideQueryStoreTargetListResponse, err error) {
	rpcResponse := CreateGuideQueryStoreTargetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuideQueryStoreTargetListRequest struct {
	*api.BaseRequest

	FiscalYear int   `json:"fiscalYear,omitempty"`
	StoreId    int64 `json:"storeId,omitempty"`
	PageNum    int   `json:"pageNum,omitempty"`
	PageSize   int   `json:"pageSize,omitempty"`
}

type GuideQueryStoreTargetListResponse struct {
}

func CreateGuideQueryStoreTargetListRequest() (request *GuideQueryStoreTargetListRequest) {
	request = &GuideQueryStoreTargetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "guide/queryStoreTargetList", "api")
	request.Method = api.POST
	return
}

func CreateGuideQueryStoreTargetListResponse() (response *api.BaseResponse[GuideQueryStoreTargetListResponse]) {
	response = api.CreateResponse[GuideQueryStoreTargetListResponse](&GuideQueryStoreTargetListResponse{})
	return
}
