package logs

import (
	"log"
)

func InitLogs() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
