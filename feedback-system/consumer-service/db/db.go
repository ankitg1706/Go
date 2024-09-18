package db

import (
    "consumer-service/models"
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

var db *sql.DB

func init() {
    var err error
    db, err = sql.Open("postgres", "postgres://user:password@db:5432/feedback?sslmode=disable")
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal("Error pinging the database:", err)
    }

    log.Println("Database connection established")
}

func SaveFeedback(feedback models.Feedback) error {
    query := `INSERT INTO feedback (username, email, message) VALUES ($1, $2, $3)`
    _, err := db.Exec(query, feedback.Username, feedback.Email, feedback.Message)
    if err != nil {
        return err
    }

    log.Println("Feedback saved to the database")
    return nil
}
