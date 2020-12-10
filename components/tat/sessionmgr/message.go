package sessionmgr

type SessionAuth struct {
	Type string
	Data struct {
		AppID      int `json:"AppId"`
		Uin        string
		SessionID  string `json:"SessionId"`
		InstanceID string `json:"InstanceId"`
	}
}

type PtyStart struct {
	Type string
	Data struct {
		SessionID string `json:"SessionId"`
		Cols      int
		Rows      int
	}
}

type PtyInput struct {
	Type string
	Data struct {
		SessionID string `json:"SessionId"`
		Input     string
	}
}

type PtyReady struct {
	Type string
	Data struct {
		SessionID string `json:"SessionId"`
	}
}

type PtyOutput struct {
	Type string
	Data struct {
		SessionID string `json:"SessionId"`
		Output    string
	}
}
