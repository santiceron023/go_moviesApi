package port

type UsersRepository interface {
	CheckUser(userId int64) error
	UpdateMovieCount(userId int64) error
}
