package redis

type redisError struct {
	Message string
}

func (err *redisError) Error() string {
	return err.Message
}

func New(text string) error {
	return &redisError{text}
}

var(
	ErrNotSameClientLock = New("not same client lock")
)