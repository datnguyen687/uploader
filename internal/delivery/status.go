package delivery

type StatusCode int

const (
	StatusCodeOk      StatusCode = 1
	StatusCodeUnknown StatusCode = 0

	// user
	StatusCodeUnableToBatchUser   StatusCode = -1000
	StatusCodeBadBatchUserRequest StatusCode = -1001
)

var errorMapping map[StatusCode]string = map[StatusCode]string{
	StatusCodeOk:      "ok",
	StatusCodeUnknown: "unknown",

	// user
	StatusCodeUnableToBatchUser:   "unable to batch user",
	StatusCodeBadBatchUserRequest: "bad batch user request",
}

func (sc StatusCode) Message() string {
	if msg, ok := errorMapping[sc]; ok {
		return msg
	}

	return errorMapping[StatusCodeUnknown]
}
