package helpers

type SIgnedDetails struct {
	Email     string
	FirstName string
	LastName  string
	Uid       string
	User_type string
	jwt.StandardClaims
}

func validateToken(signedToken string) (claims *SIgnedDetails, msg string) {

}
