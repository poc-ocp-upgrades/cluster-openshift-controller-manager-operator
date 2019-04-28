package framework

import (
	"testing"
)

type Logger interface{ Logf(string, ...interface{}) }

var _ Logger = &testing.T{}
