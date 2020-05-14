package exception

type InvalidMovieName struct {
	ErrMessage string
}

func (e InvalidMovieName) Error() string {
	return e.ErrMessage
}