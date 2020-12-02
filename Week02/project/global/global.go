package global

import "project/model/database"

var (
	DB *database.Memory
)

func init() {
	database := database.NewMemory()
	DB = database
}
