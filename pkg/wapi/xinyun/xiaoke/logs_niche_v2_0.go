package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// LogsNiche
// @id 3068
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询商机跟进记录（下线）)
func (client *Client) LogsNicheV2_0(request *LogsNicheRequestV2_0) (response *LogsNicheResponseV2_0, err error) {
	rpcResponse := CreateLogsNicheResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type LogsNicheRequestV2_0 struct {
	*api.BaseRequest
}

type LogsNicheResponseV2_0 struct {
}

func CreateLogsNicheRequestV2_0() (request *LogsNicheRequestV2_0) {
	request = &LogsNicheRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "2_0", "logs/niche", "api")
	request.Method = api.POST
	return
}

func CreateLogsNicheResponseV2_0() (response *api.BaseResponse[LogsNicheResponseV2_0]) {
	response = api.CreateResponse[LogsNicheResponseV2_0](&LogsNicheResponseV2_0{})
	return
}
