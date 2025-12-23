package main

import "time"

type User struct {
	ID        int `gorm:"primaryKey,autoIncrement"`
	Name      string
	Email     string
	Password  string
	Expense   []Expense
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Expense struct {
	ID          int `gorm:"primaryKey,autoIncrement"`
	Title       string
	Description string
	Amount      float64
	Category    string
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
