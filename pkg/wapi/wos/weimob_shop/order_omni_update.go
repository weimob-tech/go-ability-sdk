package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderOmniUpdate
// @id 1986
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1986?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=外部订单更新)
func (client *Client) OrderOmniUpdate(request *OrderOmniUpdateRequest) (response *OrderOmniUpdateResponse, err error) {
	rpcResponse := CreateOrderOmniUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderOmniUpdateRequest struct {
	*api.BaseRequest

	CancelInfo     OrderOmniUpdateRequestCancelInfo       `json:"cancelInfo,omitempty"`
	OrderTimeInfos []OrderOmniUpdateRequestOrderTimeInfos `json:"orderTimeInfos,omitempty"`
	PackageList    []OrderOmniUpdateRequestPackageList    `json:"packageList,omitempty"`
	PayFinishInfos []OrderOmniUpdateRequestPayFinishInfos `json:"payFinishInfos,omitempty"`
	OrderNo        int64                                  `json:"orderNo,omitempty"`
	OrderStatus    int64                                  `json:"orderStatus,omitempty"`
}

type OrderOmniUpdateResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateOrderOmniUpdateRequest() (request *OrderOmniUpdateRequest) {
	request = &OrderOmniUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "order/omni/update", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderOmniUpdateResponse() (response *api.BaseResponse[OrderOmniUpdateResponse]) {
	response = api.CreateResponse[OrderOmniUpdateResponse](&OrderOmniUpdateResponse{})
	return
}

type OrderOmniUpdateRequestCancelInfo struct {
	CancelType    int64  `json:"cancelType,omitempty"`
	Reason        string `json:"reason,omitempty"`
	SpecialReason string `json:"specialReason,omitempty"`
}

type OrderOmniUpdateRequestOrderTimeInfos struct {
	Type  int64 `json:"type,omitempty"`
	Value int64 `json:"value,omitempty"`
}

type OrderOmniUpdateRequestPackageList struct {
	DeliveryImportInfo OrderOmniUpdateRequestDeliveryImportInfo `json:"deliveryImportInfo,omitempty"`
	PackageItems       []OrderOmniUpdateRequestPackageItems     `json:"packageItems,omitempty"`
	ReceiveInfo        OrderOmniUpdateRequestReceiveInfo        `json:"receiveInfo,omitempty"`
	SendInfo           OrderOmniUpdateRequestSendInfo           `json:"sendInfo,omitempty"`
	BuyerRemark        string                                   `json:"buyerRemark,omitempty"`
	ConfirmType        int64                                    `json:"confirmType,omitempty"`
	DeliveryMethod     int64                                    `json:"deliveryMethod,omitempty"`
	DeliveryTime       int64                                    `json:"deliveryTime,omitempty"`
	FulfillNo          int64                                    `json:"fulfillNo,omitempty"`
	PackageName        string                                   `json:"packageName,omitempty"`
	ReceiveTime        int64                                    `json:"receiveTime,omitempty"`
	Remark             string                                   `json:"remark,omitempty"`
}

type OrderOmniUpdateRequestDeliveryImportInfo struct {
	CompanyCode             string `json:"companyCode,omitempty"`
	CompanyName             string `json:"companyName,omitempty"`
	ExpectReceivedDate      int64  `json:"expectReceivedDate,omitempty"`
	ExpectReceivedEndTime   int64  `json:"expectReceivedEndTime,omitempty"`
	ExpectReceivedStartTime int64  `json:"expectReceivedStartTime,omitempty"`
	ExpectReceivedType      int64  `json:"expectReceivedType,omitempty"`
	Number                  string `json:"number,omitempty"`
	WriteOffName            string `json:"writeOffName,omitempty"`
	WriteOffId              int64  `json:"writeOffId,omitempty"`
}

type OrderOmniUpdateRequestPackageItems struct {
	OutItemId string `json:"outItemId,omitempty"`
	SkuNum    string `json:"skuNum,omitempty"`
}

type OrderOmniUpdateRequestReceiveInfo struct {
	AddressInfo     OrderOmniUpdateRequestAddressInfo  `json:"addressInfo,omitempty"`
	ReceiverInfo    OrderOmniUpdateRequestReceiverInfo `json:"receiverInfo,omitempty"`
	PickUpName      string                             `json:"pickUpName,omitempty"`
	PickUpVid       int64                              `json:"pickUpVid,omitempty"`
	ReceiverAddress string                             `json:"receiverAddress,omitempty"`
}

type OrderOmniUpdateRequestAddressInfo struct {
	AddressExt OrderOmniUpdateRequestAddressExt `json:"addressExt,omitempty"`
	Address    string                           `json:"address,omitempty"`
	Area       string                           `json:"area,omitempty"`
	City       string                           `json:"city,omitempty"`
	County     string                           `json:"county,omitempty"`
	Latitude   string                           `json:"latitude,omitempty"`
	Longitude  string                           `json:"longitude,omitempty"`
	Province   string                           `json:"province,omitempty"`
	Zip        string                           `json:"zip,omitempty"`
}

type OrderOmniUpdateRequestAddressExt struct {
	AreaCode     string `json:"areaCode,omitempty"`
	CityCode     string `json:"cityCode,omitempty"`
	CountyCode   string `json:"countyCode,omitempty"`
	ProvinceCode string `json:"provinceCode,omitempty"`
}

type OrderOmniUpdateRequestReceiverInfo struct {
	IdCardExt      OrderOmniUpdateRequestIdCardExt `json:"idCardExt,omitempty"`
	IdCardNo       string                          `json:"idCardNo,omitempty"`
	ReceiverMobile string                          `json:"receiverMobile,omitempty"`
	ReceiverName   string                          `json:"receiverName,omitempty"`
}

type OrderOmniUpdateRequestIdCardExt struct {
	BehindImg string `json:"behindImg,omitempty"`
	FrontImg  string `json:"frontImg,omitempty"`
	UserName  string `json:"userName,omitempty"`
}

type OrderOmniUpdateRequestSendInfo struct {
	AddressInfo    OrderOmniUpdateRequestAddressInfo2 `json:"addressInfo,omitempty"`
	SenderInfo     OrderOmniUpdateRequestSenderInfo   `json:"senderInfo,omitempty"`
	ProcessVid     int64                              `json:"processVid,omitempty"`
	ProcessVidName string                             `json:"processVidName,omitempty"`
	SenderAddress  string                             `json:"senderAddress,omitempty"`
}

type OrderOmniUpdateRequestAddressInfo2 struct {
	AddressExt OrderOmniUpdateRequestAddressExt2 `json:"addressExt,omitempty"`
	Address    string                            `json:"address,omitempty"`
	Area       string                            `json:"area,omitempty"`
	City       string                            `json:"city,omitempty"`
	County     string                            `json:"county,omitempty"`
	Latitude   string                            `json:"latitude,omitempty"`
	Longitude  string                            `json:"longitude,omitempty"`
	Province   string                            `json:"province,omitempty"`
	Zip        string                            `json:"zip,omitempty"`
}

type OrderOmniUpdateRequestAddressExt2 struct {
	AreaCode     string `json:"areaCode,omitempty"`
	CityCode     string `json:"cityCode,omitempty"`
	CountyCode   string `json:"countyCode,omitempty"`
	ProvinceCode string `json:"provinceCode,omitempty"`
}

type OrderOmniUpdateRequestSenderInfo struct {
	SenderMobile string `json:"senderMobile,omitempty"`
	SenderName   string `json:"senderName,omitempty"`
}

type OrderOmniUpdateRequestPayFinishInfos struct {
	ChannelTrxNo string `json:"channelTrxNo,omitempty"`
	PayAmount    string `json:"payAmount,omitempty"`
	PayMethod    int64  `json:"payMethod,omitempty"`
	PayTime      int64  `json:"payTime,omitempty"`
}
