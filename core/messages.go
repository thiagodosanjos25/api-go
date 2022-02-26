package core

// ErrMessage is return message default
type ErrMessage struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Erro    string `json:"erro"`
}

// ErrDetail ...
type ErrDetail struct {
	Err error `json:"erro"`
}

// SuccessMessage is return Zen message
type SuccessMessage struct {
	Message string `json:"message"`
}

// VersionMessage ...
type VersionMessage struct {
	AppID          string `json:"appID"`
	AppName        string `json:"appName"`
	ServerID       string `json:"serverID"`
	CreatedAt      string `json:"createdAt"`
	ReleaseVersion string `json:"version"`
	Commit         string `json:"commit"`
	Description    string `json:"description"`
}

// StatusSessaoAuthResp ...
type StatusSessaoAuthResp struct {
	CodResposta string `json:"codResposta"`
	Mensagem    string `json:"mensagem"`
	IDStatus    int    `json:"idStatus"`
}
