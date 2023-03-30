package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuideAddGuiderReturnWid
// @id 1704
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增导购返回wid)
func (client *Client) GuideAddGuiderReturnWid(request *GuideAddGuiderReturnWidRequest) (response *GuideAddGuiderReturnWidResponse, err error) {
	rpcResponse := CreateGuideAddGuiderReturnWidResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuideAddGuiderReturnWidRequest struct {
	*api.BaseRequest

	Data           GuideAddGuiderReturnWidRequestData `json:"data,omitempty"`
	StoreNumber    string                             `json:"storeNumber,omitempty"`
	Zone           string                             `json:"zone,omitempty"`
	GuiderPhone    string                             `json:"guiderPhone,omitempty"`
	GuiderName     string                             `json:"guiderName,omitempty"`
	JobNumber      string                             `json:"jobNumber,omitempty"`
	IsUsed         int                                `json:"isUsed,omitempty"`
	IsExclusiveCus int                                `json:"isExclusiveCus,omitempty"`
}

type GuideAddGuiderReturnWidResponse struct {
}

func CreateGuideAddGuiderReturnWidRequest() (request *GuideAddGuiderReturnWidRequest) {
	request = &GuideAddGuiderReturnWidRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "guide/addGuiderReturnWid", "api")
	request.Method = api.POST
	return
}

func CreateGuideAddGuiderReturnWidResponse() (response *api.BaseResponse[GuideAddGuiderReturnWidResponse]) {
	response = api.CreateResponse[GuideAddGuiderReturnWidResponse](&GuideAddGuiderReturnWidResponse{})
	return
}

type GuideAddGuiderReturnWidRequestData struct {
}
