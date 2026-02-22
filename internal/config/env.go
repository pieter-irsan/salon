package config

type Env string

const (
	EnvProduction  Env = "production"
	EnvDevelopment Env = "development"
)

func (e Env) IsProduction() bool {
	return e == EnvProduction
}
