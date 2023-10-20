package types

type BaseDeegoo struct {
	ID         string `json:"id"`
	UserId     string `json:"user_id"`
	Perception int8   `json:"perception"`
	Focus      int8   `json:"focus"`
	Execution  int8   `json:"execution"`
	Memory     int8   `json:"memory"`
	Language   int8   `json:"language"`
}
