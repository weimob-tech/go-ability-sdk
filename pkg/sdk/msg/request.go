package msg

type MsgRequest[T any] interface {
	GetId() string
	GetTopic() string
	GetEvent() string
	GetSign() string
	GetVersion() string
	GetRawMsgBody() string
	GetMsgBody() *T
}

type baseRequest struct {
	Id         string `json:"id,omitempty"`
	Topic      string `json:"topic,omitempty"`
	Event      string `json:"event,omitempty"`
	Sign       string `json:"sign,omitempty"`
	RawMsgBody string `json:"msgBody,omitempty"`
	Version    string `json:"-"`
}

func (r *baseRequest) GetId() string {
	return r.Id
}

func (r *baseRequest) GetTopic() string {
	return r.Topic
}

func (r *baseRequest) GetEvent() string {
	return r.Event
}

func (r *baseRequest) GetVersion() string {
	return r.Version
}

func (r *baseRequest) GetSign() string {
	return r.Sign
}

func (r *baseRequest) GetRawMsgBody() string {
	return r.RawMsgBody
}

type XinyunOpenMessage[T any] struct {
	*baseRequest
	// todo
	Data            *T     `json:"data,omitempty"`
	BusinessId      string `json:"business_id,omitempty"`
	PublicAccountId string `json:"public_account_id,omitempty"`
	MsgSignature    string `json:"msgSignature,omitempty"`
	SaasChannel     string `json:"saas_channel,omitempty"`
	SaasClientId    string `json:"saas_clientId,omitempty"`
	MsgBody         *T
}

func (r *XinyunOpenMessage[T]) GetMsgBody() *T {
	return r.MsgBody
}

type WosOpenMessage[T any] struct {
	*baseRequest
	// todo
	BosId        int64  `json:"bosId,omitempty"`
	SaasChannel  string `json:"saasChannel,omitempty"`
	SaasClientId bool   `json:"saasClientId,omitempty"`
	MsgBody      *T
}

func (r *WosOpenMessage[T]) GetMsgBody() *T {
	return r.MsgBody
}
