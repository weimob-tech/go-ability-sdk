package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpBatchRollWidByTagId
// @id 1669
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=滚动查询标签的客户)
func (client *Client) MbpBatchRollWidByTagIdV2_0(request *MbpBatchRollWidByTagIdRequestV2_0) (response *MbpBatchRollWidByTagIdResponseV2_0, err error) {
	rpcResponse := CreateMbpBatchRollWidByTagIdResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpBatchRollWidByTagIdRequestV2_0 struct {
	*api.BaseRequest

	StoreId  int64 `json:"storeId,omitempty"`
	TagId    int64 `json:"tagId,omitempty"`
	Cursor   int64 `json:"cursor,omitempty"`
	PageSize int   `json:"pageSize,omitempty"`
}

type MbpBatchRollWidByTagIdResponseV2_0 struct {
}

func CreateMbpBatchRollWidByTagIdRequestV2_0() (request *MbpBatchRollWidByTagIdRequestV2_0) {
	request = &MbpBatchRollWidByTagIdRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "2_0", "mbp/batchRollWidByTagId", "api")
	request.Method = api.POST
	return
}

func CreateMbpBatchRollWidByTagIdResponseV2_0() (response *api.BaseResponse[MbpBatchRollWidByTagIdResponseV2_0]) {
	response = api.CreateResponse[MbpBatchRollWidByTagIdResponseV2_0](&MbpBatchRollWidByTagIdResponseV2_0{})
	return
}
