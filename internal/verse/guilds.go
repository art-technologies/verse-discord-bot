package verse

type Guilds []GuildDetails

type GuildDetails struct {
	ServerID  string  `json:"server_id"`
	ChannelID string  `json:"channel_id"`
	Filters   *Filter `json:"filters"`
}

type Filter struct {
	Collections   *[]string `json:"collections"`
	Collaborators *[]string `json:"collaborators"`
	Artists       *[]string `json:"artists"`
	Events        *[]string `json:"events"`
}
