package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabaseConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_SCHEMA"),
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return db, err
	}

	return db, err
}

func MigrateTables(databaseConnection *gorm.DB) {
	databaseConnection.AutoMigrate(
		&User{},
		&Expense{},
	)
}

func CreateUser(user User) error {
	err := DBConnection.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func GetUser(email string) (User, error) {
	var user User
	err := DBConnection.Where("email = ?", email).First(&user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func GetExpense(expenseId int) (Expense, error) {
	var expense Expense
	err := DBConnection.Where("id = ?", expenseId).First(&expense).Error
	if err != nil {
		return Expense{}, err
	}

	return expense, nil
}

func CreateExpense(expense Expense) (Expense, error) {
	err := DBConnection.Create(&expense).Error
	if err != nil {
		return Expense{}, err
	}

	return expense, nil
}

func UpdateExpense(expense Expense) (Expense, error) {
	err := DBConnection.Save(&expense).Error
	if err != nil {
		return Expense{}, err
	}

	return expense, nil
}

func DeleteExpense(expenseId int, userId int) error {
	err := DBConnection.Where("id = ? AND user_id = ?", expenseId, userId).Delete(&Expense{}).Error
	if err != nil {
		return err
	}

	return nil
}

func GetAllExpenses(userId int, startDate time.Time, endDate time.Time, category string) ([]Expense, float64, error) {
	var (
		expenses []Expense
		total    float64
	)

	query := DBConnection.Model(&Expense{}).Where("user_id = ?", userId)

	if category != "" {
		query = query.Where("category = ?", category)
	}

	if !startDate.IsZero() {
		query = query.Where("created_at >= ?", startDate)
	}

	if !endDate.IsZero() {
		query = query.Where("created_at <= ?", endDate)
	}

	err := query.Order("id DESC").Find(&expenses).Error
	if err != nil {
		return expenses, total, err
	}

	err = query.Select("COALESCE(SUM(amount), 0)").Scan(&total).Error
	if err != nil {
		return expenses, total, err
	}

	return expenses, total, nil
}
