package weimob_guide

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuiderGet
// @id 1388
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1388?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询导购信息)
func (client *Client) GuiderGet(request *GuiderGetRequest) (response *GuiderGetResponse, err error) {
	rpcResponse := CreateGuiderGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuiderGetRequest struct {
	*api.BaseRequest

	GuiderWid        int64  `json:"guiderWid,omitempty"`
	NeedGuiderPicUrl int64  `json:"needGuiderPicUrl,omitempty"`
	GuiderPhone      string `json:"guiderPhone,omitempty"`
	GuiderId         string `json:"guiderId,omitempty"`
	GuiderVid        int64  `json:"guiderVid,omitempty"`
	JobNumber        string `json:"jobNumber,omitempty"`
}

type GuiderGetResponse struct {
	GuiderVid      int64  `json:"guiderVid,omitempty"`
	Zone           string `json:"zone,omitempty"`
	VidName        string `json:"vidName,omitempty"`
	JobNumber      string `json:"jobNumber,omitempty"`
	IsUsed         int64  `json:"isUsed,omitempty"`
	GuiderWid      int64  `json:"guiderWid,omitempty"`
	IsExclusiveCus int64  `json:"isExclusiveCus,omitempty"`
	GuiderName     string `json:"guiderName,omitempty"`
	GuiderPhone    string `json:"guiderPhone,omitempty"`
	GuiderId       string `json:"guiderId,omitempty"`
}

func CreateGuiderGetRequest() (request *GuiderGetRequest) {
	request = &GuiderGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_guide", "v2.0", "guider/get", "apigw")
	request.Method = api.POST
	return
}

func CreateGuiderGetResponse() (response *api.BaseResponse[GuiderGetResponse]) {
	response = api.CreateResponse[GuiderGetResponse](&GuiderGetResponse{})
	return
}
