package exception

type InvalidMovieDuration struct {
	ErrMessage string
}

func (e InvalidMovieDuration) Error() string {
	return e.ErrMessage
}