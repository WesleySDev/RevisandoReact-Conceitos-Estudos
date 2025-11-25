package models

import "time"

type StudyGoal struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.time `json:"deadline"`
}
type StudeTask struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	GoalID    int       `json:"goal_id"`
	Duration  int       `json:"duration"` // in minutes
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"username"`
	Email string `json:"email"`
}
