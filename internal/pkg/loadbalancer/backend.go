package loadbalancer

import "sync"

type backend struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	isDead bool
	mu     sync.Mutex
}

func (b *backend) SetState(state bool) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.isDead = state
}

func (b *backend) IsDown() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	isAlive := b.isDead

	return isAlive
}
