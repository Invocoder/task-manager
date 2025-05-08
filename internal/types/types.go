package types
import "time"

// type Task struct {
//     ID        int       `json:"id"`
//     Title     string    `json:"title"`
//     Status    string    `json:"status"` 
//     CreatedAt time.Time `json:"created_at"`
// }

type Task struct {
    ID        int       
    Title     string `validate:"required"` 
    Status    string  `validate:"required"` 
    CreatedAt time.Time `validate:"required"` 
}