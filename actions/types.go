package actions

type ResponseError struct {
	FriendlyErrorMessage string
	ErrorMessage         string
	ErrorCode            string
}

func (r ResponseError) Error() string {
	return r.ErrorMessage
}

type GetStatusResponse struct {
	CurrentState string
}
