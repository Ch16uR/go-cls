package uuid

import (
	_ "fmt"
	"github.com/satori/go.uuid"
)

func CreateUuid() string {

	id := uuid.NewV4().String()
	//	t := id.String()
	return id
}
