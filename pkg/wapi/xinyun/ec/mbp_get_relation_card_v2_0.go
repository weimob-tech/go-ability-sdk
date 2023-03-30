package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpGetRelationCard
// @id 2701
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=领取副卡)
func (client *Client) MbpGetRelationCardV2_0(request *MbpGetRelationCardRequestV2_0) (response *MbpGetRelationCardResponseV2_0, err error) {
	rpcResponse := CreateMbpGetRelationCardResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpGetRelationCardRequestV2_0 struct {
	*api.BaseRequest

	Pid            int64  `json:"pid,omitempty"`
	Wid            int64  `json:"wid,omitempty"`
	BelongWid      int64  `json:"belongWid,omitempty"`
	CardTemplateId int64  `json:"cardTemplateId,omitempty"`
	Phone          string `json:"phone,omitempty"`
	StoreId        int64  `json:"storeId,omitempty"`
	AppChannel     int    `json:"appChannel,omitempty"`
	GetChannelId   int64  `json:"getChannelId,omitempty"`
	GetChannel     int    `json:"getChannel,omitempty"`
	Name           string `json:"name,omitempty"`
}

type MbpGetRelationCardResponseV2_0 struct {
}

func CreateMbpGetRelationCardRequestV2_0() (request *MbpGetRelationCardRequestV2_0) {
	request = &MbpGetRelationCardRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "2_0", "mbp/getRelationCard", "api")
	request.Method = api.POST
	return
}

func CreateMbpGetRelationCardResponseV2_0() (response *api.BaseResponse[MbpGetRelationCardResponseV2_0]) {
	response = api.CreateResponse[MbpGetRelationCardResponseV2_0](&MbpGetRelationCardResponseV2_0{})
	return
}
