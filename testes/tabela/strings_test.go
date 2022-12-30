package strings

const msgIndex = "%s (parte: %s) - indices: esperando (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Cod3r eh show", "Cod3r", 0},
		{"", "", 0},
		{"Opa", "opoa", -1},
		{"leopnardo", "o", 2},
	}
}

for _, teste := range testes {
	t.Logf("Massa: %v", teste)
	atual := string.Index(teste.texto, teste.parte)

	if atual != teste.esperado {
		t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado)
	}
}
