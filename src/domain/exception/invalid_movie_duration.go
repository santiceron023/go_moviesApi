package exception

type InvalidMovieDuration struct {
	ErrMessage string
}

func (e InvalidMovieDuration) Error() string {
	return e.ErrMessage
}

func (e InvalidMovieDuration) IsBusinessLogic() bool {
	return true
}
