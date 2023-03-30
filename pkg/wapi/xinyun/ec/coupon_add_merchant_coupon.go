package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponAddMerchantCoupon
// @id 1112
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新建优惠券模板)
func (client *Client) CouponAddMerchantCoupon(request *CouponAddMerchantCouponRequest) (response *CouponAddMerchantCouponResponse, err error) {
	rpcResponse := CreateCouponAddMerchantCouponResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponAddMerchantCouponRequest struct {
	*api.BaseRequest

	AcceptGoodsIds        []map[string]any                           `json:"acceptGoodsIds,omitempty"`
	SelectStoreIdList     []int                                      `json:"selectStoreIdList,omitempty"`
	UseSceneList          []int                                      `json:"useSceneList,omitempty"`
	MerchantIssueChannels []map[string]any                           `json:"merchantIssueChannels,omitempty"`
	CodeRule              CouponAddMerchantCouponRequestCodeRule     `json:"codeRule,omitempty"`
	TextImages            []CouponAddMerchantCouponRequestTextImages `json:"textImages,omitempty"`
	ExtraData             CouponAddMerchantCouponRequestExtraData    `json:"extraData,omitempty"`
	LimitPromotionType    []map[string]any                           `json:"limitPromotionType,omitempty"`
	Name                  string                                     `json:"name,omitempty"`
	CouponDesc            string                                     `json:"couponDesc,omitempty"`
	Type                  int                                        `json:"type,omitempty"`
	CashTicketCondition   float64                                    `json:"cashTicketCondition,omitempty"`
	CashTicketAmt         float64                                    `json:"cashTicketAmt,omitempty"`
	Discount              float64                                    `json:"discount,omitempty"`
	SelectLimitType       int                                        `json:"selectLimitType,omitempty"`
	MaxDiscountAmount     float64                                    `json:"maxDiscountAmount,omitempty"`
	MaxGoodsAmount        int                                        `json:"maxGoodsAmount,omitempty"`
	ExpireDateType        int                                        `json:"expireDateType,omitempty"`
	StartDate             int                                        `json:"startDate,omitempty"`
	ExpireDate            int                                        `json:"expireDate,omitempty"`
	StartDayCount         int                                        `json:"startDayCount,omitempty"`
	ExpDayCount           int                                        `json:"expDayCount,omitempty"`
	Stock                 int                                        `json:"stock,omitempty"`
	AcceptGoodsType       int                                        `json:"acceptGoodsType,omitempty"`
	SelectStoreType       int                                        `json:"selectStoreType,omitempty"`
	StoreLimitNum         int                                        `json:"storeLimitNum,omitempty"`
	UserTakeLimit         int                                        `json:"userTakeLimit,omitempty"`
	ColorType             int                                        `json:"colorType,omitempty"`
	UserGuide             string                                     `json:"userGuide,omitempty"`
	ActivityIssueFlag     int                                        `json:"activityIssueFlag,omitempty"`
	MerchantIssueFlag     int                                        `json:"merchantIssueFlag,omitempty"`
	RecommendIssueFlag    int                                        `json:"recommendIssueFlag,omitempty"`
	RecommendStartTime    int                                        `json:"recommendStartTime,omitempty"`
	RecommendEndTime      int                                        `json:"recommendEndTime,omitempty"`
	ReceiveDesc           string                                     `json:"receiveDesc,omitempty"`
	Detail                string                                     `json:"detail,omitempty"`
	UseNotice             string                                     `json:"useNotice,omitempty"`
	CanGiveFriend         bool                                       `json:"canGiveFriend,omitempty"`
	CanShare              bool                                       `json:"canShare,omitempty"`
	IssueChannel          int                                        `json:"issueChannel,omitempty"`
	MinDiscountGoodsNum   int                                        `json:"minDiscountGoodsNum,omitempty"`
	MaxDiscountGoodsNum   int                                        `json:"maxDiscountGoodsNum,omitempty"`
	CanWeimobEdit         bool                                       `json:"canWeimobEdit,omitempty"`
	SyncWeChatCardBag     bool                                       `json:"syncWeChatCardBag,omitempty"`
	MerchantLogoImg       string                                     `json:"merchantLogoImg,omitempty"`
	MerchantName          string                                     `json:"merchantName,omitempty"`
	SurfaceSetting        bool                                       `json:"surfaceSetting,omitempty"`
	SurfaceImg            string                                     `json:"surfaceImg,omitempty"`
	SurfaceTitle          string                                     `json:"surfaceTitle,omitempty"`
	CanDonate             bool                                       `json:"canDonate,omitempty"`
	ServicePhone          string                                     `json:"servicePhone,omitempty"`
	ThirdId               string                                     `json:"thirdId,omitempty"`
	OperatorId            int64                                      `json:"operatorId,omitempty"`
	SubCouponType         int                                        `json:"subCouponType,omitempty"`
	MinCashTicketAmt      float64                                    `json:"minCashTicketAmt,omitempty"`
	MaxCashTicketAmt      float64                                    `json:"maxCashTicketAmt,omitempty"`
	CreateType            int                                        `json:"createType,omitempty"`
	CreatorId             int64                                      `json:"creatorId,omitempty"`
}

type CouponAddMerchantCouponResponse struct {
	CouponTemplateId int64 `json:"couponTemplateId,omitempty"`
}

func CreateCouponAddMerchantCouponRequest() (request *CouponAddMerchantCouponRequest) {
	request = &CouponAddMerchantCouponRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "coupon/addMerchantCoupon", "api")
	request.Method = api.POST
	return
}

func CreateCouponAddMerchantCouponResponse() (response *api.BaseResponse[CouponAddMerchantCouponResponse]) {
	response = api.CreateResponse[CouponAddMerchantCouponResponse](&CouponAddMerchantCouponResponse{})
	return
}

type CouponAddMerchantCouponRequestCodeRule struct {
	Prefix   string `json:"prefix,omitempty"`
	Suffix   string `json:"suffix,omitempty"`
	MinValue int64  `json:"minValue,omitempty"`
	MaxValue int64  `json:"maxValue,omitempty"`
}

type CouponAddMerchantCouponRequestTextImages struct {
	ImageUrl string `json:"imageUrl,omitempty"`
	Text     string `json:"text,omitempty"`
}

type CouponAddMerchantCouponRequestExtraData struct {
	UseSceneEditable bool `json:"useSceneEditable,omitempty"`
}
