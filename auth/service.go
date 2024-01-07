package auth

import "github.com/dgrijalva/jwt-go"

// Generate JWT token
type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

// Temporary secret key
var SECRET_KEY = []byte("CROWDFUNDING_SECRET_KEY")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// Verify signature
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
