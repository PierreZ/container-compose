package env

import (
	"os"
	"fmt"
	"strings"
	"github.com/pkg/errors"
)

const envFilesToTemplate string = "CONTAINER_COMPOSE_TEMPLATES"
const prefixEnv string = "CONTAINER_COMPOSE"
const listSeparator string = ","

// GetFilesToTemplates is reading env vars to get files that 
// need to be templated
func GetFilesToTemplates() ([]string, error) {

	return readListFromEnv(envFilesToTemplate)
}

func readListFromEnv(key string) ([]string, error) {

	env := os.Getenv(key)

	if len(env) == 0 {
		return nil, errors.Wrap(fmt.Errorf("no env for '%s'", key), "read from environnement failed")
	}

	envs := strings.Split(env, listSeparator)

	return envs, nil
}