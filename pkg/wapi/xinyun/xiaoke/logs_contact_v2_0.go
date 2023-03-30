package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// LogsContact
// @id 3067
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询联系人跟进记录)
func (client *Client) LogsContactV2_0(request *LogsContactRequestV2_0) (response *LogsContactResponseV2_0, err error) {
	rpcResponse := CreateLogsContactResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type LogsContactRequestV2_0 struct {
	*api.BaseRequest
}

type LogsContactResponseV2_0 struct {
}

func CreateLogsContactRequestV2_0() (request *LogsContactRequestV2_0) {
	request = &LogsContactRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "2_0", "logs/contact", "api")
	request.Method = api.POST
	return
}

func CreateLogsContactResponseV2_0() (response *api.BaseResponse[LogsContactResponseV2_0]) {
	response = api.CreateResponse[LogsContactResponseV2_0](&LogsContactResponseV2_0{})
	return
}
