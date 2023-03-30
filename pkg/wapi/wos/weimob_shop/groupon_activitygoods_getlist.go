package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GrouponActivitygoodsGetlist
// @id 1797
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1797?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询拼团活动适用的商品列表)
func (client *Client) GrouponActivitygoodsGetlist(request *GrouponActivitygoodsGetlistRequest) (response *GrouponActivitygoodsGetlistResponse, err error) {
	rpcResponse := CreateGrouponActivitygoodsGetlistResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GrouponActivitygoodsGetlistRequest struct {
	*api.BaseRequest

	BasicInfo      GrouponActivitygoodsGetlistRequestBasicInfo      `json:"basicInfo,omitempty"`
	QueryParameter GrouponActivitygoodsGetlistRequestQueryParameter `json:"queryParameter,omitempty"`
	PageNum        int64                                            `json:"pageNum,omitempty"`
	PageSize       int64                                            `json:"pageSize,omitempty"`
}

type GrouponActivitygoodsGetlistResponse struct {
	Data           GrouponActivitygoodsGetlistResponseData    `json:"data,omitempty"`
	ErrData        GrouponActivitygoodsGetlistResponseErrData `json:"errData,omitempty"`
	Errcode        string                                     `json:"errcode,omitempty"`
	Errmsg         string                                     `json:"errmsg,omitempty"`
	ServerTime     int64                                      `json:"serverTime,omitempty"`
	GlobalTicket   string                                     `json:"globalTicket,omitempty"`
	MonitorTrackId string                                     `json:"monitorTrackId,omitempty"`
}

func CreateGrouponActivitygoodsGetlistRequest() (request *GrouponActivitygoodsGetlistRequest) {
	request = &GrouponActivitygoodsGetlistRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "groupon/activitygoods/getlist", "apigw")
	request.Method = api.POST
	return
}

func CreateGrouponActivitygoodsGetlistResponse() (response *api.BaseResponse[GrouponActivitygoodsGetlistResponse]) {
	response = api.CreateResponse[GrouponActivitygoodsGetlistResponse](&GrouponActivitygoodsGetlistResponse{})
	return
}

type GrouponActivitygoodsGetlistRequestBasicInfo struct {
	Vid     int64 `json:"vid,omitempty"`
	VidType int64 `json:"vidType,omitempty"`
}

type GrouponActivitygoodsGetlistRequestQueryParameter struct {
	ActivityId  int64   `json:"activityId,omitempty"`
	GoodsIdList []int64 `json:"goodsIdList,omitempty"`
}

type GrouponActivitygoodsGetlistResponseData struct {
	PageList   []GrouponActivitygoodsGetlistResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                         `json:"pageNum,omitempty"`
	PageSize   int64                                         `json:"pageSize,omitempty"`
	TotalCount int64                                         `json:"totalCount,omitempty"`
}

type GrouponActivitygoodsGetlistResponsePageList struct {
	SkuList           []GrouponActivitygoodsGetlistResponseSkuList `json:"skuList,omitempty"`
	MerchantId        int64                                        `json:"merchantId,omitempty"`
	BosId             int64                                        `json:"bosId,omitempty"`
	Vid               int64                                        `json:"vid,omitempty"`
	VidType           int64                                        `json:"vidType,omitempty"`
	ProductId         int64                                        `json:"productId,omitempty"`
	ProductVersionId  int64                                        `json:"productVersionId,omitempty"`
	ProductInstanceId int64                                        `json:"productInstanceId,omitempty"`
	MallVersionType   string                                       `json:"mallVersionType,omitempty"`
	Cid               string                                       `json:"cid,omitempty"`
	Globalvid         string                                       `json:"globalvid,omitempty"`
	Tcode             string                                       `json:"tcode,omitempty"`
	ActivityId        int64                                        `json:"activityId,omitempty"`
	GoodsId           int64                                        `json:"goodsId,omitempty"`
	StockAmount       int64                                        `json:"stockAmount,omitempty"`
	IsLeaderFree      int64                                        `json:"isLeaderFree,omitempty"`
	BuyNum            int64                                        `json:"buyNum,omitempty"`
	GoodsLimit        int64                                        `json:"goodsLimit,omitempty"`
	GrouponPeopleNum  int64                                        `json:"grouponPeopleNum,omitempty"`
	IsAutoGroup       int64                                        `json:"isAutoGroup,omitempty"`
	AutoGroupNum      int64                                        `json:"autoGroupNum,omitempty"`
	InitPeopleNum     int64                                        `json:"initPeopleNum,omitempty"`
	PresentInfo       string                                       `json:"presentInfo,omitempty"`
	Sort              int64                                        `json:"sort,omitempty"`
	ActivityStock     int64                                        `json:"activityStock,omitempty"`
	CreateTime        string                                       `json:"createTime,omitempty"`
	UpdateTime        string                                       `json:"updateTime,omitempty"`
}

type GrouponActivitygoodsGetlistResponseSkuList struct {
	ActivityScopeGoodsSkuStepVOS []GrouponActivitygoodsGetlistResponseActivityScopeGoodsSkuStepVOS `json:"activityScopeGoodsSkuStepVOS,omitempty"`
	GoodsId                      int64                                                             `json:"goodsId,omitempty"`
	SkuId                        int64                                                             `json:"skuId,omitempty"`
	SkuStock                     int64                                                             `json:"skuStock,omitempty"`
	UsedStock                    int64                                                             `json:"usedStock,omitempty"`
}

type GrouponActivitygoodsGetlistResponseActivityScopeGoodsSkuStepVOS struct {
	GrouponPrice   int64 `json:"grouponPrice,omitempty"`
	GroupPeopleNum int64 `json:"groupPeopleNum,omitempty"`
	Step           int64 `json:"step,omitempty"`
}

type GrouponActivitygoodsGetlistResponseErrData struct {
}
