package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ClueGiveupPage
// @id 1654
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询放弃原因)
func (client *Client) ClueGiveupPage(request *ClueGiveupPageRequest) (response *ClueGiveupPageResponse, err error) {
	rpcResponse := CreateClueGiveupPageResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ClueGiveupPageRequest struct {
	*api.BaseRequest

	Wid int `json:"wid,omitempty"`
}

type ClueGiveupPageResponse struct {
	ChooseOptionKey  string `json:"chooseOptionKey,omitempty"`
	Code             string `json:"code,omitempty"`
	FieldValue       string `json:"fieldValue,omitempty"`
	Id               int    `json:"id,omitempty"`
	IsCustom         int    `json:"isCustom,omitempty"`
	IsDefaultDisplay int    `json:"isDefaultDisplay,omitempty"`
	Key              int    `json:"key,omitempty"`
	OptionKey        int    `json:"optionKey,omitempty"`
	Sort             int    `json:"sort,omitempty"`
}

func CreateClueGiveupPageRequest() (request *ClueGiveupPageRequest) {
	request = &ClueGiveupPageRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "clue/giveupPage", "api")
	request.Method = api.POST
	return
}

func CreateClueGiveupPageResponse() (response *api.BaseResponse[ClueGiveupPageResponse]) {
	response = api.CreateResponse[ClueGiveupPageResponse](&ClueGiveupPageResponse{})
	return
}
