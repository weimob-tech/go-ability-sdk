package weimob_bi

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserBehaviorsGet
// @id 2184
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2184?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取用户行为记录)
func (client *Client) UserBehaviorsGet(request *UserBehaviorsGetRequest) (response *UserBehaviorsGetResponse, err error) {
	rpcResponse := CreateUserBehaviorsGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserBehaviorsGetRequest struct {
	*api.BaseRequest

	SearchDate string `json:"searchDate,omitempty"`
	Hour       int64  `json:"hour,omitempty"`
	Limit      int64  `json:"limit,omitempty"`
	CursorId   string `json:"cursorId,omitempty"`
}

type UserBehaviorsGetResponse struct {
	List         []UserBehaviorsGetResponselist `json:"list,omitempty"`
	Total        int64                          `json:"total,omitempty"`
	NextCursorId string                         `json:"nextCursorId,omitempty"`
}

func CreateUserBehaviorsGetRequest() (request *UserBehaviorsGetRequest) {
	request = &UserBehaviorsGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_bi", "v2.0", "user/behaviors/get", "apigw")
	request.Method = api.POST
	return
}

func CreateUserBehaviorsGetResponse() (response *api.BaseResponse[UserBehaviorsGetResponse]) {
	response = api.CreateResponse[UserBehaviorsGetResponse](&UserBehaviorsGetResponse{})
	return
}

type UserBehaviorsGetResponselist struct {
	TargetProductInstanceId string `json:"targetProductInstanceId,omitempty"`
	VisitDuration           string `json:"visitDuration,omitempty"`
	Cuid                    string `json:"cuid,omitempty"`
	SourceEnd               string `json:"sourceEnd,omitempty"`
	OpenID                  string `json:"openID,omitempty"`
	IP                      string `json:"IP,omitempty"`
	En                      string `json:"en,omitempty"`
	Sourcename              string `json:"sourcename,omitempty"`
	SessionId               string `json:"sessionId,omitempty"`
	PageName                string `json:"pageName,omitempty"`
	Url                     string `json:"url,omitempty"`
	Scene                   string `json:"scene,omitempty"`
	Vid                     string `json:"vid,omitempty"`
	TargetProductId         string `json:"targetProductId,omitempty"`
	Wid                     string `json:"wid,omitempty"`
	JsonData                string `json:"jsonData,omitempty"`
	BosId                   string `json:"bosId,omitempty"`
	DataSource              string `json:"dataSource,omitempty"`
}
