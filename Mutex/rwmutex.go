package main

import (
	"sync"
	"time"
)

type InfoClient struct {
	mu        sync.RWMutex
	token     string
	tokenTime time.Time
	TTL       time.Duration
}

func (i *InfoClient) CheckToken() (string, time.Duration) {
	// Allows multiple readers to use data at the same time but one writer at a time
	i.mu.Lock()
	defer i.mu.RUnlock()

	return i.token, i.TTL - time.Since(i.tokenTime)
}
