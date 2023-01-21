package models
import (
    "fmt"
    "net/http"
)

type Item struct {
    ID int `json:"id"`
    UserName string `json:"username"`
    Password string `json:"password"`
    Email string `json:"email"`
    CreatedAt string `json:"created_at"`
}
