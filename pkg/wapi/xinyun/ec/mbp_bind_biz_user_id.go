package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpBindBizUserId
// @id 1817
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=绑定实体卡号到wid)
func (client *Client) MbpBindBizUserId(request *MbpBindBizUserIdRequest) (response *MbpBindBizUserIdResponse, err error) {
	rpcResponse := CreateMbpBindBizUserIdResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpBindBizUserIdRequest struct {
	*api.BaseRequest

	Pid         int64  `json:"pid,omitempty"`
	Wid         int64  `json:"wid,omitempty"`
	BizUserId   string `json:"bizUserId,omitempty"`
	OperateType int    `json:"operateType,omitempty"`
}

type MbpBindBizUserIdResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateMbpBindBizUserIdRequest() (request *MbpBindBizUserIdRequest) {
	request = &MbpBindBizUserIdRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "mbp/bindBizUserId", "api")
	request.Method = api.POST
	return
}

func CreateMbpBindBizUserIdResponse() (response *api.BaseResponse[MbpBindBizUserIdResponse]) {
	response = api.CreateResponse[MbpBindBizUserIdResponse](&MbpBindBizUserIdResponse{})
	return
}
