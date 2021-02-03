package main

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	resultado := greeting("teste")
	if resultado != "<b>teste</b>" {
		t.Errorf("Texto esperado: %s, obtido: %s", "<b>teste</b>", resultado)
	}
}
