package schemas

type RequestEvent struct {
	Type        string `json:"type"`
	Destination string `json:"destination,omitempty"`
	Origin      string `json:"origin,omitempty"`
	Amount      int    `json:"amount"`
}

type Account struct {
	ID      string `json:"id"`
	Balance int    `json:"balance"`
}

type EventResponse struct {
	Origin      *Account `json:"origin,omitempty"`
	Destination *Account `json:"destination,omitempty"`
}
