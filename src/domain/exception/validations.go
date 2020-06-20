package exception

type Validations interface {
	Error() string
	IsBussinesLogic() bool
}