package errors

type ApiError struct {
	Status           int    `json:"-"`
	ShowInProduction bool   `json:"-"`
	Code             int    `json:"code"`
	Message          string `json:"message"`
}

func (err ApiError) Error() string {
	return err.Message
}
