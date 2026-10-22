package application

//==========================================
// error code enum
//==========================================

type ErrorCode int

const (
	// chat
	ChatUserCancelledError ErrorCode = 2000
	// network
	NetworkRequestError ErrorCode = 4000
	// internal
	InternalServerError ErrorCode = 5000
)

func (s ErrorCode) Type() string {
	switch s {
	case ChatUserCancelledError:
		return "ChatUserCancelledError"
	case InternalServerError:
		return "InternalServerError"
	case NetworkRequestError:
		return "NetworkRequestError"
	default:
		return ""
	}
}
