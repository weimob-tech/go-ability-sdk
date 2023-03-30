package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// LogsCustomer
// @id 1723
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询客户跟进日志)
func (client *Client) LogsCustomer(request *LogsCustomerRequest) (response *LogsCustomerResponse, err error) {
	rpcResponse := CreateLogsCustomerResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type LogsCustomerRequest struct {
	*api.BaseRequest

	Key      string `json:"key,omitempty"`
	Wid      int64  `json:"wid,omitempty"`
	PageNum  int    `json:"pageNum,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
}

type LogsCustomerResponse struct {
}

func CreateLogsCustomerRequest() (request *LogsCustomerRequest) {
	request = &LogsCustomerRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "logs/customer", "api")
	request.Method = api.POST
	return
}

func CreateLogsCustomerResponse() (response *api.BaseResponse[LogsCustomerResponse]) {
	response = api.CreateResponse[LogsCustomerResponse](&LogsCustomerResponse{})
	return
}
