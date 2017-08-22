package server

import "testing"

func TestULIDGenerateToken(t *testing.T) {
	n := 1000
	results := make([]string, n)
	for i := 0; i < n; i++ {
		v := GenerateToken()
		for j := 0; j < i; j++ {
			if results[j] == v {
				t.Errorf("duplicate token generated %s", v)
				t.FailNow()
			}
		}
		results[i] = v
	}
}
