package streamlabs

type DonationEvent struct {
	EventID string `json:"event_id"`
	For     string `json:"for"`
	Message []struct {
		ID              string `json:"_id"`
		Amount          int    `json:"amount"`
		Currency        string `json:"currency"`
		FormattedAmount string `json:"formatted_amount"`
		From            string `json:"from"`
		FromUserID      int    `json:"from_user_id"`
		IsTest          bool   `json:"isTest"`
		Message         string `json:"message"`
		Name            string `json:"name"`
		Priority        int    `json:"priority"`
		To              struct {
			Name string `json:"name"`
		} `json:"to"`
	} `json:"message"`
}

type SuperchatEvent struct {
	EventID string `json:"event_id"`
	For     string `json:"for"`
	Message []struct {
		ID            string  `json:"_id"`
		Amount        float64 `json:"amount"`
		ChannelID     string  `json:"channelId"`
		Comment       string  `json:"comment"`
		Currency      string  `json:"currency"`
		DisplayString string  `json:"displayString"`
		EventID       int     `json:"id"`
		IsTest        bool    `json:"isTest"`
		Name          string  `json:"name"`
		Payload       struct {
			Currency string `json:"currency"`
		} `json:"payload"`
		Priority int `json:"priority"`
	} `json:"message"`
}

type FollowEvent struct {
	EventID string `json:"event_id"`
	For     string `json:"for"`
	Message []struct {
		ID       string `json:"_id"`
		IsTest   bool   `json:"isTest"`
		Name     string `json:"name"`
		Priority int    `json:"priority"`
	} `json:"message"`
}

type SubscriptionEvent struct {
	EventID string `json:"event_id"`
	For     string `json:"for"`
	Message []struct {
		ID       string      `json:"_id"`
		Emotes   interface{} `json:"emotes"` // Adjust the type based on the actual data type
		IsTest   bool        `json:"isTest"`
		Message  string      `json:"message"`
		Months   int         `json:"months"`
		Name     string      `json:"name"`
		Priority int         `json:"priority"`
		SubPlan  string      `json:"sub_plan"`
	} `json:"message"`
}

func (d DonationEvent) Type() string {
	return "donation"
}

func (f FollowEvent) Type() string {
	return "follow"
}

func (s SuperchatEvent) Type() string {
	return "superchat"
}

func (s SubscriptionEvent) Type() string {
	return "subscription"
}
