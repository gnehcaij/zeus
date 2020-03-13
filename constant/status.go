package constant

type CommonStatus struct {
	StatusCode int32
	StatusMsg  string
}

const (
	COMMON int32 = (1 << 10) * 0 // 0
	API    int32 = (1<<10)*1 + 1 // 2049
)

var (
	SUCCESS      = &CommonStatus{COMMON + 0, ""}
	ERR_INTERNAL = &CommonStatus{COMMON + 4, "internal error"}

	ERR_NO_USER = &CommonStatus{COMMON + 11, "the user does not exist"}
)
