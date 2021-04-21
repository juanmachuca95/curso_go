package main 

/*

TEST: http://localhost:3000/login

JSON ENVIADO POR POST - USANDO POSTMAN
{
	"email":"my@gmail.com",
	"password":"password123"
}


RESULTADO ESPERADO:
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIxLTA0LTI3VDIxOjExOjU1Ljc0NzA1NjQtMDM6MDAiLCJzdWIiOiIxIn0.um2k52Pn_S_bTNtXlm67SUWqAgwFzW862Yssij-IgYU",
    "user": {
        "id": 1,
        "email": "my@gmail.com",
        "name": "JuanTest"
    }
}

*/


import (
	"time"
	"github.com/gofiber/fiber"
	"github.com/dgrijalva/jwt-go"
)
 
const jwtSecret = "asecret"

func main(){
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx){
		ctx.Send("Middleware example");
	})

	app.Post("/login", login)

	err := app.Listen( 3000 )
	if err != nil {
		panic(err)
	}
}

func login (ctx *fiber.Ctx){
	type request struct {
		Email string `json:email`
		Password string `json:password`
	}

	var body request
	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Cannot parse json",
		})
		return
	}

	if body.Email != "my@gmail.com" || body.Password != "password123" {
		ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":"Bad credentials",
		})
		return 
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "1"
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7) // a week

	s, err := token.SignedString([]byte( jwtSecret ))

	if err != nil {
		ctx.SendStatus(fiber.StatusInternalServerError)
		return
	}

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": s,
		"user" : struct {
			Id int `json:"id"`
			Email string `json:"email"`
			Name string `json:"name"`
		}{
			Id: 1,
			Email: "my@gmail.com",
			Name: "JuanTest",
		},
 	})

}