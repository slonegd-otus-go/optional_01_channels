package channels

func Merge(args ... <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	for _, channel := range args {
		go func(channel <-chan interface{}) {
			for {
				value, ok := <- channel
				if !ok {
					return
				}
				out <- value
			}
		}(channel)
	}
	return out
}

