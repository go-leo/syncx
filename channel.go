package syncx

import "sync"

func CombineChannels[T any](buf int, channels ...<-chan T) <-chan T {
	c := make(chan T, buf)
	var wg sync.WaitGroup
	for _, ch := range channels {
		wg.Add(1)
		go func(ch <-chan T) {
			defer wg.Done()
			for v := range ch {
				c <- v
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(c)
	}()
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
