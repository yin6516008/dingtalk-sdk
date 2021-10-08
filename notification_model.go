package dingtalk

type SendWorkerNotificationParams struct {
	ToAllUser  string      `json:"to_all_user,omitempty"`
	AgentID    string      `json:"agent_id"`
	DeptIDList string      `json:"dept_id_list,omitempty"`
	UseridList string      `json:"userid_list,omitempty"`
	MSG        interface{} `json:"msg"`
}

type SendWorkerNotificationRes struct {
	Errcode   int64  `json:"errcode"`
	TaskID    int64  `json:"task_id"`
	RequestID string `json:"request_id"`
}

type GetWorkerNotificationResultParams struct {
	AgentID int64 `json:"agent_id"`
	TaskID  int64 `json:"task_id"`
}

type GetWorkerNotificationResultRes struct {
	Errcode    int64      `json:"errcode"`
	SendResult SendResult `json:"send_result"`
	RequestID  string     `json:"request_id"`
}

type SendResult struct {
	FailedUserIDList  []interface{} `json:"failed_user_id_list"`
	ForbiddenList     []interface{} `json:"forbidden_list"`
	InvalidDeptIDList []interface{} `json:"invalid_dept_id_list"`
	InvalidUserIDList []interface{} `json:"invalid_user_id_list"`
	ReadUserIDList    []string      `json:"read_user_id_list"`
	UnreadUserIDList  []interface{} `json:"unread_user_id_list"`
}

// 文本类型消息结构体
type TextMessage struct {
	Msgtype string `json:"msgtype"`
	Text    Text   `json:"text"`
}

type Text struct {
	Content string `json:"content"`
}

// Markdown类型消息结构体
type MarkdownMessage struct {
	Msgtype  string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
