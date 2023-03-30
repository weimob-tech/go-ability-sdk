package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderLogisticNewCreateOrder
// @id 1867
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=ISV向字节电子面单系统获取单号和打印信息)
func (client *Client) OrderLogisticNewCreateOrderV1_3(request *OrderLogisticNewCreateOrderRequestV1_3) (response *OrderLogisticNewCreateOrderResponseV1_3, err error) {
	rpcResponse := CreateOrderLogisticNewCreateOrderResponseV1_3()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderLogisticNewCreateOrderRequestV1_3 struct {
	*api.BaseRequest

	SenderInfo    OrderLogisticNewCreateOrderRequestV1_3SenderInfo   `json:"senderInfo,omitempty"`
	OrderInfos    []OrderLogisticNewCreateOrderRequestV1_3OrderInfos `json:"orderInfos,omitempty"`
	DeliveryReq   OrderLogisticNewCreateOrderRequestV1_3DeliveryReq  `json:"deliveryReq,omitempty"`
	LogisticsCode string                                             `json:"logisticsCode,omitempty"`
	UserId        int64                                              `json:"userId,omitempty"`
	OrderChannel  string                                             `json:"orderChannel,omitempty"`
}

type OrderLogisticNewCreateOrderResponseV1_3 struct {
}

func CreateOrderLogisticNewCreateOrderRequestV1_3() (request *OrderLogisticNewCreateOrderRequestV1_3) {
	request = &OrderLogisticNewCreateOrderRequestV1_3{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_3", "order/logisticNewCreateOrder", "api")
	request.Method = api.POST
	return
}

func CreateOrderLogisticNewCreateOrderResponseV1_3() (response *api.BaseResponse[OrderLogisticNewCreateOrderResponseV1_3]) {
	response = api.CreateResponse[OrderLogisticNewCreateOrderResponseV1_3](&OrderLogisticNewCreateOrderResponseV1_3{})
	return
}

type OrderLogisticNewCreateOrderRequestV1_3SenderInfo struct {
	Address OrderLogisticNewCreateOrderRequestV1_3Address `json:"address,omitempty"`
	Contact OrderLogisticNewCreateOrderRequestV1_3Contact `json:"contact,omitempty"`
}

type OrderLogisticNewCreateOrderRequestV1_3Address struct {
	CountryCode   string `json:"countryCode,omitempty"`
	ProvinceName  string `json:"provinceName,omitempty"`
	CityName      string `json:"cityName,omitempty"`
	DistrictName  string `json:"districtName,omitempty"`
	StreetName    string `json:"streetName,omitempty"`
	DetailAddress string `json:"detailAddress,omitempty"`
}

type OrderLogisticNewCreateOrderRequestV1_3Contact struct {
	Name   string `json:"name,omitempty"`
	Phone  string `json:"phone,omitempty"`
	Mobile string `json:"mobile,omitempty"`
}

type OrderLogisticNewCreateOrderRequestV1_3OrderInfos struct {
	ServiceList     []OrderLogisticNewCreateOrderRequestV1_3ServiceList   `json:"serviceList,omitempty"`
	PodModelAddress OrderLogisticNewCreateOrderRequestV1_3PodModelAddress `json:"podModelAddress,omitempty"`
	ReceiverInfo    OrderLogisticNewCreateOrderRequestV1_3ReceiverInfo    `json:"receiverInfo,omitempty"`
	Items           []OrderLogisticNewCreateOrderRequestV1_3Items         `json:"items,omitempty"`
	OrderNo         int64                                                 `json:"orderNo,omitempty"`
	PackId          string                                                `json:"packId,omitempty"`
	ProductType     string                                                `json:"productType,omitempty"`
	PayMethod       int64                                                 `json:"payMethod,omitempty"`
	PayAmount       int64                                                 `json:"payAmount,omitempty"`
}

type OrderLogisticNewCreateOrderRequestV1_3ServiceList struct {
	ServiceCode  string `json:"serviceCode,omitempty"`
	ServiceValue string `json:"serviceValue,omitempty"`
}

type OrderLogisticNewCreateOrderRequestV1_3PodModelAddress struct {
	CountryCode   string `json:"countryCode,omitempty"`
	ProvinceName  string `json:"provinceName,omitempty"`
	CityName      string `json:"cityName,omitempty"`
	DistrictName  string `json:"districtName,omitempty"`
	StreetName    string `json:"streetName,omitempty"`
	DetailAddress string `json:"detailAddress,omitempty"`
}

type OrderLogisticNewCreateOrderRequestV1_3ReceiverInfo struct {
	Address OrderLogisticNewCreateOrderRequestV1_3Address2 `json:"address,omitempty"`
	Contact OrderLogisticNewCreateOrderRequestV1_3Contact2 `json:"contact,omitempty"`
}

type OrderLogisticNewCreateOrderRequestV1_3Address2 struct {
	CountryCode   string `json:"countryCode,omitempty"`
	ProvinceName  string `json:"provinceName,omitempty"`
	CityName      string `json:"cityName,omitempty"`
	DistrictName  string `json:"districtName,omitempty"`
	StreetName    string `json:"streetName,omitempty"`
	DetailAddress string `json:"detailAddress,omitempty"`
}

type OrderLogisticNewCreateOrderRequestV1_3Contact2 struct {
	Name   string `json:"name,omitempty"`
	Phone  string `json:"phone,omitempty"`
	Mobile string `json:"mobile,omitempty"`
}

type OrderLogisticNewCreateOrderRequestV1_3Items struct {
	ItemName        string `json:"itemName,omitempty"`
	ItemSpecs       string `json:"itemSpecs,omitempty"`
	ItemCount       int64  `json:"itemCount,omitempty"`
	ItemVolume      int64  `json:"itemVolume,omitempty"`
	ItemWeight      int64  `json:"itemWeight,omitempty"`
	ItemNetWeight   int64  `json:"itemNetWeight,omitempty"`
	SenderFetchTime string `json:"senderFetchTime,omitempty"`
	IsSignVack      int64  `json:"isSignVack,omitempty"`
	Remark          string `json:"remark,omitempty"`
	Extra           string `json:"extra,omitempty"`
	TotalPackCount  int64  `json:"totalPackCount,omitempty"`
	TotalWeight     string `json:"totalWeight,omitempty"`
	TotalLength     string `json:"totalLength,omitempty"`
	TotalWidth      string `json:"totalWidth,omitempty"`
	TotalHeight     string `json:"totalHeight,omitempty"`
	Volume          string `json:"volume,omitempty"`
}

type OrderLogisticNewCreateOrderRequestV1_3DeliveryReq struct {
	IsCenterDelivery bool `json:"isCenterDelivery,omitempty"`
	IsPickupSelf     bool `json:"isPickupSelf,omitempty"`
}
