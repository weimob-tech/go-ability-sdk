package sdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// WeikeGetMerchantWeikeList
// @id 742
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取商户微客列表)
func (client *Client) WeikeGetMerchantWeikeList(request *WeikeGetMerchantWeikeListRequest) (response *WeikeGetMerchantWeikeListResponse, err error) {
	rpcResponse := CreateWeikeGetMerchantWeikeListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WeikeGetMerchantWeikeListRequest struct {
	*api.BaseRequest

	QueryPamater WeikeGetMerchantWeikeListRequestQueryPamater `json:"queryPamater,omitempty"`
	PageNum      int64                                        `json:"pageNum,omitempty"`
	PageSize     int64                                        `json:"pageSize,omitempty"`
	StartTime    int64                                        `json:"startTime,omitempty"`
	EndTime      int64                                        `json:"endTime,omitempty"`
}

type WeikeGetMerchantWeikeListResponse struct {
	PageList   []WeikeGetMerchantWeikeListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                       `json:"pageNum,omitempty"`
	PageSize   int64                                       `json:"pageSize,omitempty"`
	TotalCount int64                                       `json:"totalCount,omitempty"`
}

func CreateWeikeGetMerchantWeikeListRequest() (request *WeikeGetMerchantWeikeListRequest) {
	request = &WeikeGetMerchantWeikeListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("sdp", "1_0", "weike/getMerchantWeikeList", "api")
	request.Method = api.POST
	return
}

func CreateWeikeGetMerchantWeikeListResponse() (response *api.BaseResponse[WeikeGetMerchantWeikeListResponse]) {
	response = api.CreateResponse[WeikeGetMerchantWeikeListResponse](&WeikeGetMerchantWeikeListResponse{})
	return
}

type WeikeGetMerchantWeikeListRequestQueryPamater struct {
}

type WeikeGetMerchantWeikeListResponsePageList struct {
	Level        WeikeGetMerchantWeikeListResponseLevel   `json:"level,omitempty"`
	Inviter      WeikeGetMerchantWeikeListResponseInviter `json:"inviter,omitempty"`
	ScrollId     int64                                    `json:"scrollId,omitempty"`
	Wid          int64                                    `json:"wid,omitempty"`
	State        int64                                    `json:"state,omitempty"`
	RegisterTime int64                                    `json:"registerTime,omitempty"`
	Name         string                                   `json:"name,omitempty"`
	Id           int64                                    `json:"id,omitempty"`
}

type WeikeGetMerchantWeikeListResponseLevel struct {
}

type WeikeGetMerchantWeikeListResponseInviter struct {
}
