// Back-end in Go server
// @jeffotoni
// 2019-01-09

package gmail

import (
	"fmt"
	"os"
)

// GMAIL
var (
	GmailUser     = os.Getenv("GMAIL_USER")
	GmailPassword = os.Getenv("GMAIL_PASSWORD")
	EmailNotify   = os.Getenv("EMAIL_NOTIFIY")

	SubjectNotify = "Notification Gologs!"
	Project       = "Gologs"
	Message       = "Send notification of gologs here when the error is critical"
)

func init() {

	if len(os.Getenv("GMAIL_USER")) <= 0 ||
		len(os.Getenv("GMAIL_PASSWORD")) <= 0 ||
		len(os.Getenv("EMAIL_NOTIFIY")) <= 0 {
		fmt.Printf("\033[0;33m")
		println("If you wanted to be notified via email, ")
		println("you can configure your username and password ")
		println("for sending emails using smtp.")
		fmt.Printf("\033[0;0m")
		showEnvSmtp()
		return
	}
}

func showEnvSmtp() {
	println("Please set the environment variable for sent emails")
	fmt.Printf("\033[0;33m")
	println("   Info:")
	println("    - GMAIL_USER=yout-user")
	println("    - GMAIL_PASSWORD=xxxxxxx")
	println("    - EMAIL_NOTIFIY=user@gmail.com")
	println("\033[0m")
}
