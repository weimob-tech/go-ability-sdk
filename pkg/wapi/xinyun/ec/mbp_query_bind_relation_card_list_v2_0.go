package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpQueryBindRelationCardList
// @id 2699
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询主卡下副卡列表)
func (client *Client) MbpQueryBindRelationCardListV2_0(request *MbpQueryBindRelationCardListRequestV2_0) (response *MbpQueryBindRelationCardListResponseV2_0, err error) {
	rpcResponse := CreateMbpQueryBindRelationCardListResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpQueryBindRelationCardListRequestV2_0 struct {
	*api.BaseRequest

	Pid            int64 `json:"pid,omitempty"`
	Wid            int64 `json:"wid,omitempty"`
	CardTemplateId int64 `json:"cardTemplateId,omitempty"`
}

type MbpQueryBindRelationCardListResponseV2_0 struct {
}

func CreateMbpQueryBindRelationCardListRequestV2_0() (request *MbpQueryBindRelationCardListRequestV2_0) {
	request = &MbpQueryBindRelationCardListRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "2_0", "mbp/queryBindRelationCardList", "api")
	request.Method = api.POST
	return
}

func CreateMbpQueryBindRelationCardListResponseV2_0() (response *api.BaseResponse[MbpQueryBindRelationCardListResponseV2_0]) {
	response = api.CreateResponse[MbpQueryBindRelationCardListResponseV2_0](&MbpQueryBindRelationCardListResponseV2_0{})
	return
}
