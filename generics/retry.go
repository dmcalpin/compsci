package generics

import "time"

func Retry(f func() error, retries int, delay time.Duration) error {
	var err error
	for i := 0; i < retries; i++ {
		err = f()
		if err != nil {
			time.Sleep(delay)
			continue
		}
		break
	}
	return err
}
