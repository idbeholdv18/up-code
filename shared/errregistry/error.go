package errregistry

type APIError struct {
	HTTPStatus int    `json:"-"`
	Code       string `json:"code"`
	Message    any    `json:"message"`
	RequestID  string `json:"request_id,omitempty"`
}
