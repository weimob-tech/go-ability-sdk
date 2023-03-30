package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ProductFieldValidateRule
// @id 2674
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询产品字段验证规则)
func (client *Client) ProductFieldValidateRule(request *ProductFieldValidateRuleRequest) (response *ProductFieldValidateRuleResponse, err error) {
	rpcResponse := CreateProductFieldValidateRuleResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ProductFieldValidateRuleRequest struct {
	*api.BaseRequest

	UserWid int64 `json:"userWid,omitempty"`
}

type ProductFieldValidateRuleResponse struct {
}

func CreateProductFieldValidateRuleRequest() (request *ProductFieldValidateRuleRequest) {
	request = &ProductFieldValidateRuleRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "product/fieldValidateRule", "api")
	request.Method = api.POST
	return
}

func CreateProductFieldValidateRuleResponse() (response *api.BaseResponse[ProductFieldValidateRuleResponse]) {
	response = api.CreateResponse[ProductFieldValidateRuleResponse](&ProductFieldValidateRuleResponse{})
	return
}
