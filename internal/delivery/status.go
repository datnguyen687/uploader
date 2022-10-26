package delivery

type StatusCode int

const (
	StatusCodeOk      StatusCode = 1
	StatusCodeUnknown StatusCode = 0

	StatusCodeUnableToFilterProducts   StatusCode = -1000
	StatusCodeBadFilterProductsRequest StatusCode = -1001
)

var errorMapping map[StatusCode]string = map[StatusCode]string{
	StatusCodeOk:      "ok",
	StatusCodeUnknown: "unknown",

	StatusCodeUnableToFilterProducts:   "unable to filter products",
	StatusCodeBadFilterProductsRequest: "bad filter products request",
}

func (sc StatusCode) Message() string {
	if msg, ok := errorMapping[sc]; ok {
		return msg
	}

	return errorMapping[StatusCodeUnknown]
}
