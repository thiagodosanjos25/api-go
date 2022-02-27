package api

import (
	"database/sql"
	"fmt"
	"math"
	"strconv"

	"github.org/api-go/internal/database"
)

// rowNil verifica se a coluna do select veio nula
func rowNil(r database.Row, column int) string {
	if r[column] != nil {
		return r.String(column)
	}
	return ""
}

//NewNullString ...
func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func encrypt(senha string) string {

	resultado := ""
	saida := ""
	byte := ""
	for i := 0; i < len(senha); i++ {
		byte = ""
		resultado = bin(int(senha[i]))
		auxResult, _ := strconv.Atoi(resultado)
		resultado = "0b" + fmt.Sprintf("%08d", auxResult)
		for index, item := range resultado {
			if index > 1 {
				if item == '0' {
					byte += "1"
				} else if item == '1' {
					byte += "0"
				} else {
					byte += fmt.Sprintf("%c", item)
				}
			}
		}
		byteAux, _ := strconv.Atoi(byte)
		saida += string(rune(dec(byteAux)))
	}
	return saida
}

// bin - Converte Decimal para Binary
func bin(numero int) string {
	return strconv.FormatInt(int64(numero), 2)
}

// dec - Converte Binary para Decimal
func dec(number int) int {

	decimal := 0
	counter := 0.0
	remainder := 0

	for number != 0 {
		remainder = number % 10
		decimal += remainder * int(math.Pow(2.0, counter))
		number = number / 10
		counter++
	}
	return decimal
}

// StrSizeCenter centralizar texto dado tamanho de acordo com tamanho do campo
func StrSizeCenter(valor string, tamanhoCampo int) string {
	var valorStr = valor

	if len(valorStr) == tamanhoCampo {
		return valorStr
	}

	var qtdDeZeros = tamanhoCampo - len(valorStr) // tamanho restante para se completar com zeros
	for i := 0; i < qtdDeZeros; i++ {
		if (i & 1) == 0 {
			valorStr = fmt.Sprintf(" %s", valorStr)
		} else {
			valorStr = fmt.Sprintf("%s ", valorStr)
		}
	}

	return fmt.Sprintf("%s", valorStr)
}
