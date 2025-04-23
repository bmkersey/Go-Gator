package config

type Config struct {
	DatabaseURL string `json:"db_url"`
	CurrentUser string `json:"current_user_name"`
}
