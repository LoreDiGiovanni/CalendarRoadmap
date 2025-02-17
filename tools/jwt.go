package tools
import(
	"os"
	"fmt"
	"RoadmapCalendar/types"
	jwt "github.com/golang-jwt/jwt/v5"
)

func CreateUserJWT(user *types.User) (string, error){
    claims:= &jwt.MapClaims{
        "uuid":   user.UUID,
        "email": user.Email,
    }
    secret:= []byte(os.Getenv("JWT_SECRET"))
    token:= jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
    return token.SignedString(secret)
} 

func ValidateJWT(tokenString string) (*jwt.Token,error){
    secret := os.Getenv("JWT_SECRET")
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		    return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	    }
	    return []byte(secret), nil
    })
}

func DecodeJWT(tokenString string) (*jwt.Token,error){
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil})
}

func GetIdFromToken(token *jwt.Token) string{ 
    claims := token.Claims.(jwt.MapClaims)
    useruuid := claims["uuid"].(string)
    return useruuid
}
