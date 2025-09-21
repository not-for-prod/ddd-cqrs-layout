package config

import "time"

type WorkerPool struct {
	Count    int
	Interval time.Duration
}
