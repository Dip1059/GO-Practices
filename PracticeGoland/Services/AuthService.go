package Services

import (
	G "PracticeGoland/Globals"
	H "PracticeGoland/Helpers"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"log"
	"net/http"
)

var(
	templateData G.UserDataForEmail
	emailData G.EmailGenerals
)


func SetRememberToken(c *gin.Context,sc *securecookie.SecureCookie) {
	val := G.User.Email
	encoded, _ := sc.Encode("remember_token", val)

	cookie1 := http.Cookie{
		Name:     "remember_token",
		Value:    encoded,
		MaxAge:   60 * 60 * 24 * 365,
	}

	http.SetCookie(c.Writer, &cookie1)
}


func SendVerificationEmail() bool {
	encEmail := base64.URLEncoding.EncodeToString([]byte(G.User.Email))
	templateData.EncEmail = encEmail
	templateData.User = G.User
	htmlString, err := H.ParseTemplate("View/Email/email-verify.html", templateData)
	if err != nil {
		log.Println("AuthService.go Log1", err.Error())
		return false
	}

	emailData.From = "Gophers <gopher@mail.com>"
	emailData.To = G.User.Email
	emailData.Subject = "Account Verification Email"
	emailData.HtmlString = htmlString
	if !SendEmail(emailData.From, emailData.To, emailData.Subject, emailData.HtmlString) {
		G.Msg.Fail = "Verification Email Not Sent, <a href='http://localhost:2000/resend-email-verification'>Click Here To Resend</a>."
		return false
	}
	return true
}


