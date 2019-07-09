package models

//CommonResponse ...
type CommonResponse struct {
	// Status contains the status of the request, and may contain debugging
	// information to help you track down why the call failed.
	Status string `json:"status,omitempty"`

	// ErrorMessage is the explanatory field added when Status is an error.
	ErrorMessage string `json:"error_message,omitempty"`
}
