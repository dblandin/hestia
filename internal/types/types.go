package types

type Input struct {
	Command string
	Args    string
}

type Event struct {
	Type           string `json:"type"`
	User           string `json:"user"`
	Text           string `json:"text"`
	Timestamp      string `json:"ts"`
	Channel        string `json:"channel"`
	EventTimestamp string `json:"event_ts"`
}

type EventCallback struct {
	Token       string   `json:"type"`
	Challenge   string   `json:"challenge"`
	TeamId      string   `json:"team_id"`
	ApiAppId    string   `json:"api_app_id"`
	Event       Event    `json:"event"`
	Type        string   `json:"type"`
	EventId     string   `json:"event_id"`
	EventTime   int      `json:"event_time"`
	AuthedUsers []string `json:"authed_users"`
}

// {
//     "token": "FeCDfP96MxGb3JA2TTmXVhmc",
//     "team_id": "T0G0RAKPG",
//     "api_app_id": "A9FDQB5V5",
//     "event": {
//         "type": "app_mention",
//         "user": "U0G0RMBC0",
//         "text": "<@U9EC5EG3U> sup?",
//         "ts": "1519519253.000080",
//         "channel": "C0G0KFXS8",
//         "event_ts": "1519519253000080"
//     },
//     "type": "event_callback",
//     "event_id": "Ev9E9CLT99",
//     "event_time": 1519519253000080,
//     "authed_users": [
//         "U9EC5EG3U"
//     ]
// }
