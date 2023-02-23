package event

type MessageTextEventSignal struct {
	Frame
	data *MesssageTextEvent `json:"d"`
}
type MesssageTextEvent struct {
	BaseEvent
	Extra *MessageTextExtra `json:"extra"`
}
type MessageTextExtra struct {
	Type         int         `json:"type"`
	GuildId      string      `json:"guild_id"`
	ChannelName  string      `json:"channel_name"`
	Mention      string      `json:"mention"`
	MentionAll   bool        `json:"mention_all"`
	MentionRoles []string    `json:"mention_roles"`
	MentionHere  bool        `json:"mention_here"`
	Author       *TextAuthor `json:"author"`
}
type TextAuthor struct {
	IdentifyNum string   `json:"identify_num"`
	Avatar      string   `json:"avatar"`
	Username    string   `json:"username"`
	Id          string   `json:"id"`
	Nickname    string   `json:"nickname"`
	roles       []string `json:"roles"`
}
