package core

import "github.com/ProtonMail/gopenpgp/v2/helper"

// Desencryptar desencripta una cadena cifrada utilizando una contraseña.
// Devuelve la cadena desencriptada o un error si ocurre uno.
func Desencryptar(PASSWORD, CADENA string) (string, error) {
	// Convertir la contraseña a bytes
	PASSWORD_ := []byte(PASSWORD)

	// Desencriptar la cadena utilizando la contraseña
	CADENA, err := helper.DecryptMessageWithPassword(PASSWORD_, CADENA)
	if err != nil {
		return "", err // Devolver el error si ocurre
	}

	// Devolver la cadena desencriptada
	return CADENA, nil
}
