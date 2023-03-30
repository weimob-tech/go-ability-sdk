package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpRemoveUserTag
// @id 1668
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量删除客户标签)
func (client *Client) MbpRemoveUserTagV2_0(request *MbpRemoveUserTagRequestV2_0) (response *MbpRemoveUserTagResponseV2_0, err error) {
	rpcResponse := CreateMbpRemoveUserTagResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpRemoveUserTagRequestV2_0 struct {
	*api.BaseRequest

	StoreId           int64   `json:"storeId,omitempty"`
	WidList           []int64 `json:"widList,omitempty"`
	RemoveUserTagList []int64 `json:"removeUserTagList,omitempty"`
}

type MbpRemoveUserTagResponseV2_0 struct {
}

func CreateMbpRemoveUserTagRequestV2_0() (request *MbpRemoveUserTagRequestV2_0) {
	request = &MbpRemoveUserTagRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "2_0", "mbp/removeUserTag", "api")
	request.Method = api.POST
	return
}

func CreateMbpRemoveUserTagResponseV2_0() (response *api.BaseResponse[MbpRemoveUserTagResponseV2_0]) {
	response = api.CreateResponse[MbpRemoveUserTagResponseV2_0](&MbpRemoveUserTagResponseV2_0{})
	return
}
