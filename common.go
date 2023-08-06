// Copyright (c) 2022 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license
package gostore

import "time"

const (
	durationMax    time.Duration = 1<<63 - 1
	periodMax      time.Duration = 10 * time.Minute
	periodMin      time.Duration = 200 * time.Millisecond
	shiftDefault   time.Duration = 5 * time.Millisecond
	ttlToPeriodDiv time.Duration = 5
)
