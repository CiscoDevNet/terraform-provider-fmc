package fmc

type semaphore chan struct{}

func Semaphore(n int) *semaphore {
	sem := make(semaphore, 10)
	return &sem
}

// acquire n resources
func (s semaphore) P(n int) {
	e := struct{}{}
	for i := 0; i < n; i++ {
		s <- e
	}
}

// release n resources
func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

/* mutexes */

func (s semaphore) Lock() {
	s.P(1)
}

func (s semaphore) Unlock() {
	s.V(1)
}
