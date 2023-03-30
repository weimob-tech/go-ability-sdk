package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CardTemplateAdd
// @id 855
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增卡券模板)
func (client *Client) CardTemplateAdd(request *CardTemplateAddRequest) (response *CardTemplateAddResponse, err error) {
	rpcResponse := CreateCardTemplateAddResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CardTemplateAddRequest struct {
	*api.BaseRequest

	TextImages              CardTemplateAddRequestTextImages `json:"textImages,omitempty"`
	AdaptGoods              []map[string]any                 `json:"adaptGoods,omitempty"`
	Name                    string                           `json:"name,omitempty"`
	SubName                 string                           `json:"subName,omitempty"`
	ColorType               int                              `json:"colorType,omitempty"`
	Type                    int                              `json:"type,omitempty"`
	ExpireDateType          int                              `json:"expireDateType,omitempty"`
	StartDateStr            string                           `json:"startDateStr,omitempty"`
	ExpDateStr              string                           `json:"expDateStr,omitempty"`
	StartDayCount           int                              `json:"startDayCount,omitempty"`
	ExpDayCount             int                              `json:"expDayCount,omitempty"`
	UsingTimeType           int                              `json:"usingTimeType,omitempty"`
	WeekDays                []string                         `json:"weekDays,omitempty"`
	StartTimeStr            string                           `json:"startTimeStr,omitempty"`
	EndTimeStr              string                           `json:"endTimeStr,omitempty"`
	TotalCount              int                              `json:"totalCount,omitempty"`
	UserTakeLimit           int                              `json:"userTakeLimit,omitempty"`
	CanUseWithOtherDiscount int                              `json:"canUseWithOtherDiscount,omitempty"`
	ApplyStoreType          int                              `json:"applyStoreType,omitempty"`
	Store                   []string                         `json:"store,omitempty"`
	UserGuide               string                           `json:"userGuide,omitempty"`
	UseNotice               string                           `json:"useNotice,omitempty"`
	ServicePhoneNo          string                           `json:"servicePhoneNo,omitempty"`
	IconUrl                 string                           `json:"iconUrl,omitempty"`
	ImageUrl                string                           `json:"image_url,omitempty"`
	Text                    string                           `json:"text,omitempty"`
	CenterScene             int                              `json:"centerScene,omitempty"`
	CenterTitle             string                           `json:"centerTitle,omitempty"`
	CenterSubTitle          string                           `json:"centerSubTitle,omitempty"`
	CenterUrl               string                           `json:"centerUrl,omitempty"`
	UsingScene              int                              `json:"usingScene,omitempty"`
	UsingSceneName          string                           `json:"usingSceneName,omitempty"`
	UsingGuide              string                           `json:"usingGuide,omitempty"`
	UsingLink               string                           `json:"usingLink,omitempty"`
	MarketingScene          int                              `json:"marketingScene,omitempty"`
	MarketingSceneName      string                           `json:"marketingSceneName,omitempty"`
	MarketingGuide          string                           `json:"marketingGuide,omitempty"`
	MarketingLink           string                           `json:"marketingLink,omitempty"`
	Detail                  string                           `json:"detail,omitempty"`
	CashTicketCondition     float32                          `json:"cashTicketCondition,omitempty"`
	CashTicketAmt           float32                          `json:"cashTicketAmt,omitempty"`
	ApplyCategoryOrNot      int                              `json:"applyCategoryOrNot,omitempty"`
	AcceptCategory          string                           `json:"acceptCategory,omitempty"`
	RejectCategory          string                           `json:"rejectCategory,omitempty"`
	Discount                float32                          `json:"discount,omitempty"`
	ObjectUseFor            string                           `json:"objectUseFor,omitempty"`
	ServiceType             string                           `json:"serviceType,omitempty"`
	SkuId                   int64                            `json:"skuId,omitempty"`
	SkuNum                  int                              `json:"skuNum,omitempty"`
	MaxDiscountAmount       float64                          `json:"maxDiscountAmount,omitempty"`
	OpenSuperposition       int                              `json:"openSuperposition,omitempty"`
	SuperpositionNum        int                              `json:"superpositionNum,omitempty"`
	UseLimitType            int                              `json:"useLimitType,omitempty"`
	UseLimitCycle           int                              `json:"useLimitCycle,omitempty"`
	UseLimitNum             int                              `json:"useLimitNum,omitempty"`
}

type CardTemplateAddResponse struct {
}

func CreateCardTemplateAddRequest() (request *CardTemplateAddRequest) {
	request = &CardTemplateAddRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "cardTemplate/add", "api")
	request.Method = api.POST
	return
}

func CreateCardTemplateAddResponse() (response *api.BaseResponse[CardTemplateAddResponse]) {
	response = api.CreateResponse[CardTemplateAddResponse](&CardTemplateAddResponse{})
	return
}

type CardTemplateAddRequestTextImages struct {
}
