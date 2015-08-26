package encontreOTelefone

import "testing"

func TestLetraA(t *testing.T) {
	numeros := getNumerosTelefone("A")

	if numeros != "2" {
		t.Errorf("Numero deveria ser 2, recebido %s", numeros)
	}
}

func TestLetraD(t *testing.T) {
	numeros := getNumerosTelefone("D")

	if numeros != "3" {
		t.Errorf("Numero deveria ser 3, recebido %s", numeros)
	}
}

func TestLetraF(t *testing.T) {
	numeros := getNumerosTelefone("F")

	if numeros != "3" {
		t.Errorf("Numero deveria ser 3, recebido %s", numeros)
	}
}

func TestLetraN(t *testing.T) {
	numeros := getNumerosTelefone("N")

	if numeros != "6" {
		t.Errorf("Numero deveria ser 6, recebido %s", numeros)
	}
}

func TestDuasLetras(t *testing.T) {
	numeros := getNumerosTelefone("AD")

	if numeros != "23" {
		t.Errorf("Numeros deveriam ser 23, recebido %s", numeros)
	}
}

func TestComHifens(t *testing.T) {
	numeros := getNumerosTelefone("1-HOME-SWEET-HOME")

	if numeros != "1-4663-79338-4663" {
		t.Errorf("Esperado 1-4663-79338-4663, recebido %s", numeros)
	}
}
