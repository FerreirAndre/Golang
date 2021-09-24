package endereco

import "strings"

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	primeiraPal := strings.Split(endereco, " ")[0]
	primeiraPalMin := strings.ToLower(primeiraPal)

	valido := false

	for _, palavra := range tiposValidos {
		if palavra == primeiraPalMin {
			valido = true
		}
	}

	if valido {
		return strings.Title(primeiraPalMin)
	}
	return "tipo invalido"
}
