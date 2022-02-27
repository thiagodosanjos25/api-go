package api

import (
	"math"
	"testing"
)

func TestValidadeFields(t *testing.T) {

	c := &Client{
		Name:   "",
		Weight: 0,
		Height: 0,
		Gender: "",
	}

	err := validateFields(c)
	if err == nil {
		t.Error("Error inesperado. Inputs vázios.")
	}

	c = &Client{
		Name:   "Name test",
		Weight: 105,
		Height: 1.77,
		Gender: "M",
	}

	err = validateFields(c)
	if err != nil {
		t.Error("Error inesperado. Inputs preenchidos.")
	}
}

func TestGenerateIMCandSituation(t *testing.T) {

	c := &Client{
		Weight: 105,
		Height: 1.77,
	}

	imcExpected := 33.52
	situationExpected := "Obesidade grau 1"

	imc, situation := generateIMCandSituation(c.Weight, c.Height)

	if imc = toFixed(imc, 2); imc != imcExpected {
		t.Errorf("Error inesperado - IMC. Situação calculada diferente do esperado. Obtido: %2.f vs esperado: %2.f", imc, imcExpected)
	}

	if situation != situationExpected {
		t.Errorf("Error inesperado - Situation. Situação calculada diferente do esperado. Obtido: %s vs esperado: %s", situation, situationExpected)
	}

	c = &Client{
		Weight: 90,
		Height: 1.77,
	}

	imcExpected = 28.73
	situationExpected = "Sobrepeso"

	imc, situation = generateIMCandSituation(c.Weight, c.Height)

	if imc = toFixed(imc, 2); imc != imcExpected {
		t.Errorf("Error inesperado - IMC. Situação calculada diferente do esperado. Obtido: %2.f vs esperado: %2.f", imc, imcExpected)
	}

	if situation != situationExpected {
		t.Errorf("Error inesperado - Situation. Situação calculada diferente do esperado. Obtido: %s vs esperado: %s", situation, situationExpected)
	}
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
