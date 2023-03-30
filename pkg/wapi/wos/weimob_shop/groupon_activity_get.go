package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GrouponActivityGet
// @id 1798
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1798?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询拼团活动的活动规则)
func (client *Client) GrouponActivityGet(request *GrouponActivityGetRequest) (response *GrouponActivityGetResponse, err error) {
	rpcResponse := CreateGrouponActivityGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GrouponActivityGetRequest struct {
	*api.BaseRequest

	BasicInfo            GrouponActivityGetRequestBasicInfo `json:"basicInfo,omitempty"`
	ActivityId           int64                              `json:"activityId,omitempty"`
	SpecificationVersion int64                              `json:"specificationVersion,omitempty"`
	Wid                  int64                              `json:"wid,omitempty"`
}

type GrouponActivityGetResponse struct {
	Data           GrouponActivityGetResponseData    `json:"data,omitempty"`
	ErrData        GrouponActivityGetResponseErrData `json:"errData,omitempty"`
	Errcode        string                            `json:"errcode,omitempty"`
	Errmsg         string                            `json:"errmsg,omitempty"`
	ServerTime     int64                             `json:"serverTime,omitempty"`
	GlobalTicket   string                            `json:"globalTicket,omitempty"`
	MonitorTrackId string                            `json:"monitorTrackId,omitempty"`
}

func CreateGrouponActivityGetRequest() (request *GrouponActivityGetRequest) {
	request = &GrouponActivityGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "groupon/activity/get", "apigw")
	request.Method = api.POST
	return
}

func CreateGrouponActivityGetResponse() (response *api.BaseResponse[GrouponActivityGetResponse]) {
	response = api.CreateResponse[GrouponActivityGetResponse](&GrouponActivityGetResponse{})
	return
}

type GrouponActivityGetRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type GrouponActivityGetResponseData struct {
	MallVersionType      GrouponActivityGetResponseMallVersionType `json:"mallVersionType,omitempty"`
	Cid                  GrouponActivityGetResponseCid             `json:"cid,omitempty"`
	Globalvid            GrouponActivityGetResponseGlobalvid       `json:"globalvid,omitempty"`
	MerchantId           int64                                     `json:"merchantId,omitempty"`
	BosId                int64                                     `json:"bosId,omitempty"`
	Vid                  int64                                     `json:"vid,omitempty"`
	VidType              int64                                     `json:"vidType,omitempty"`
	ProductId            int64                                     `json:"productId,omitempty"`
	ProductVersionId     int64                                     `json:"productVersionId,omitempty"`
	ProductInstanceId    int64                                     `json:"productInstanceId,omitempty"`
	Tcode                string                                    `json:"tcode,omitempty"`
	ActivityId           int64                                     `json:"activityId,omitempty"`
	BelongVid            int64                                     `json:"belongVid,omitempty"`
	BelongVidType        int64                                     `json:"belongVidType,omitempty"`
	Title                string                                    `json:"title,omitempty"`
	StartDate            string                                    `json:"startDate,omitempty"`
	EndDate              string                                    `json:"endDate,omitempty"`
	ActivityType         int64                                     `json:"activityType,omitempty"`
	TagUser              int64                                     `json:"tagUser,omitempty"`
	UserTagContent       string                                    `json:"userTagContent,omitempty"`
	UserTagId            string                                    `json:"userTagId,omitempty"`
	Participators        int64                                     `json:"participators,omitempty"`
	ParticipatorsContent string                                    `json:"participatorsContent,omitempty"`
	IsRecommend          int64                                     `json:"isRecommend,omitempty"`
	DiscountType         string                                    `json:"discountType,omitempty"`
	DurationTime         int64                                     `json:"durationTime,omitempty"`
	OrderAutoCloseTime   int64                                     `json:"orderAutoCloseTime,omitempty"`
	StoreRelationType    int64                                     `json:"storeRelationType,omitempty"`
	StoreRelationList    []int64                                   `json:"storeRelationList,omitempty"`
	IsPreheat            int64                                     `json:"isPreheat,omitempty"`
	PreheatDate          string                                    `json:"preheatDate,omitempty"`
	ImgUrls              string                                    `json:"imgUrls,omitempty"`
	Notes                string                                    `json:"notes,omitempty"`
	OsType               int64                                     `json:"osType,omitempty"`
}

type GrouponActivityGetResponseMallVersionType struct {
}

type GrouponActivityGetResponseCid struct {
}

type GrouponActivityGetResponseGlobalvid struct {
}

type GrouponActivityGetResponseErrData struct {
}
