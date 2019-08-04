package config

import (
	"fmt"
	"os"

	"github.com/davidbanham/required_env"
)

func main() {
	fmt.Println("vim-go")
}

var RENDER_ERRORS bool
var PORT string

func init() {
	required_env.Ensure(map[string]string{
		"RENDER_ERRORS": "false",
		"PORT":          "3000",
	})

	PORT = os.Getenv("PORT")
	RENDER_ERRORS = os.Getenv("RENDER_ERRORS") == "true"
}
