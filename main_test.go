package envstringmap_test

import (
	"os"
	"testing"

	"github.com/fabiante/envstringmap"
)

func Test(t *testing.T) {
	var setupValues = map[string]string{
		"APP_VALUE_ABC_1": "1",
		"APP_VALUE_ABC_2": "2",
		"APP_ELSE":        "unrelated",
	}

	// Cleanup the environment after the test.
	t.Cleanup(func() {
		for k := range setupValues {
			os.Unsetenv(k)
		}
	})

	// Set env variables before the test
	for k, v := range setupValues {
		os.Setenv(k, v)
	}

	// Actually test
	values := envstringmap.GetMap("APP_VALUE")
	if len(values) != 2 {
		t.Errorf("Expected 2 values, got %d", len(values))
	}

	t.Logf("values: %+v", values)

	if abc1, ok := values["ABC_1"]; !ok {
		t.Error("Expected key ABC_1 to be present")
	} else if abc1 != "1" {
		t.Errorf("Expected value 1, got %s", abc1)
	}

	if abc2, ok := values["ABC_2"]; !ok {
		t.Error("Expected key ABC_2 to be present")
	} else if abc2 != "2" {
		t.Errorf("Expected value 2, got %s", abc2)
	}
}
