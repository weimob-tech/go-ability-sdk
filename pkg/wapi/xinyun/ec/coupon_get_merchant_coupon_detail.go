package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponGetMerchantCouponDetail
// @id 1079
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取优惠券模板详情)
func (client *Client) CouponGetMerchantCouponDetail(request *CouponGetMerchantCouponDetailRequest) (response *CouponGetMerchantCouponDetailResponse, err error) {
	rpcResponse := CreateCouponGetMerchantCouponDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponGetMerchantCouponDetailRequest struct {
	*api.BaseRequest

	StoreId        int64 `json:"storeId,omitempty"`
	CardTemplateId int64 `json:"cardTemplateId,omitempty"`
}

type CouponGetMerchantCouponDetailResponse struct {
	SceneList           []map[string]any `json:"sceneList,omitempty"`
	MerchantName        string           `json:"merchantName,omitempty"`
	Name                string           `json:"name,omitempty"`
	SubName             string           `json:"subName,omitempty"`
	UseNotice           string           `json:"useNotice,omitempty"`
	UseRule             string           `json:"useRule,omitempty"`
	UserGuide           string           `json:"userGuide,omitempty"`
	CardTemplateId      int64            `json:"cardTemplateId,omitempty"`
	IconUrl             string           `json:"iconUrl,omitempty"`
	Detail              string           `json:"detail,omitempty"`
	Status              int              `json:"status,omitempty"`
	Type                int              `json:"type,omitempty"`
	IsSyncWechat        int              `json:"isSyncWechat,omitempty"`
	ExpireDateType      int              `json:"expireDateType,omitempty"`
	StartDate           int64            `json:"startDate,omitempty"`
	ExpireDate          int64            `json:"expireDate,omitempty"`
	CreateDate          int64            `json:"createDate,omitempty"`
	Discount            float64          `json:"discount,omitempty"`
	CashTicketAmt       float64          `json:"cashTicketAmt,omitempty"`
	CashTicketCondition float64          `json:"cashTicketCondition,omitempty"`
	CreateStoreId       int64            `json:"createStoreId,omitempty"`
	PublishChannels     []int            `json:"publishChannels,omitempty"`
	UseSceneList        []int            `json:"useSceneList,omitempty"`
	ColorType           int              `json:"colorType,omitempty"`
	Left                int64            `json:"left,omitempty"`
	ServicePhone        string           `json:"servicePhone,omitempty"`
	UserTakeLimit       int              `json:"userTakeLimit,omitempty"`
	AcceptGoodsType     int              `json:"acceptGoodsType,omitempty"`
	SelectStoreType     int              `json:"selectStoreType,omitempty"`
	SelectStoreIdList   []int64          `json:"selectStoreIdList,omitempty"`
	StoreLimitNum       int              `json:"storeLimitNum,omitempty"`
	ScenePageId         int              `json:"scenePageId,omitempty"`
	Index               int              `json:"index,omitempty"`
	SceneId             int64            `json:"sceneId,omitempty"`
	SceneType           int              `json:"sceneType,omitempty"`
	SceneTitle          string           `json:"sceneTitle,omitempty"`
	IsOpen              int              `json:"isOpen,omitempty"`
	SceneTypeTrunText   string           `json:"sceneTypeTrunText,omitempty"`
	SceneSubTitle       string           `json:"sceneSubTitle,omitempty"`
	H5Url               string           `json:"h5Url,omitempty"`
	AcceptGoodsIdList   []int64          `json:"acceptGoodsIdList,omitempty"`
	StartDayCount       int              `json:"startDayCount,omitempty"`
	ExpDayCount         int              `json:"expDayCount,omitempty"`
	ExistExcludeGoods   bool             `json:"existExcludeGoods,omitempty"`
	ExcludeGoodsList    []int64          `json:"excludeGoodsList,omitempty"`
	CouponDesc          string           `json:"couponDesc,omitempty"`
	LimitPromotionType  []int            `json:"limitPromotionType,omitempty"`
	CanDeductionType    []int            `json:"canDeductionType,omitempty"`
	IssueChannel        int              `json:"issueChannel,omitempty"`
	ActivityIssueFlag   int              `json:"activityIssueFlag,omitempty"`
	MerchantIssueFlag   int              `json:"merchantIssueFlag,omitempty"`
	ReceiveDesc         string           `json:"receiveDesc,omitempty"`
	RecommendStartTime  int              `json:"recommendStartTime,omitempty"`
	RecommendEndTime    int              `json:"recommendEndTime,omitempty"`
	CouponCostType      int              `json:"couponCostType,omitempty"`
	CouponCost          float64          `json:"couponCost,omitempty"`
	AcceptObjectIdList  []int64          `json:"acceptObjectIdList,omitempty"`
	OperatorId          int64            `json:"operatorId,omitempty"`
	ThirdId             string           `json:"thirdId,omitempty"`
	MinDiscountGoodsNum int              `json:"minDiscountGoodsNum,omitempty"`
	MaxDiscountGoodsNum int              `json:"maxDiscountGoodsNum,omitempty"`
	CodeGenerateType    byte             `json:"codeGenerateType,omitempty"`
	CustomCodePrefix    string           `json:"customCodePrefix,omitempty"`
	CustomCodeSuffix    string           `json:"customCodeSuffix,omitempty"`
	CustomCodeMinValue  string           `json:"customCodeMinValue,omitempty"`
	CustomCodeMaxValue  string           `json:"customCodeMaxValue,omitempty"`
	SubCouponType       int              `json:"subCouponType,omitempty"`
	MinCashTicketAmt    float64          `json:"minCashTicketAmt,omitempty"`
	MaxCashTicketAmt    float64          `json:"maxCashTicketAmt,omitempty"`
}

func CreateCouponGetMerchantCouponDetailRequest() (request *CouponGetMerchantCouponDetailRequest) {
	request = &CouponGetMerchantCouponDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "coupon/getMerchantCouponDetail", "api")
	request.Method = api.POST
	return
}

func CreateCouponGetMerchantCouponDetailResponse() (response *api.BaseResponse[CouponGetMerchantCouponDetailResponse]) {
	response = api.CreateResponse[CouponGetMerchantCouponDetailResponse](&CouponGetMerchantCouponDetailResponse{})
	return
}
