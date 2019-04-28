package e2e_test

import (
	"os"
	"testing"
)

type devnullLogger struct{}

func (_ devnullLogger) Logf(string, ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func TestMain(m *testing.M) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	os.Exit(m.Run())
}
