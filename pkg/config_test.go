package pkg

import (
	"testing"

	"github.com/spf13/viper"
)

func TestInitConfig(t *testing.T) {
	viper.Reset()

	err := InitConfig("../config/config.test.yml")
	if err != nil {
		t.Fatalf("failed to load test config: %v", err)
	}

	requiredFields := []string{
		"db.host",
		"db.port",
		"db.username",
		"db.dbname",
		"db.sslmode",
		"port",
	}

	for _, field := range requiredFields {
		if !viper.IsSet(field) {
			t.Errorf("config field %q is not set", field)
		}
	}
}
