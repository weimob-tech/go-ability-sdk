package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerPublicseaObtain
// @id 1777
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=领取/分配客户公海数据)
func (client *Client) CustomerPublicseaObtain(request *CustomerPublicseaObtainRequest) (response *CustomerPublicseaObtainResponse, err error) {
	rpcResponse := CreateCustomerPublicseaObtainResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerPublicseaObtainRequest struct {
	*api.BaseRequest

	Owner      int64  `json:"owner,omitempty"`
	Keys       string `json:"keys,omitempty"`
	Wid        int64  `json:"wid,omitempty"`
	OptionType int    `json:"optionType,omitempty"`
}

type CustomerPublicseaObtainResponse struct {
}

func CreateCustomerPublicseaObtainRequest() (request *CustomerPublicseaObtainRequest) {
	request = &CustomerPublicseaObtainRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "customer/publicseaObtain", "api")
	request.Method = api.POST
	return
}

func CreateCustomerPublicseaObtainResponse() (response *api.BaseResponse[CustomerPublicseaObtainResponse]) {
	response = api.CreateResponse[CustomerPublicseaObtainResponse](&CustomerPublicseaObtainResponse{})
	return
}
