package ugo

import "time"

func Protect(f func(), onPanic ...func(exception interface{})) {
	defer func() {
		if err := recover(); err != nil {
			if len(onPanic) > 0 {
				onPanic[0](err)
			}
		}
	}()
	f()
}

func Retry(f func() error, maxRetries int, interval time.Duration) error {
	var err error
	for i := 0; i <= maxRetries; i++ {
		err = f()
		if err == nil {
			return nil
		}
		if i < maxRetries {
			time.Sleep(interval)
		}
	}
	return err
}
