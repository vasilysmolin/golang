package main

import (
 "main/tests"
 "testing"
 "main/bootstrap"
)

func TestMain(t *testing.T,) {
      app := bootstrap.SetupApp()
      tests.TestRegister(t, app)
}
