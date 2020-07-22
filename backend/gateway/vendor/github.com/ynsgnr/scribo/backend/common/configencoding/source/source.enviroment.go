package source

import "os"

const EnvTag = "env"

type Environment struct{}

func (Environment) GetValue(tag string) (string, bool) {
	return os.LookupEnv(tag)
}
