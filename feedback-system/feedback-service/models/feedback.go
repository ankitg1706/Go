package models

type Feedback struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Message  string `json:"message"`
}
