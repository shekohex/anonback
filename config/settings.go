package config

import (
	"os"
)

// Settings The application settings.
type Settings struct {
	title      string
	adminName  string
	passPhrase string
	userInfo   *UserInformation
}

// UserInformation The UserI nformation.
type UserInformation struct {
	username   string
	name       string
	profilePic string
	headerPic  string
	facebook   string
	twitter    string
	github     string
}

// Create create new settings...
func (settings *Settings) Create() *Settings {
	settings.title = "Anonback - New Application"
	settings.adminName = getEnv("ADMIN_NAME", "admin")
	settings.passPhrase = getEnv("PASS_PHRASE", "123456")
	settings.userInfo = &UserInformation{
		username:   getEnv("USERNAME", "username"),
		name:       getEnv("USERNAME", "username"),
		profilePic: getEnv("PROFILE_PIC", "unknown"),
		headerPic:  getEnv("HEADER_PIC", "unknown"),
		facebook:   getEnv("FACEBOOK", "unknown"),
		twitter:    getEnv("TWITTER", "unknown"),
		github:     getEnv("GITHUB", "unknown"),
	}
	return settings
}
func getEnv(key string, defaultValue string) string {
	value, OK := os.LookupEnv(key)
	if OK {
		return value
	}
	return defaultValue
}
