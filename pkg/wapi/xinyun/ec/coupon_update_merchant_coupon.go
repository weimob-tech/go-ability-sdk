package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponUpdateMerchantCoupon
// @id 1694
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新优惠券模板)
func (client *Client) CouponUpdateMerchantCoupon(request *CouponUpdateMerchantCouponRequest) (response *CouponUpdateMerchantCouponResponse, err error) {
	rpcResponse := CreateCouponUpdateMerchantCouponResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponUpdateMerchantCouponRequest struct {
	*api.BaseRequest

	AcceptGoodsIds        []string                                      `json:"acceptGoodsIds,omitempty"`
	SelectStoreIdList     []string                                      `json:"selectStoreIdList,omitempty"`
	UseSceneList          []string                                      `json:"useSceneList,omitempty"`
	LimitPromotionType    []string                                      `json:"limitPromotionType,omitempty"`
	CanDeductionType      []string                                      `json:"canDeductionType,omitempty"`
	MerchantIssueChannels []string                                      `json:"merchantIssueChannels,omitempty"`
	TextImages            []CouponUpdateMerchantCouponRequestTextImages `json:"textImages,omitempty"`
	ExtraData             CouponUpdateMerchantCouponRequestExtraData    `json:"extraData,omitempty"`
	StoreId               int64                                         `json:"storeId,omitempty"`
	CardTemplateId        int64                                         `json:"cardTemplateId,omitempty"`
	CouponDesc            string                                        `json:"couponDesc,omitempty"`
	ExpireDate            int                                           `json:"expireDate,omitempty"`
	CouponCostType        int                                           `json:"couponCostType,omitempty"`
	CouponCost            float64                                       `json:"couponCost,omitempty"`
	AcceptGoodsType       int                                           `json:"acceptGoodsType,omitempty"`
	SelectStoreType       int                                           `json:"selectStoreType,omitempty"`
	ActivityIssueFlag     int                                           `json:"activityIssueFlag,omitempty"`
	MerchantIssueFlag     int                                           `json:"merchantIssueFlag,omitempty"`
	StoreLimitNum         int                                           `json:"storeLimitNum,omitempty"`
	RecommendIssueFlag    int                                           `json:"recommendIssueFlag,omitempty"`
	RecommendStartTime    int                                           `json:"recommendStartTime,omitempty"`
	RecommendEndTime      int                                           `json:"recommendEndTime,omitempty"`
	UserTakeLimit         int                                           `json:"userTakeLimit,omitempty"`
	StoreLaunchFlag       bool                                          `json:"storeLaunchFlag,omitempty"`
	CanShare              bool                                          `json:"canShare,omitempty"`
	CanGiveFriend         bool                                          `json:"canGiveFriend,omitempty"`
	CanWeimobEdit         bool                                          `json:"canWeimobEdit,omitempty"`
	ReceiveDesc           string                                        `json:"receiveDesc,omitempty"`
	UseNotice             string                                        `json:"useNotice,omitempty"`
	SyncWeChatCardBag     bool                                          `json:"syncWeChatCardBag,omitempty"`
	ColorType             int                                           `json:"colorType,omitempty"`
	MerchantLogoImg       string                                        `json:"merchantLogoImg,omitempty"`
	MerchantName          string                                        `json:"merchantName,omitempty"`
	SurfaceSetting        bool                                          `json:"surfaceSetting,omitempty"`
	SurfaceImg            string                                        `json:"surfaceImg,omitempty"`
	SurfaceTitle          string                                        `json:"surfaceTitle,omitempty"`
	CanDonate             bool                                          `json:"canDonate,omitempty"`
	UserGuide             string                                        `json:"userGuide,omitempty"`
	ServicePhone          string                                        `json:"servicePhone,omitempty"`
	OperatorId            int64                                         `json:"operatorId,omitempty"`
}

type CouponUpdateMerchantCouponResponse struct {
}

func CreateCouponUpdateMerchantCouponRequest() (request *CouponUpdateMerchantCouponRequest) {
	request = &CouponUpdateMerchantCouponRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "coupon/updateMerchantCoupon", "api")
	request.Method = api.POST
	return
}

func CreateCouponUpdateMerchantCouponResponse() (response *api.BaseResponse[CouponUpdateMerchantCouponResponse]) {
	response = api.CreateResponse[CouponUpdateMerchantCouponResponse](&CouponUpdateMerchantCouponResponse{})
	return
}

type CouponUpdateMerchantCouponRequestTextImages struct {
	ImageUrl string `json:"imageUrl,omitempty"`
	Text     string `json:"text,omitempty"`
}

type CouponUpdateMerchantCouponRequestExtraData struct {
	UseSceneEditable bool `json:"useSceneEditable,omitempty"`
}
