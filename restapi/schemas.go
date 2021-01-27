package restapi

type Tags struct {
	Tags     []string `json:"tags"`
	TagName  string   `json:"-"`
	TagId    string   `json:"tag_id_123"`
	TestName string   `json:"tag_name,omitempty"` //не появляется если пустое
}
