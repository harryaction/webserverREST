package tools

import guuid "github.com/google/uuid"

func GenUUID() *string {
	id, _ := guuid.NewRandom()
	uuid := id.String()
	return &uuid
}
