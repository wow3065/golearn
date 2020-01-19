package testuuid

import (
	"github.com/google/uuid"
)

func GetUUID() string {
	var normal string = uuid.New().String()
	return normal
}

func LdkAdd(a int, b int) int {
	return a + b
}
