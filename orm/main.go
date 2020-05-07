package orm

import (
	"log"
	"os"
)

func init() {
	var log1 *log.Logger

    log1 = log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime)
	log1.Print("Test1\n")
}
