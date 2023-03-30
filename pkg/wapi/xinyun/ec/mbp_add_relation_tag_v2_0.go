package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpAddRelationTag
// @id 1665
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增标签关联关系)
func (client *Client) MbpAddRelationTagV2_0(request *MbpAddRelationTagRequestV2_0) (response *MbpAddRelationTagResponseV2_0, err error) {
	rpcResponse := CreateMbpAddRelationTagResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpAddRelationTagRequestV2_0 struct {
	*api.BaseRequest

	StoreId            int64   `json:"storeId,omitempty"`
	RelationType       int64   `json:"relationType,omitempty"`
	TagId              int64   `json:"tagId,omitempty"`
	RelationId         string  `json:"relationId,omitempty"`
	RelationAttrVoList []int64 `json:"relationAttrVoList,omitempty"`
	Pid                string  `json:"pid,omitempty"`
}

type MbpAddRelationTagResponseV2_0 struct {
}

func CreateMbpAddRelationTagRequestV2_0() (request *MbpAddRelationTagRequestV2_0) {
	request = &MbpAddRelationTagRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "2_0", "mbp/addRelationTag", "api")
	request.Method = api.POST
	return
}

func CreateMbpAddRelationTagResponseV2_0() (response *api.BaseResponse[MbpAddRelationTagResponseV2_0]) {
	response = api.CreateResponse[MbpAddRelationTagResponseV2_0](&MbpAddRelationTagResponseV2_0{})
	return
}
