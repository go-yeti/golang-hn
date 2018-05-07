package configs

type EnvInterface interface {
	GetVars() map[string]string
	SetVars(map[string]string)
}
