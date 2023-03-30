package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GrouponGoodsactivityGetlist
// @id 1796
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1796?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询商品适用的拼团活动详情)
func (client *Client) GrouponGoodsactivityGetlist(request *GrouponGoodsactivityGetlistRequest) (response *GrouponGoodsactivityGetlistResponse, err error) {
	rpcResponse := CreateGrouponGoodsactivityGetlistResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GrouponGoodsactivityGetlistRequest struct {
	*api.BaseRequest

	BasicInfo            GrouponGoodsactivityGetlistRequestBasicInfo `json:"basicInfo,omitempty"`
	GoodsIdList          []int64                                     `json:"goodsIdList,omitempty"`
	SpecificationVersion int64                                       `json:"specificationVersion,omitempty"`
	Wid                  int64                                       `json:"wid,omitempty"`
}

type GrouponGoodsactivityGetlistResponse struct {
	Data           []GrouponGoodsactivityGetlistResponseData  `json:"data,omitempty"`
	ErrData        GrouponGoodsactivityGetlistResponseErrData `json:"errData,omitempty"`
	Errcode        string                                     `json:"errcode,omitempty"`
	Errmsg         string                                     `json:"errmsg,omitempty"`
	ServerTime     int64                                      `json:"serverTime,omitempty"`
	GlobalTicket   string                                     `json:"globalTicket,omitempty"`
	MonitorTrackId string                                     `json:"monitorTrackId,omitempty"`
}

func CreateGrouponGoodsactivityGetlistRequest() (request *GrouponGoodsactivityGetlistRequest) {
	request = &GrouponGoodsactivityGetlistRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "groupon/goodsactivity/getlist", "apigw")
	request.Method = api.POST
	return
}

func CreateGrouponGoodsactivityGetlistResponse() (response *api.BaseResponse[GrouponGoodsactivityGetlistResponse]) {
	response = api.CreateResponse[GrouponGoodsactivityGetlistResponse](&GrouponGoodsactivityGetlistResponse{})
	return
}

type GrouponGoodsactivityGetlistRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type GrouponGoodsactivityGetlistResponseData struct {
	MerchantId         int64  `json:"merchantId,omitempty"`
	BosId              int64  `json:"bosId,omitempty"`
	Vid                int64  `json:"vid,omitempty"`
	VidType            int64  `json:"vidType,omitempty"`
	ProductId          int64  `json:"productId,omitempty"`
	ProductVersionId   int64  `json:"productVersionId,omitempty"`
	ProductInstanceId  int64  `json:"productInstanceId,omitempty"`
	MallVersionType    int64  `json:"mallVersionType,omitempty"`
	Cid                int64  `json:"cid,omitempty"`
	Globalvid          int64  `json:"globalvid,omitempty"`
	Tcode              string `json:"tcode,omitempty"`
	ActivityId         int64  `json:"activityId,omitempty"`
	GoodsId            int64  `json:"goodsId,omitempty"`
	ActivityName       string `json:"activityName,omitempty"`
	BelongVid          int64  `json:"belongVid,omitempty"`
	BelongVidType      int64  `json:"belongVidType,omitempty"`
	CreateTime         string `json:"createTime,omitempty"`
	StartTime          string `json:"startTime,omitempty"`
	EndTime            string `json:"endTime,omitempty"`
	ActivityStatusDesc string `json:"activityStatusDesc,omitempty"`
	ActivityStatus     int64  `json:"activityStatus,omitempty"`
}

type GrouponGoodsactivityGetlistResponseErrData struct {
}
