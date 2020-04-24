package main

import (
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func TestCompatibility(t *testing.T) {

	rand.Seed(time.Now().Unix())
	isCompatibility := rand.Intn(2)

	if runtime.Version() == "go1.12.17" && isCompatibility == 0 {
		t.Errorf("Version: [%s] incompatible", runtime.Version())
	} else {
		t.Logf("Version: [%s] compatible", runtime.Version())
	}
}
