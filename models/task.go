package models

import "time"

// Task field composition
type Task struct {
	ID, Priority                               int
	Title, Description, Status                 string
	Dependency                                 []int
	Deadline, TimeEstimate, workStart, workEnd time.Time
}
