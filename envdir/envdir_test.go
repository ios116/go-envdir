package envdir

import "testing"

func TestEnvdir(t *testing.T) {

	b, err := RunWithEnv("./test_var", "env")
	if err != nil {
		t.Fatal(err)

	}
	t.Log(string(b))
}
