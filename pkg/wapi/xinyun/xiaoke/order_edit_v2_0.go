package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderEdit
// @id 2680
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新商机)
func (client *Client) OrderEditV2_0(request *OrderEditRequestV2_0) (response *OrderEditResponseV2_0, err error) {
	rpcResponse := CreateOrderEditResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderEditRequestV2_0 struct {
	*api.BaseRequest

	List        []OrderEditRequestV2_0list     `json:"list,omitempty"`
	Products    []OrderEditRequestV2_0Products `json:"products,omitempty"`
	UserWid     int64                          `json:"userWid,omitempty"`
	OrderNumber string                         `json:"orderNumber,omitempty"`
}

type OrderEditResponseV2_0 struct {
}

func CreateOrderEditRequestV2_0() (request *OrderEditRequestV2_0) {
	request = &OrderEditRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "2_0", "order/edit", "api")
	request.Method = api.POST
	return
}

func CreateOrderEditResponseV2_0() (response *api.BaseResponse[OrderEditResponseV2_0]) {
	response = api.CreateResponse[OrderEditResponseV2_0](&OrderEditResponseV2_0{})
	return
}

type OrderEditRequestV2_0list struct {
	FieldType []map[string]any `json:"fieldType,omitempty"`
	FieldKey  int              `json:"fieldKey,omitempty"`
	Value     string           `json:"value,omitempty"`
}

type OrderEditRequestV2_0Products struct {
	ProductUniqueNo string  `json:"productUniqueNo,omitempty"`
	Discount        float64 `json:"discount,omitempty"`
	SellingPrice    float64 `json:"sellingPrice,omitempty"`
	Number          int     `json:"number,omitempty"`
	Remark          string  `json:"remark,omitempty"`
	Version         int     `json:"version,omitempty"`
}
