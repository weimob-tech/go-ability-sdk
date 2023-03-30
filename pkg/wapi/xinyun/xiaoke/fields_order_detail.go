package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FieldsOrderDetail
// @id 2660
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询订单字段校验规则)
func (client *Client) FieldsOrderDetail(request *FieldsOrderDetailRequest) (response *FieldsOrderDetailResponse, err error) {
	rpcResponse := CreateFieldsOrderDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FieldsOrderDetailRequest struct {
	*api.BaseRequest

	UserWid int64  `json:"userWid,omitempty"`
	Stage   string `json:"stage,omitempty"`
}

type FieldsOrderDetailResponse struct {
}

func CreateFieldsOrderDetailRequest() (request *FieldsOrderDetailRequest) {
	request = &FieldsOrderDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "fields/orderDetail", "api")
	request.Method = api.POST
	return
}

func CreateFieldsOrderDetailResponse() (response *api.BaseResponse[FieldsOrderDetailResponse]) {
	response = api.CreateResponse[FieldsOrderDetailResponse](&FieldsOrderDetailResponse{})
	return
}
