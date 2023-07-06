package tests

import (
	"main/bootstrap"
	"testing"
)

func TestMain(t *testing.T) {
	root := ".."
	bootstrap.SetupApp(root)
}
