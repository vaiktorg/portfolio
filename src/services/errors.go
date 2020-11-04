package services

type ErrorManager struct {
	err     string
	updater func(error)
}

var (
	ErrMan = new(ErrorManager)
)

func CheckError(err error) {
	defer func() {
		if e := recover(); e != nil && ErrMan.updater != nil {
			ErrMan.updater(e.(error))
		}
	}()

	if err != nil {
		panic(err)
	}
}

func (e *ErrorManager) OnError(updater func(err error)) {
	e.updater = updater
}
