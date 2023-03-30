package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SettingGetPointUseRuleV2
// @id 1487
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询积分抵扣规则V2)
func (client *Client) SettingGetPointUseRuleV2V2_0(request *SettingGetPointUseRuleV2RequestV2_0) (response *SettingGetPointUseRuleV2ResponseV2_0, err error) {
	rpcResponse := CreateSettingGetPointUseRuleV2ResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SettingGetPointUseRuleV2RequestV2_0 struct {
	*api.BaseRequest

	Type int `json:"type,omitempty"`
}

type SettingGetPointUseRuleV2ResponseV2_0 struct {
}

func CreateSettingGetPointUseRuleV2RequestV2_0() (request *SettingGetPointUseRuleV2RequestV2_0) {
	request = &SettingGetPointUseRuleV2RequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "2_0", "setting/getPointUseRuleV2", "api")
	request.Method = api.POST
	return
}

func CreateSettingGetPointUseRuleV2ResponseV2_0() (response *api.BaseResponse[SettingGetPointUseRuleV2ResponseV2_0]) {
	response = api.CreateResponse[SettingGetPointUseRuleV2ResponseV2_0](&SettingGetPointUseRuleV2ResponseV2_0{})
	return
}
