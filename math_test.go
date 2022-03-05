package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 15)
	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado é %d. Esperado %d", total, 30)
	}
}

func TestGreeting(t *testing.T) {
	greet := greeting("any_name")
	if greet != "Hello any_name" {
		t.Errorf("Greeting result is wrong. Result should be: %s", greet)
	}
}
