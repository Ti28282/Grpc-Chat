package config

import "time"

type GRPCConfig struct {
	Port int           `yaml:"port"`
	Time time.Duration `yaml:"timeout"`
}

type Config struct {
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	GRPC        GRPCConfig    `yaml:"grpc" `
	TokenTTL    time.Duration `yaml:"tokenttl" env-required:"true"`
}
