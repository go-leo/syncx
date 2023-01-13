package syncx

func CombineChan[T any](buf int, chans ...<-chan T) <-chan T {
	c := make(chan T, buf)
	for _, ch := range chans {
		go func() {
			for v := range ch {
				c <- v
			}
		}()
	}
	return c
}
