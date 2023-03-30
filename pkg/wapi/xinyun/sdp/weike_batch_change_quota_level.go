package sdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// WeikeBatchChangeQuotaLevel
// @id 1329
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量修改微客等级)
func (client *Client) WeikeBatchChangeQuotaLevel(request *WeikeBatchChangeQuotaLevelRequest) (response *WeikeBatchChangeQuotaLevelResponse, err error) {
	rpcResponse := CreateWeikeBatchChangeQuotaLevelResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WeikeBatchChangeQuotaLevelRequest struct {
	*api.BaseRequest

	WidList       []int64 `json:"widList,omitempty"`
	TargetLevelId int     `json:"targetLevelId,omitempty"`
}

type WeikeBatchChangeQuotaLevelResponse struct {
	FailSize        int     `json:"failSize,omitempty"`
	FailWeikeInfos  []int64 `json:"failWeikeInfos,omitempty"`
	NotWeikeWidList []int64 `json:"notWeikeWidList,omitempty"`
	RequestSize     []int64 `json:"requestSize,omitempty"`
}

func CreateWeikeBatchChangeQuotaLevelRequest() (request *WeikeBatchChangeQuotaLevelRequest) {
	request = &WeikeBatchChangeQuotaLevelRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("sdp", "1_0", "weike/batchChangeQuotaLevel", "api")
	request.Method = api.POST
	return
}

func CreateWeikeBatchChangeQuotaLevelResponse() (response *api.BaseResponse[WeikeBatchChangeQuotaLevelResponse]) {
	response = api.CreateResponse[WeikeBatchChangeQuotaLevelResponse](&WeikeBatchChangeQuotaLevelResponse{})
	return
}
