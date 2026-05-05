package c

type JSONTagMismatchAllowed struct {
	ID     string `json:"vmID,omitempty"`
	IPAddr string `json:"vmIp,omitempty"`
}
