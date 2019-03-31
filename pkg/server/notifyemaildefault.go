// Go in Action
// @jeffotoni
// 2019-03-30

// server.go
package server

import (
	"fmt"
	"log"

	"github.com/jeffotoni/gologs/pkg/gmail"
)

// you can parameterize
// this function as needed
func notifyEmailDefault() {

	to := []string{gmail.EmailNotify}
	subject := gmail.SubjectNotify
	project := gmail.Project
	message := gmail.Message

	// if the message comes with some critical rules, send emails
	if gmail.Send(to, subject, project, message) {
		log.Println("Mail sent successfully!")
	} else {
		fmt.Println("error sending email!")
	}
}
