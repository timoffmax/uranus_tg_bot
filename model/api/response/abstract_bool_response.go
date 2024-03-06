package response

/*
*
Represents a bool response
Is used in order to have unified general and processing logic with JSON responses
*/
type AbstractBoolResponse struct {
	Success bool
}

/*
*
Parses response body into the struct
*/
func (r *AbstractBoolResponse) UnmarshalJSON(body []byte) error {
	var isSuccessful bool

	respString := string(body)

	switch respString {
	case "true":
		isSuccessful = true
		break

	default:
		isSuccessful = false
	}

	r.Success = isSuccessful

	return nil
}
