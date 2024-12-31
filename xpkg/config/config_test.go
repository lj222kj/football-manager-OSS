package config

import (
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	t.Run("default", func(t *testing.T) {

		var appName = "football-manager-oss"
		var env = "dev"
		var host = "127.0.0.1"
		var port = "8080"
		var redisHost = "127.0.0.1"
		var redisPort = "6379"
		var redisPassword = ""

		envVars := []string{
			"APP_NAME", "ENV", "HOST", "PORT",
			"REDIS_HOST", "REDIS_PORT", "REDIS_PASSWORD",
		}
		for _, env := range envVars {
			old := os.Getenv(env)
			_ = os.Unsetenv(env)
			t.Cleanup(func() { _ = os.Setenv(env, old) })
		}

		cfg, err := NewConfig()
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if cfg.AppName != appName {
			t.Errorf("AppName default mismatch: got %q, want %q", cfg.AppName, appName)
		}
		if cfg.Env != env {
			t.Errorf("Env default mismatch: got %q, want %q", cfg.Env, env)
		}
		if cfg.Host != host {
			t.Errorf("Host default mismatch: got %q, want %q", cfg.Host, host)
		}
		if cfg.Port != port {
			t.Errorf("Port default mismatch: got %q, want %q", cfg.Port, port)
		}

		if cfg.Redis == nil {
			t.Fatalf("expected Redis to be non-nil")
		}
		if cfg.Redis.Host != redisHost {
			t.Errorf("Redis.Host default mismatch: got %q, want %q", cfg.Redis.Host, redisHost)
		}
		if cfg.Redis.Port != redisPort {
			t.Errorf("Redis.Port default mismatch: got %q, want %q", cfg.Redis.Port, redisPort)
		}
		if cfg.Redis.Password != redisPassword {
			t.Errorf("Redis.Password default mismatch: got %q, want empty string", redisPassword)
		}
	})

	t.Run("set", func(t *testing.T) {
		var appName = "override-default-fb-manager-oss"
		var env = "override-default-dev"
		var host = "override-default-127.0.0.1"
		var port = "override-default-8080"
		var redisHost = "override-default-127.0.0.1"
		var redisPort = "override-default-6379"
		var redisPassword = "override-default-password"

		envSetup := map[string]string{
			"APP_NAME":       appName,
			"ENV":            env,
			"HOST":           host,
			"PORT":           port,
			"REDIS_HOST":     redisHost,
			"REDIS_PORT":     redisPort,
			"REDIS_PASSWORD": redisPassword,
		}

		for k, v := range envSetup {
			old := os.Getenv(k)
			_ = os.Setenv(k, v)
			t.Cleanup(func() { _ = os.Setenv(k, old) })
		}

		cfg, err := NewConfig()
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		// Verify overrides
		if cfg.AppName != appName {
			t.Errorf("AppName mismatch: got %q, want %q", cfg.AppName, appName)
		}
		if cfg.Env != env {
			t.Errorf("Env mismatch: got %q, want %q", cfg.Env, env)
		}
		if cfg.Host != host {
			t.Errorf("Host mismatch: got %q, want %q", cfg.Host, host)
		}
		if cfg.Port != port {
			t.Errorf("Port mismatch: got %q, want %q", cfg.Port, port)
		}

		if cfg.Redis == nil {
			t.Fatalf("expected Redis to be defined")
		}
		if cfg.Redis.Host != redisHost {
			t.Errorf("Redis.Host mismatch: got %q, want %q", cfg.Redis.Host, redisHost)
		}
		if cfg.Redis.Port != redisPort {
			t.Errorf("Redis.Port mismatch: got %q, want %q", cfg.Redis.Port, redisPort)
		}
		if cfg.Redis.Password != redisPassword {
			t.Errorf("Redis.Password mismatch: got %q, want %q", cfg.Redis.Password, redisPassword)
		}
	})
}
