package models

// Player this is Player
type Player struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetPlayer this is GetPlayer
func GetPlayer(username, password string) Player {
	var player Player
	db.Select("id,username,password,created_at,updated_at").Where(Player{Username: username, Password: password}).First(&player)
	return player
}
