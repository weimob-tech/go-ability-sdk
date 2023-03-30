package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomOptionCreate
// @id 2132
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2132?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=指定父节点新增自定义字段选项信息)
func (client *Client) CustomOptionCreate(request *CustomOptionCreateRequest) (response *CustomOptionCreateResponse, err error) {
	rpcResponse := CreateCustomOptionCreateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomOptionCreateRequest struct {
	*api.BaseRequest

	BasicInfo        CustomOptionCreateRequestBasicInfo          `json:"basicInfo,omitempty"`
	CustomOptionList []CustomOptionCreateRequestCustomOptionList `json:"customOptionList,omitempty"`
	ParentId         int64                                       `json:"parentId,omitempty"`
	AttributeId      string                                      `json:"attributeId,omitempty"`
}

type CustomOptionCreateResponse struct {
	CustomOptionsResultInfoList []CustomOptionCreateResponseCustomOptionsResultInfoList `json:"customOptionsResultInfoList,omitempty"`
	IsSuccess                   bool                                                    `json:"isSuccess,omitempty"`
	RemainOptionNum             int64                                                   `json:"remainOptionNum,omitempty"`
}

func CreateCustomOptionCreateRequest() (request *CustomOptionCreateRequest) {
	request = &CustomOptionCreateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "custom/option/create", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomOptionCreateResponse() (response *api.BaseResponse[CustomOptionCreateResponse]) {
	response = api.CreateResponse[CustomOptionCreateResponse](&CustomOptionCreateResponse{})
	return
}

type CustomOptionCreateRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type CustomOptionCreateRequestCustomOptionList struct {
	OptionName string `json:"optionName,omitempty"`
}

type CustomOptionCreateResponseCustomOptionsResultInfoList struct {
	Id         int64  `json:"id,omitempty"`
	IsSuccess  bool   `json:"isSuccess,omitempty"`
	OptionName string `json:"optionName,omitempty"`
	ReturnMsg  string `json:"returnMsg,omitempty"`
}
