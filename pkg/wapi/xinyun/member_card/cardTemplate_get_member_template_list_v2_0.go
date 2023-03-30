package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CardTemplateGetMemberTemplateList
// @id 1450
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询会员卡模板列表)
func (client *Client) CardTemplateGetMemberTemplateListV2_0(request *CardTemplateGetMemberTemplateListRequestV2_0) (response *CardTemplateGetMemberTemplateListResponseV2_0, err error) {
	rpcResponse := CreateCardTemplateGetMemberTemplateListResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CardTemplateGetMemberTemplateListRequestV2_0 struct {
	*api.BaseRequest

	Page     int `json:"page,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
}

type CardTemplateGetMemberTemplateListResponseV2_0 struct {
}

func CreateCardTemplateGetMemberTemplateListRequestV2_0() (request *CardTemplateGetMemberTemplateListRequestV2_0) {
	request = &CardTemplateGetMemberTemplateListRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "2_0", "cardTemplate/getMemberTemplateList", "api")
	request.Method = api.POST
	return
}

func CreateCardTemplateGetMemberTemplateListResponseV2_0() (response *api.BaseResponse[CardTemplateGetMemberTemplateListResponseV2_0]) {
	response = api.CreateResponse[CardTemplateGetMemberTemplateListResponseV2_0](&CardTemplateGetMemberTemplateListResponseV2_0{})
	return
}
