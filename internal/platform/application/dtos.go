package application

//==========================================
// oarbit error
//==========================================

type OarbitError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Type    string    `json:"type"`
}

func NewOarbitError(code ErrorCode, message string) OarbitError {
	return OarbitError{
		Code:    code,
		Message: message,
		Type:    code.Type(),
	}
}
