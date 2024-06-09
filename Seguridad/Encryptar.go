package seguridad

import (
	"BOLUDO.NET/config"
	"github.com/ProtonMail/gopenpgp/v2/helper"
)

func Encryptar(Cadena string) (string, error) {
	EncrypCadena, err := helper.EncryptMessageWithPassword(config.KEY, Cadena)
	if err != nil {
		return "", err
	} else {
		return EncrypCadena, nil
	}
}

func Desencryptar(Cadena string) (string, error) {
	DecrypCadena, err := helper.DecryptMessageWithPassword(config.KEY, Cadena)
	if err != nil {
		return "", err
	} else {
		return DecrypCadena, nil
	}
}
