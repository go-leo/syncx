package syncx

func CombineChannels[T any](buf int, channels ...<-chan T) <-chan T {
	c := make(chan T, buf)
	for _, ch := range channels {
		go func() {
			for v := range ch {
				c <- v
			}
		}()
	}
	return c
}

func ToSendChannels[T any](channels ...chan T) []<-chan T {
	c := make([]<-chan T, len(channels))
	for _, ch := range channels {
		c = append(c, ch)
	}
	return c
}

func ToReceiveChannels[T any](channels ...chan T) []chan<- T {
	c := make([]chan<- T, len(channels))
	for _, ch := range channels {
		c = append(c, ch)
	}
	return c
}
