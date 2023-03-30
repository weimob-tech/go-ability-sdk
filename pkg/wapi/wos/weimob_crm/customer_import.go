package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerImport
// @id 1503
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1503?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=导入客户)
func (client *Client) CustomerImport(request *CustomerImportRequest) (response *CustomerImportResponse, err error) {
	rpcResponse := CreateCustomerImportResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerImportRequest struct {
	*api.BaseRequest

	UserList         []CustomerImportRequestUserList `json:"userList,omitempty"`
	ImportType       int64                           `json:"importType,omitempty"`
	MembershipPlanId int64                           `json:"membershipPlanId,omitempty"`
	NeedGetCardGift  int64                           `json:"needGetCardGift,omitempty"`
}

type CustomerImportResponse struct {
	SuccessList []CustomerImportResponseSuccessList `json:"successList,omitempty"`
	ErrorList   []CustomerImportResponseErrorList   `json:"errorList,omitempty"`
	SuccessCnt  int64                               `json:"successCnt,omitempty"`
	ErrorCnt    int64                               `json:"errorCnt,omitempty"`
}

func CreateCustomerImportRequest() (request *CustomerImportRequest) {
	request = &CustomerImportRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "customer/import", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerImportResponse() (response *api.BaseResponse[CustomerImportResponse]) {
	response = api.CreateResponse[CustomerImportResponse](&CustomerImportResponse{})
	return
}

type CustomerImportRequestUserList struct {
	ExtMap                CustomerImportRequestExtMap `json:"extMap,omitempty"`
	Phone                 string                      `json:"phone,omitempty"`
	RegionCode            string                      `json:"regionCode,omitempty"`
	UnionId               string                      `json:"unionId,omitempty"`
	OpenId                string                      `json:"openId,omitempty"`
	AppId                 string                      `json:"appId,omitempty"`
	AppChannel            int64                       `json:"appChannel,omitempty"`
	UserName              string                      `json:"userName,omitempty"`
	Gender                int64                       `json:"gender,omitempty"`
	Birthday              int64                       `json:"birthday,omitempty"`
	IdentityCardNum       string                      `json:"identityCardNum,omitempty"`
	Education             string                      `json:"education,omitempty"`
	Income                string                      `json:"income,omitempty"`
	Industry              string                      `json:"industry,omitempty"`
	Hobby                 string                      `json:"hobby,omitempty"`
	Province              string                      `json:"province,omitempty"`
	City                  string                      `json:"city,omitempty"`
	Area                  string                      `json:"area,omitempty"`
	Address               string                      `json:"address,omitempty"`
	BelongVidName         string                      `json:"belongVidName,omitempty"`
	BelongVid             int64                       `json:"belongVid,omitempty"`
	CustomCardNo          string                      `json:"customCardNo,omitempty"`
	LevelName             string                      `json:"levelName,omitempty"`
	Growth                int64                       `json:"growth,omitempty"`
	Balance               int64                       `json:"balance,omitempty"`
	TotalBalance          int64                       `json:"totalBalance,omitempty"`
	Point                 int64                       `json:"point,omitempty"`
	TotalPoint            int64                       `json:"totalPoint,omitempty"`
	Tag                   string                      `json:"tag,omitempty"`
	GuiderPhone           int64                       `json:"guiderPhone,omitempty"`
	SourceVid             int64                       `json:"sourceVid,omitempty"`
	UserKey               string                      `json:"userKey,omitempty"`
	BecomeMemberTime      int64                       `json:"becomeMemberTime,omitempty"`
	MembershipCardChannel int64                       `json:"membershipCardChannel,omitempty"`
}

type CustomerImportRequestExtMap struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type CustomerImportResponseSuccessList struct {
	Wid     int64  `json:"wid,omitempty"`
	Phone   string `json:"phone,omitempty"`
	UserKey string `json:"userKey,omitempty"`
}

type CustomerImportResponseErrorList struct {
	Phone        string `json:"phone,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	Wid          int64  `json:"wid,omitempty"`
	UserKey      string `json:"userKey,omitempty"`
}
