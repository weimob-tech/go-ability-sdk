package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsOmniImport
// @id 1931
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1931?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=全渠道导入售后单)
func (client *Client) RightsOmniImport(request *RightsOmniImportRequest) (response *RightsOmniImportResponse, err error) {
	rpcResponse := CreateRightsOmniImportResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsOmniImportRequest struct {
	*api.BaseRequest

	CancelInfo           RightsOmniImportRequestCancelInfo           `json:"cancelInfo,omitempty"`
	ExchangeDeliveryInfo RightsOmniImportRequestExchangeDeliveryInfo `json:"exchangeDeliveryInfo,omitempty"`
	MerchantInfo         RightsOmniImportRequestMerchantInfo         `json:"merchantInfo,omitempty"`
	Operator             RightsOmniImportRequestOperator             `json:"operator,omitempty"`
	OriginOrder          RightsOmniImportRequestOriginOrder          `json:"originOrder,omitempty"`
	RefundAccount        RightsOmniImportRequestRefundAccount        `json:"refundAccount,omitempty"`
	RefundDetail         RightsOmniImportRequestRefundDetail         `json:"refundDetail,omitempty"`
	RefundPayInfoList    []RightsOmniImportRequestRefundPayInfoList  `json:"refundPayInfoList,omitempty"`
	RefuseInfo           RightsOmniImportRequestRefuseInfo           `json:"refuseInfo,omitempty"`
	ReturnAddressInfo    RightsOmniImportRequestReturnAddressInfo    `json:"returnAddressInfo,omitempty"`
	ReturnDeliveryInfo   RightsOmniImportRequestReturnDeliveryInfo   `json:"returnDeliveryInfo,omitempty"`
	RightsAssets         []RightsOmniImportRequestRightsAssets       `json:"rightsAssets,omitempty"`
	RightsItems          []RightsOmniImportRequestRightsItems        `json:"rightsItems,omitempty"`
	RightsOrder          RightsOmniImportRequestRightsOrder          `json:"rightsOrder,omitempty"`
	RightsReasonInfo     RightsOmniImportRequestRightsReasonInfo     `json:"rightsReasonInfo,omitempty"`
}

type RightsOmniImportResponse struct {
	Success                bool  `json:"success,omitempty"`
	RightsOrderNo          int64 `json:"rightsOrderNo,omitempty"`
	FulfillOrderNo         int64 `json:"fulfillOrderNo,omitempty"`
	ExchangeOrderNo        int64 `json:"exchangeOrderNo,omitempty"`
	ExchangeFulfillOrderNo int64 `json:"exchangeFulfillOrderNo,omitempty"`
}

func CreateRightsOmniImportRequest() (request *RightsOmniImportRequest) {
	request = &RightsOmniImportRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "rights/omni/import", "apigw")
	request.Method = api.POST
	return
}

func CreateRightsOmniImportResponse() (response *api.BaseResponse[RightsOmniImportResponse]) {
	response = api.CreateResponse[RightsOmniImportResponse](&RightsOmniImportResponse{})
	return
}

type RightsOmniImportRequestCancelInfo struct {
	CancelType int64 `json:"cancelType,omitempty"`
}

type RightsOmniImportRequestExchangeDeliveryInfo struct {
	PackageInfoList         []RightsOmniImportRequestPackageInfoList `json:"packageInfoList,omitempty"`
	ReceiveInfo             RightsOmniImportRequestReceiveInfo2      `json:"receiveInfo,omitempty"`
	SenderInfo              RightsOmniImportRequestSenderInfo2       `json:"senderInfo,omitempty"`
	DeliveryType            int64                                    `json:"deliveryType,omitempty"`
	ExpectReceivedDate      int64                                    `json:"expectReceivedDate,omitempty"`
	ExpectReceivedEndTime   int64                                    `json:"expectReceivedEndTime,omitempty"`
	ExpectReceivedStartTime int64                                    `json:"expectReceivedStartTime,omitempty"`
	ExpectReceivedType      int64                                    `json:"expectReceivedType,omitempty"`
}

type RightsOmniImportRequestPackageInfoList struct {
	DeliveryInfo   RightsOmniImportRequestDeliveryInfo   `json:"deliveryInfo,omitempty"`
	PackageItems   []RightsOmniImportRequestPackageItems `json:"packageItems,omitempty"`
	ReceiveInfo    RightsOmniImportRequestReceiveInfo    `json:"receiveInfo,omitempty"`
	SenderInfo     RightsOmniImportRequestSenderInfo     `json:"senderInfo,omitempty"`
	BuyerRemark    string                                `json:"buyerRemark,omitempty"`
	DeliveryMethod int64                                 `json:"deliveryMethod,omitempty"`
	DeliveryTime   int64                                 `json:"deliveryTime,omitempty"`
	OutPackageId   string                                `json:"outPackageId,omitempty"`
	ReceiveTime    int64                                 `json:"receiveTime,omitempty"`
	Remark         string                                `json:"remark,omitempty"`
}

type RightsOmniImportRequestDeliveryInfo struct {
	CompanyCode             string `json:"companyCode,omitempty"`
	CompanyName             string `json:"companyName,omitempty"`
	ExpectReceivedDate      int64  `json:"expectReceivedDate,omitempty"`
	ExpectReceivedEndTime   int64  `json:"expectReceivedEndTime,omitempty"`
	ExpectReceivedStartTime int64  `json:"expectReceivedStartTime,omitempty"`
	ExpectReceivedType      int64  `json:"expectReceivedType,omitempty"`
	ExpectReceivedTypeName  string `json:"expectReceivedTypeName,omitempty"`
	Number                  string `json:"number,omitempty"`
	WriteOffId              int64  `json:"writeOffId,omitempty"`
	WriteOffName            string `json:"writeOffName,omitempty"`
}

type RightsOmniImportRequestPackageItems struct {
	Num             string `json:"num,omitempty"`
	OutRightsItemId string `json:"outRightsItemId,omitempty"`
}

type RightsOmniImportRequestReceiveInfo struct {
	AddressInfo     RightsOmniImportRequestAddressInfo `json:"addressInfo,omitempty"`
	Receiver        RightsOmniImportRequestReceiver    `json:"receiver,omitempty"`
	PickUpName      string                             `json:"pickUpName,omitempty"`
	PickUpVid       int64                              `json:"pickUpVid,omitempty"`
	ReceiverAddress string                             `json:"receiverAddress,omitempty"`
}

type RightsOmniImportRequestAddressInfo struct {
	AddressExt RightsOmniImportRequestAddressExt `json:"addressExt,omitempty"`
	Address    string                            `json:"address,omitempty"`
	Area       string                            `json:"area,omitempty"`
	City       string                            `json:"city,omitempty"`
	County     string                            `json:"county,omitempty"`
	Latitude   string                            `json:"latitude,omitempty"`
	Longitude  string                            `json:"longitude,omitempty"`
	Province   string                            `json:"province,omitempty"`
	Zip        string                            `json:"zip,omitempty"`
}

type RightsOmniImportRequestAddressExt struct {
	AreaCode     string `json:"areaCode,omitempty"`
	CityCode     string `json:"cityCode,omitempty"`
	CountyCode   string `json:"countyCode,omitempty"`
	ProvinceCode string `json:"provinceCode,omitempty"`
}

type RightsOmniImportRequestReceiver struct {
	ReceiverMobile string `json:"receiverMobile,omitempty"`
	ReceiverName   string `json:"receiverName,omitempty"`
}

type RightsOmniImportRequestSenderInfo struct {
	AddressInfo    RightsOmniImportRequestAddressInfo2 `json:"addressInfo,omitempty"`
	Sender         RightsOmniImportRequestSender       `json:"sender,omitempty"`
	ProcessVid     int64                               `json:"processVid,omitempty"`
	ProcessVidName string                              `json:"processVidName,omitempty"`
}

type RightsOmniImportRequestAddressInfo2 struct {
	AddressExt RightsOmniImportRequestAddressExt2 `json:"addressExt,omitempty"`
	Address    string                             `json:"address,omitempty"`
	Area       string                             `json:"area,omitempty"`
	City       string                             `json:"city,omitempty"`
	County     string                             `json:"county,omitempty"`
	Latitude   string                             `json:"latitude,omitempty"`
	Longitude  string                             `json:"longitude,omitempty"`
	Province   string                             `json:"province,omitempty"`
	Zip        string                             `json:"zip,omitempty"`
}

type RightsOmniImportRequestAddressExt2 struct {
	AreaCode     string `json:"areaCode,omitempty"`
	CityCode     string `json:"cityCode,omitempty"`
	CountyCode   string `json:"countyCode,omitempty"`
	ProvinceCode string `json:"provinceCode,omitempty"`
}

type RightsOmniImportRequestSender struct {
	SenderMobile string `json:"senderMobile,omitempty"`
	SenderName   string `json:"senderName,omitempty"`
}

type RightsOmniImportRequestReceiveInfo2 struct {
	AddressInfo     RightsOmniImportRequestAddressInfo3 `json:"addressInfo,omitempty"`
	Receiver        RightsOmniImportRequestReceiver2    `json:"receiver,omitempty"`
	PickUpName      string                              `json:"pickUpName,omitempty"`
	PickUpVid       int64                               `json:"pickUpVid,omitempty"`
	ReceiverAddress string                              `json:"receiverAddress,omitempty"`
}

type RightsOmniImportRequestAddressInfo3 struct {
	AddressExt RightsOmniImportRequestAddressExt3 `json:"addressExt,omitempty"`
	Address    string                             `json:"address,omitempty"`
	Area       string                             `json:"area,omitempty"`
	City       string                             `json:"city,omitempty"`
	County     string                             `json:"county,omitempty"`
	Latitude   string                             `json:"latitude,omitempty"`
	Longitude  string                             `json:"longitude,omitempty"`
	Province   string                             `json:"province,omitempty"`
	Zip        string                             `json:"zip,omitempty"`
}

type RightsOmniImportRequestAddressExt3 struct {
	AreaCode     string `json:"areaCode,omitempty"`
	CityCode     string `json:"cityCode,omitempty"`
	CountyCode   string `json:"countyCode,omitempty"`
	ProvinceCode string `json:"provinceCode,omitempty"`
}

type RightsOmniImportRequestReceiver2 struct {
	ReceiverMobile string `json:"receiverMobile,omitempty"`
	ReceiverName   string `json:"receiverName,omitempty"`
}

type RightsOmniImportRequestSenderInfo2 struct {
	AddressInfo    RightsOmniImportRequestAddressInfo4 `json:"addressInfo,omitempty"`
	Sender         RightsOmniImportRequestSender2      `json:"sender,omitempty"`
	ProcessVid     int64                               `json:"processVid,omitempty"`
	ProcessVidName string                              `json:"processVidName,omitempty"`
}

type RightsOmniImportRequestAddressInfo4 struct {
	AddressExt RightsOmniImportRequestAddressExt4 `json:"addressExt,omitempty"`
	Address    string                             `json:"address,omitempty"`
	Area       string                             `json:"area,omitempty"`
	City       string                             `json:"city,omitempty"`
	County     string                             `json:"county,omitempty"`
	Latitude   string                             `json:"latitude,omitempty"`
	Longitude  string                             `json:"longitude,omitempty"`
	Province   string                             `json:"province,omitempty"`
	Zip        string                             `json:"zip,omitempty"`
}

type RightsOmniImportRequestAddressExt4 struct {
	AreaCode     string `json:"areaCode,omitempty"`
	CityCode     string `json:"cityCode,omitempty"`
	CountyCode   string `json:"countyCode,omitempty"`
	ProvinceCode string `json:"provinceCode,omitempty"`
}

type RightsOmniImportRequestSender2 struct {
	SenderMobile string `json:"senderMobile,omitempty"`
	SenderName   string `json:"senderName,omitempty"`
}

type RightsOmniImportRequestMerchantInfo struct {
	BosName         string `json:"bosName,omitempty"`
	OuterMerchantId string `json:"outerMerchantId,omitempty"`
	ProcessVid      int64  `json:"processVid,omitempty"`
	ProcessVidName  string `json:"processVidName,omitempty"`
	ProcessVidType  int64  `json:"processVidType,omitempty"`
	Vid             int64  `json:"vid,omitempty"`
	VidName         string `json:"vidName,omitempty"`
}

type RightsOmniImportRequestOperator struct {
	OperatorId    string `json:"operatorId,omitempty"`
	OperatorName  string `json:"operatorName,omitempty"`
	OperatorPhone string `json:"operatorPhone,omitempty"`
}

type RightsOmniImportRequestOriginOrder struct {
	OrderNo    int64  `json:"orderNo,omitempty"`
	OutOrderNo string `json:"outOrderNo,omitempty"`
}

type RightsOmniImportRequestRefundAccount struct {
	AccountNum     string `json:"accountNum,omitempty"`
	OpeningBank    string `json:"openingBank,omitempty"`
	RefundMethodId int64  `json:"refundMethodId,omitempty"`
	UserName       string `json:"userName,omitempty"`
}

type RightsOmniImportRequestRefundDetail struct {
	ApplyAmount  string `json:"applyAmount,omitempty"`
	RefundAmount string `json:"refundAmount,omitempty"`
}

type RightsOmniImportRequestRefundPayInfoList struct {
	Amount             string `json:"amount,omitempty"`
	RefundFailedReason string `json:"refundFailedReason,omitempty"`
	RefundMethodId     int64  `json:"refundMethodId,omitempty"`
	RefundStatus       int64  `json:"refundStatus,omitempty"`
}

type RightsOmniImportRequestRefuseInfo struct {
	RefusedReason string `json:"refusedReason,omitempty"`
	RefusedRemark string `json:"refusedRemark,omitempty"`
}

type RightsOmniImportRequestReturnAddressInfo struct {
	AddressInfo RightsOmniImportRequestAddressInfo5 `json:"addressInfo,omitempty"`
	Name        string                              `json:"name,omitempty"`
	Phone       string                              `json:"phone,omitempty"`
}

type RightsOmniImportRequestAddressInfo5 struct {
	AddressExt RightsOmniImportRequestAddressExt5 `json:"addressExt,omitempty"`
	Address    string                             `json:"address,omitempty"`
	Area       string                             `json:"area,omitempty"`
	City       string                             `json:"city,omitempty"`
	County     string                             `json:"county,omitempty"`
	Latitude   string                             `json:"latitude,omitempty"`
	Longitude  string                             `json:"longitude,omitempty"`
	Province   string                             `json:"province,omitempty"`
	Zip        string                             `json:"zip,omitempty"`
}

type RightsOmniImportRequestAddressExt5 struct {
	AreaCode     string `json:"areaCode,omitempty"`
	CityCode     string `json:"cityCode,omitempty"`
	CountyCode   string `json:"countyCode,omitempty"`
	ProvinceCode string `json:"provinceCode,omitempty"`
}

type RightsOmniImportRequestReturnDeliveryInfo struct {
	DeliveryInfo   RightsOmniImportRequestDeliveryInfo2 `json:"deliveryInfo,omitempty"`
	DeliveryMethod int64                                `json:"deliveryMethod,omitempty"`
	DeliveryType   int64                                `json:"deliveryType,omitempty"`
	OutPackageId   string                               `json:"outPackageId,omitempty"`
}

type RightsOmniImportRequestDeliveryInfo2 struct {
	CompanyCode string `json:"companyCode,omitempty"`
	CompanyName string `json:"companyName,omitempty"`
	Number      string `json:"number,omitempty"`
}

type RightsOmniImportRequestRightsAssets struct {
	AssetsAmount string `json:"assetsAmount,omitempty"`
	AssetsNum    string `json:"assetsNum,omitempty"`
	AssetsTarget int64  `json:"assetsTarget,omitempty"`
	AssetsType   int64  `json:"assetsType,omitempty"`
}

type RightsOmniImportRequestRightsItems struct {
	ExchangeGoods       RightsOmniImportRequestExchangeGoods       `json:"exchangeGoods,omitempty"`
	RefundDetail        RightsOmniImportRequestRefundDetail2       `json:"refundDetail,omitempty"`
	RightsAssets        []RightsOmniImportRequestRightsAssets2     `json:"rightsAssets,omitempty"`
	RightsItemExtraInfo RightsOmniImportRequestRightsItemExtraInfo `json:"rightsItemExtraInfo,omitempty"`
	ApplyNum            string                                     `json:"applyNum,omitempty"`
	GoodsReceiveType    int64                                      `json:"goodsReceiveType,omitempty"`
	OrderItemId         int64                                      `json:"orderItemId,omitempty"`
	OutOrderItemId      string                                     `json:"outOrderItemId,omitempty"`
	OutRightsItemId     string                                     `json:"outRightsItemId,omitempty"`
}

type RightsOmniImportRequestExchangeGoods struct {
	GoodsCode    string  `json:"goodsCode,omitempty"`
	GoodsId      int64   `json:"goodsId,omitempty"`
	GoodsTitle   string  `json:"goodsTitle,omitempty"`
	GoodsType    int64   `json:"goodsType,omitempty"`
	ImageUrl     []int64 `json:"imageUrl,omitempty"`
	Price        string  `json:"price,omitempty"`
	SkuCode      string  `json:"skuCode,omitempty"`
	SkuId        int64   `json:"skuId,omitempty"`
	SkuName      string  `json:"skuName,omitempty"`
	SkuNum       string  `json:"skuNum,omitempty"`
	SubGoodsType int64   `json:"subGoodsType,omitempty"`
}

type RightsOmniImportRequestRefundDetail2 struct {
	ApplyAmount  string `json:"applyAmount,omitempty"`
	RefundAmount string `json:"refundAmount,omitempty"`
}

type RightsOmniImportRequestRightsAssets2 struct {
	AssetsAmount string `json:"assetsAmount,omitempty"`
	AssetsNum    string `json:"assetsNum,omitempty"`
	AssetsTarget int64  `json:"assetsTarget,omitempty"`
	AssetsType   int64  `json:"assetsType,omitempty"`
}

type RightsOmniImportRequestRightsItemExtraInfo struct {
	ReturnGoodsNum int64 `json:"returnGoodsNum,omitempty"`
}

type RightsOmniImportRequestRightsOrder struct {
	Operations      []RightsOmniImportRequestOperations `json:"operations,omitempty"`
	ChannelType     int64                               `json:"channelType,omitempty"`
	OutCreateTime   int64                               `json:"outCreateTime,omitempty"`
	OutRightsId     string                              `json:"outRightsId,omitempty"`
	RefundType      int64                               `json:"refundType,omitempty"`
	RightsCauseType int64                               `json:"rightsCauseType,omitempty"`
	RightsSource    int64                               `json:"rightsSource,omitempty"`
	RightsStatus    int64                               `json:"rightsStatus,omitempty"`
	RightsType      int64                               `json:"rightsType,omitempty"`
}

type RightsOmniImportRequestOperations struct {
	Operator RightsOmniImportRequestOperator2 `json:"operator,omitempty"`
	Datetime int64                            `json:"datetime,omitempty"`
	Type     int64                            `json:"type,omitempty"`
}

type RightsOmniImportRequestOperator2 struct {
	OperatorId    string `json:"operatorId,omitempty"`
	OperatorName  string `json:"operatorName,omitempty"`
	OperatorPhone string `json:"operatorPhone,omitempty"`
}

type RightsOmniImportRequestRightsReasonInfo struct {
	ReasonImageUrls []int64 `json:"reasonImageUrls,omitempty"`
	RightsReason    string  `json:"rightsReason,omitempty"`
	RightsRemark    string  `json:"rightsRemark,omitempty"`
}
