package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomAttributeCreate
// @id 2137
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2137?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=新增自定义字段)
func (client *Client) CustomAttributeCreate(request *CustomAttributeCreateRequest) (response *CustomAttributeCreateResponse, err error) {
	rpcResponse := CreateCustomAttributeCreateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomAttributeCreateRequest struct {
	*api.BaseRequest

	BasicInfo           CustomAttributeCreateRequestBasicInfo             `json:"basicInfo,omitempty"`
	TitleNameList       []CustomAttributeCreateRequestTitleNameList       `json:"titleNameList,omitempty"`
	DimensionInfoVoList []CustomAttributeCreateRequestDimensionInfoVoList `json:"dimensionInfoVoList,omitempty"`
	AttributeName       string                                            `json:"attributeName,omitempty"`
	FormType            int64                                             `json:"formType,omitempty"`
	IsMust              int64                                             `json:"isMust,omitempty"`
	Sort                string                                            `json:"sort,omitempty"`
	Level               int64                                             `json:"level,omitempty"`
	Status              int64                                             `json:"status,omitempty"`
	BizType             string                                            `json:"bizType,omitempty"`
	BizId               int64                                             `json:"bizId,omitempty"`
	DefaultValue        string                                            `json:"defaultValue,omitempty"`
}

type CustomAttributeCreateResponse struct {
	TitleInfoList []CustomAttributeCreateResponseTitleInfoList `json:"titleInfoList,omitempty"`
	AttributeName string                                       `json:"attributeName,omitempty"`
	BizId         int64                                        `json:"bizId,omitempty"`
	BizType       int64                                        `json:"bizType,omitempty"`
	BosId         int64                                        `json:"bosId,omitempty"`
	FormType      int64                                        `json:"formType,omitempty"`
	Id            int64                                        `json:"id,omitempty"`
	IsMust        int64                                        `json:"isMust,omitempty"`
	Level         int64                                        `json:"level,omitempty"`
	MerchantId    int64                                        `json:"merchantId,omitempty"`
	Sort          int64                                        `json:"sort,omitempty"`
	Status        int64                                        `json:"status,omitempty"`
	Vid           int64                                        `json:"vid,omitempty"`
	DefaultValue  string                                       `json:"defaultValue,omitempty"`
}

func CreateCustomAttributeCreateRequest() (request *CustomAttributeCreateRequest) {
	request = &CustomAttributeCreateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "custom/attribute/create", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomAttributeCreateResponse() (response *api.BaseResponse[CustomAttributeCreateResponse]) {
	response = api.CreateResponse[CustomAttributeCreateResponse](&CustomAttributeCreateResponse{})
	return
}

type CustomAttributeCreateRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type CustomAttributeCreateRequestTitleNameList struct {
	TitleName string `json:"titleName,omitempty"`
	Level     int64  `json:"level,omitempty"`
}

type CustomAttributeCreateRequestDimensionInfoVoList struct {
	DimensionType      int64   `json:"dimensionType,omitempty"`
	DimensionValueList []int64 `json:"dimensionValueList,omitempty"`
}

type CustomAttributeCreateResponseTitleInfoList struct {
	AttributeId int64  `json:"attributeId,omitempty"`
	Id          int64  `json:"id,omitempty"`
	IsTitle     int64  `json:"isTitle,omitempty"`
	Level       int64  `json:"level,omitempty"`
	OptionName  string `json:"optionName,omitempty"`
	ParentId    int64  `json:"parentId,omitempty"`
	SortNum     int64  `json:"sortNum,omitempty"`
}
