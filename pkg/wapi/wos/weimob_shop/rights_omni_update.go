package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsOmniUpdate
// @id 1930
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1930?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=全渠道更新售后单)
func (client *Client) RightsOmniUpdate(request *RightsOmniUpdateRequest) (response *RightsOmniUpdateResponse, err error) {
	rpcResponse := CreateRightsOmniUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsOmniUpdateRequest struct {
	*api.BaseRequest

	RightsOrder          RightsOmniUpdateRequestRightsOrder            `json:"rightsOrder,omitempty"`
	RefundPayInfoList    []RightsOmniUpdateRequestRefundPayInfoList    `json:"refundPayInfoList,omitempty"`
	ReturnAddressInfo    RightsOmniUpdateRequestReturnAddressInfo      `json:"returnAddressInfo,omitempty"`
	RefundAccount        RightsOmniUpdateRequestRefundAccount          `json:"refundAccount,omitempty"`
	ExchangeDeliveryInfo []RightsOmniUpdateRequestExchangeDeliveryInfo `json:"exchangeDeliveryInfo,omitempty"`
	ReturnDeliveryInfo   RightsOmniUpdateRequestReturnDeliveryInfo     `json:"returnDeliveryInfo,omitempty"`
	CancelInfo           RightsOmniUpdateRequestCancelInfo             `json:"cancelInfo,omitempty"`
	RefuseInfo           RightsOmniUpdateRequestRefuseInfo             `json:"refuseInfo,omitempty"`
	RightsReasonInfo     RightsOmniUpdateRequestRightsReasonInfo       `json:"rightsReasonInfo,omitempty"`
}

type RightsOmniUpdateResponse struct {
}

func CreateRightsOmniUpdateRequest() (request *RightsOmniUpdateRequest) {
	request = &RightsOmniUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "rights/omni/update", "apigw")
	request.Method = api.POST
	return
}

func CreateRightsOmniUpdateResponse() (response *api.BaseResponse[RightsOmniUpdateResponse]) {
	response = api.CreateResponse[RightsOmniUpdateResponse](&RightsOmniUpdateResponse{})
	return
}

type RightsOmniUpdateRequestRightsOrder struct {
	Operations   []RightsOmniUpdateRequestOperations `json:"operations,omitempty"`
	RightsId     int64                               `json:"rightsId,omitempty"`
	RightsStatus int64                               `json:"rightsStatus,omitempty"`
	OutRightsId  string                              `json:"outRightsId,omitempty"`
}

type RightsOmniUpdateRequestOperations struct {
	Operator RightsOmniUpdateRequestOperator `json:"operator,omitempty"`
	Datetime int64                           `json:"datetime,omitempty"`
	Type     int64                           `json:"type,omitempty"`
}

type RightsOmniUpdateRequestOperator struct {
	OperatorId    string `json:"operatorId,omitempty"`
	OperatorName  string `json:"operatorName,omitempty"`
	OperatorPhone string `json:"operatorPhone,omitempty"`
}

type RightsOmniUpdateRequestRefundPayInfoList struct {
	RefundStatus       int64  `json:"refundStatus,omitempty"`
	RefundFailedReason string `json:"refundFailedReason,omitempty"`
	Amount             string `json:"amount,omitempty"`
	RefundMethodId     int64  `json:"refundMethodId,omitempty"`
}

type RightsOmniUpdateRequestReturnAddressInfo struct {
	AddressInfo RightsOmniUpdateRequestAddressInfo `json:"addressInfo,omitempty"`
	Name        string                             `json:"name,omitempty"`
	Phone       string                             `json:"phone,omitempty"`
}

type RightsOmniUpdateRequestAddressInfo struct {
	AddressExt RightsOmniUpdateRequestAddressExt `json:"addressExt,omitempty"`
	Longitude  string                            `json:"longitude,omitempty"`
	County     string                            `json:"county,omitempty"`
	Zip        string                            `json:"zip,omitempty"`
	Area       string                            `json:"area,omitempty"`
	Province   string                            `json:"province,omitempty"`
	City       string                            `json:"city,omitempty"`
	Address    string                            `json:"address,omitempty"`
	Latitude   string                            `json:"latitude,omitempty"`
}

type RightsOmniUpdateRequestAddressExt struct {
	AreaCode     string `json:"areaCode,omitempty"`
	ProvinceCode string `json:"provinceCode,omitempty"`
	CountyCode   string `json:"countyCode,omitempty"`
	CityCode     string `json:"cityCode,omitempty"`
}

type RightsOmniUpdateRequestRefundAccount struct {
	RefundMethodId int64  `json:"refundMethodId,omitempty"`
	OpeningBank    string `json:"openingBank,omitempty"`
	AccountNum     string `json:"accountNum,omitempty"`
	UserName       string `json:"userName,omitempty"`
}

type RightsOmniUpdateRequestExchangeDeliveryInfo struct {
	PackageItems   []RightsOmniUpdateRequestPackageItems `json:"packageItems,omitempty"`
	SenderInfo     RightsOmniUpdateRequestSenderInfo     `json:"senderInfo,omitempty"`
	ReceiveInfo    RightsOmniUpdateRequestReceiveInfo    `json:"receiveInfo,omitempty"`
	DeliveryInfo   RightsOmniUpdateRequestDeliveryInfo   `json:"deliveryInfo,omitempty"`
	DeliveryMethod int64                                 `json:"deliveryMethod,omitempty"`
	OutPackageId   string                                `json:"outPackageId,omitempty"`
	Remark         string                                `json:"remark,omitempty"`
}

type RightsOmniUpdateRequestPackageItems struct {
	OutRightsItemId string `json:"outRightsItemId,omitempty"`
	Num             string `json:"num,omitempty"`
}

type RightsOmniUpdateRequestSenderInfo struct {
	Sender         RightsOmniUpdateRequestSender       `json:"sender,omitempty"`
	AddressInfo    RightsOmniUpdateRequestAddressInfo2 `json:"addressInfo,omitempty"`
	ProcessVidName string                              `json:"processVidName,omitempty"`
	ProcessVid     int64                               `json:"processVid,omitempty"`
}

type RightsOmniUpdateRequestSender struct {
	SenderMobile string `json:"senderMobile,omitempty"`
	SenderName   string `json:"senderName,omitempty"`
}

type RightsOmniUpdateRequestAddressInfo2 struct {
	AddressExt RightsOmniUpdateRequestAddressExt2 `json:"addressExt,omitempty"`
	Latitude   string                             `json:"latitude,omitempty"`
	Longitude  string                             `json:"longitude,omitempty"`
	Province   string                             `json:"province,omitempty"`
	Address    string                             `json:"address,omitempty"`
	County     string                             `json:"county,omitempty"`
	Area       string                             `json:"area,omitempty"`
	Zip        string                             `json:"zip,omitempty"`
	City       string                             `json:"city,omitempty"`
}

type RightsOmniUpdateRequestAddressExt2 struct {
	CityCode     string `json:"cityCode,omitempty"`
	CountyCode   string `json:"countyCode,omitempty"`
	AreaCode     string `json:"areaCode,omitempty"`
	ProvinceCode string `json:"provinceCode,omitempty"`
}

type RightsOmniUpdateRequestReceiveInfo struct {
	Receiver        RightsOmniUpdateRequestReceiver     `json:"receiver,omitempty"`
	AddressInfo     RightsOmniUpdateRequestAddressInfo3 `json:"addressInfo,omitempty"`
	PickUpVid       int64                               `json:"pickUpVid,omitempty"`
	PickUpName      string                              `json:"pickUpName,omitempty"`
	ReceiverAddress string                              `json:"receiverAddress,omitempty"`
}

type RightsOmniUpdateRequestReceiver struct {
	ReceiverMobile string `json:"receiverMobile,omitempty"`
	ReceiverName   string `json:"receiverName,omitempty"`
}

type RightsOmniUpdateRequestAddressInfo3 struct {
	AddressExt RightsOmniUpdateRequestAddressExt3 `json:"addressExt,omitempty"`
	County     string                             `json:"county,omitempty"`
	Area       string                             `json:"area,omitempty"`
	Longitude  string                             `json:"longitude,omitempty"`
	Province   string                             `json:"province,omitempty"`
	Latitude   string                             `json:"latitude,omitempty"`
	Address    string                             `json:"address,omitempty"`
	Zip        string                             `json:"zip,omitempty"`
	City       string                             `json:"city,omitempty"`
}

type RightsOmniUpdateRequestAddressExt3 struct {
	ProvinceCode string `json:"provinceCode,omitempty"`
	CityCode     string `json:"cityCode,omitempty"`
	CountyCode   string `json:"countyCode,omitempty"`
	AreaCode     string `json:"areaCode,omitempty"`
}

type RightsOmniUpdateRequestDeliveryInfo struct {
	ExpectReceivedEndTime   int64  `json:"expectReceivedEndTime,omitempty"`
	WriteOffId              int64  `json:"writeOffId,omitempty"`
	CompanyName             string `json:"companyName,omitempty"`
	ExpectReceivedTypeName  string `json:"expectReceivedTypeName,omitempty"`
	ExpectReceivedType      int64  `json:"expectReceivedType,omitempty"`
	WriteOffName            string `json:"writeOffName,omitempty"`
	ExpectReceivedStartTime int64  `json:"expectReceivedStartTime,omitempty"`
	Number                  string `json:"number,omitempty"`
	CompanyCode             string `json:"companyCode,omitempty"`
	ExpectReceivedDate      int64  `json:"expectReceivedDate,omitempty"`
}

type RightsOmniUpdateRequestReturnDeliveryInfo struct {
	DeliveryInfo   RightsOmniUpdateRequestDeliveryInfo2 `json:"deliveryInfo,omitempty"`
	DeliveryMethod int64                                `json:"deliveryMethod,omitempty"`
	OutPackageId   string                               `json:"outPackageId,omitempty"`
	DeliveryType   int64                                `json:"deliveryType,omitempty"`
	DeliveryTime   int64                                `json:"deliveryTime,omitempty"`
}

type RightsOmniUpdateRequestDeliveryInfo2 struct {
	CompanyCode string `json:"companyCode,omitempty"`
	Number      string `json:"number,omitempty"`
	CompanyName string `json:"companyName,omitempty"`
}

type RightsOmniUpdateRequestCancelInfo struct {
	CancelType int64 `json:"cancelType,omitempty"`
}

type RightsOmniUpdateRequestRefuseInfo struct {
	RefusedRemark string `json:"refusedRemark,omitempty"`
	RefusedReason string `json:"refusedReason,omitempty"`
}

type RightsOmniUpdateRequestRightsReasonInfo struct {
	RightsReason    string  `json:"rightsReason,omitempty"`
	ReasonImageUrls []int64 `json:"reasonImageUrls,omitempty"`
	RightsRemark    string  `json:"rightsRemark,omitempty"`
}
