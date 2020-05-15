package exception

type InvalidMovieLength struct {
	ErrMessage string
}

func (e InvalidMovieLength) Error() string {
	return e.ErrMessage
}

func (e InvalidMovieLength) IsBusinessLogic() bool {
	return true
}
