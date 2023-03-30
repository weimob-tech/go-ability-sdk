package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ContactCreate
// @id 1729
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新建联系人)
func (client *Client) ContactCreate(request *ContactCreateRequest) (response *ContactCreateResponse, err error) {
	rpcResponse := CreateContactCreateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ContactCreateRequest struct {
	*api.BaseRequest

	ContactsInfo ContactCreateRequestContactsInfo `json:"contactsInfo,omitempty"`
	Wid          int64                            `json:"wid,omitempty"`
}

type ContactCreateResponse struct {
}

func CreateContactCreateRequest() (request *ContactCreateRequest) {
	request = &ContactCreateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "contact/create", "api")
	request.Method = api.POST
	return
}

func CreateContactCreateResponse() (response *api.BaseResponse[ContactCreateResponse]) {
	response = api.CreateResponse[ContactCreateResponse](&ContactCreateResponse{})
	return
}

type ContactCreateRequestContactsInfo struct {
	SAvatar        string `json:"s_avatar,omitempty"`
	SBirthday      string `json:"s_birthday,omitempty"`
	SContactName   string `json:"s_contact_name,omitempty"`
	SCreateTime    string `json:"s_create_time,omitempty"`
	SCreateUserWid int64  `json:"s_create_user_wid,omitempty"`
	SCustomerKey   int64  `json:"s_customer_key,omitempty"`
	SEmail         string `json:"s_email,omitempty"`
	SFormerOwner   int64  `json:"s_former_owner,omitempty"`
	SGender        string `json:"s_gender,omitempty"`
	SMobilePhone   string `json:"s_mobile_phone,omitempty"`
	SOwner         int64  `json:"s_owner,omitempty"`
	SPost          string `json:"s_post,omitempty"`
	SQq            string `json:"s_qq,omitempty"`
	STelephone     string `json:"s_telephone,omitempty"`
	SUpdateTime    string `json:"s_update_time,omitempty"`
	SUpdateUserWid int64  `json:"s_update_user_wid,omitempty"`
	SWechat        string `json:"s_wechat,omitempty"`
}
