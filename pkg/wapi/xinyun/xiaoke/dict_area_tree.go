package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DictAreaTree
// @id 1644
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询全部省市区)
func (client *Client) DictAreaTree(request *DictAreaTreeRequest) (response *DictAreaTreeResponse, err error) {
	rpcResponse := CreateDictAreaTreeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DictAreaTreeRequest struct {
	*api.BaseRequest
}

type DictAreaTreeResponse struct {
	AreaInfoList []map[string]any `json:"areaInfoList,omitempty"`
	AreaCode     string           `json:"areaCode,omitempty"`
	AreaName     string           `json:"areaName,omitempty"`
	ParentCode   string           `json:"parentCode,omitempty"`
	LevelType    int64            `json:"levelType,omitempty"`
	Id           int64            `json:"id,omitempty"`
}

func CreateDictAreaTreeRequest() (request *DictAreaTreeRequest) {
	request = &DictAreaTreeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "dict/areaTree", "api")
	request.Method = api.POST
	return
}

func CreateDictAreaTreeResponse() (response *api.BaseResponse[DictAreaTreeResponse]) {
	response = api.CreateResponse[DictAreaTreeResponse](&DictAreaTreeResponse{})
	return
}
