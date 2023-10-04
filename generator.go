package goslices

func Generator[T ~[]V, V any](s *T) <-chan V {
	ch := make(chan V)
	go func() {
		for _, v := range *s {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// generator for Map
func MapGen[T ~[]V, V any](s *T, f func(v V) V) <-chan V {
	ch := make(chan V)
	if IsEmpty(s) {
		close(ch)
		return nil
	}
	go func() {
		for _, v := range *s {
			ch <- f(v)
		}
		close(ch)
	}()
	return ch
}

// generator for Filter
func FilterGen[T ~[]V, V any](s *T, f func(v V) bool) <-chan V {
	ch := make(chan V)
	if IsEmpty(s) {
		close(ch)
		return nil
	}
	go func() {
		for _, v := range *s {
			if f(v) {
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}
