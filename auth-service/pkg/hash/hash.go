package hash

//helper para el uso de la libreria bycrypt para el hasheo de las contraseñas
import (
	"golang.org/x/crypto/bcrypt"
)

//HashPassword genera un hash de la contraseña
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

//ComparePassword compara la contraseña con su hash
func ComparePassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
