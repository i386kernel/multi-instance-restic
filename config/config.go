package config

import (
	"log"
	"os"
)

type KubeConfig struct {
	KubePR      string
	KubeDR      string
	KubePRToken string
	KubeDRToken string
}

type AwsConfig struct {
	AwsAccessKeyID     string
	AwsSecretAccessKey string
}

type ResticConfig struct {
	ResticPassword string
}

type Config struct {
	Kube KubeConfig
	AWS  AwsConfig
	Restic ResticConfig
}

func New() *Config {
	return &Config{
		Kube: KubeConfig{
			KubePR:      getEnv("KUBE_PR"),
			KubeDR:      getEnv("KUBE_DR"),
			KubePRToken: getEnv("KUBE_PR_TOKEN"),
			KubeDRToken: getEnv("KUBE_DR_TOKEN"),
		},
		AWS: AwsConfig{
			AwsAccessKeyID:     getEnv("AWS_ACCESS_KEY_ID"),
			AwsSecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY"),
		},
		Restic: ResticConfig{ResticPassword: getEnv("RESTIC_PASSWORD")},
	}
}

func getEnv(key string) string {
	env, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("%s Env is not set, please set appropriate ENV's and rerun the program, exiting program...", key)
	}
	return env
}
