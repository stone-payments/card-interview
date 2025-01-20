package schema

type AuthorizerResponse struct {
	Status      string `json:"status"`
	AuthorizeID string `json:"authorize_id"`
	Error       string `json:"errors"`
	Warning     string `json:"warning"`
}
