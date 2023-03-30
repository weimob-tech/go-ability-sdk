package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomAttributeUpdate
// @id 2136
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2136?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=修改自定义字段)
func (client *Client) CustomAttributeUpdate(request *CustomAttributeUpdateRequest) (response *CustomAttributeUpdateResponse, err error) {
	rpcResponse := CreateCustomAttributeUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomAttributeUpdateRequest struct {
	*api.BaseRequest

	BasicInfo           CustomAttributeUpdateRequestBasicInfo           `json:"basicInfo,omitempty"`
	DimensionInfoVoList CustomAttributeUpdateRequestDimensionInfoVoList `json:"dimensionInfoVoList,omitempty"`
	Id                  string                                          `json:"id,omitempty"`
	AttributeName       string                                          `json:"attributeName,omitempty"`
	IsMust              int64                                           `json:"isMust,omitempty"`
	Sort                int64                                           `json:"sort,omitempty"`
	Status              int64                                           `json:"status,omitempty"`
	DefaultValue        string                                          `json:"defaultValue,omitempty"`
	BizId               int64                                           `json:"bizId,omitempty"`
	BizType             int64                                           `json:"bizType,omitempty"`
	FormType            int64                                           `json:"formType,omitempty"`
}

type CustomAttributeUpdateResponse struct {
}

func CreateCustomAttributeUpdateRequest() (request *CustomAttributeUpdateRequest) {
	request = &CustomAttributeUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "custom/attribute/update", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomAttributeUpdateResponse() (response *api.BaseResponse[CustomAttributeUpdateResponse]) {
	response = api.CreateResponse[CustomAttributeUpdateResponse](&CustomAttributeUpdateResponse{})
	return
}

type CustomAttributeUpdateRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type CustomAttributeUpdateRequestDimensionInfoVoList struct {
	DimensionType      int64   `json:"dimensionType,omitempty"`
	DimensionValueList []int64 `json:"dimensionValueList,omitempty"`
}
