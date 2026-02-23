package d

type JSONTagMismatch struct {
	ID             string   `json:"vmID,omitempty"` // want "field JSONTagMismatch.ID json tag should match the camelCase field name \"id\": got \"vmID\""
	IPAddr         string   `json:"vmIp,omitempty"` // want "field JSONTagMismatch.IPAddr json tag should match the camelCase field name \"ipAddr\": got \"vmIp\""
	IPv6Address    string   `json:"ipv6Address,omitempty"`
	IPv6AddressBad string   `json:"ipv6addressBad,omitempty"` // want "field JSONTagMismatch.IPv6AddressBad json tag should match the camelCase field name \"ipv6AddressBad\": got \"ipv6addressBad\""
	WWIDs          []string `json:"wwids,omitempty"`
	WWIDsBad       []string `json:"wwiDs,omitempty"` // want "field JSONTagMismatch.WWIDsBad json tag should match the camelCase field name \"wwidsBad\": got \"wwiDs\""
	URLs           []string `json:"urls,omitempty"`
	URLsBad        []string `json:"urLs,omitempty"` // want "field JSONTagMismatch.URLsBad json tag should match the camelCase field name \"urlsBad\": got \"urLs\""
}
