// Back-end in Go server
// @jeffotoni
// 2019-01-09

package gmail

import "os"

// GMAIL
var (
	GmailUser     = os.Getenv("GMAIL_USER")
	GmailPassword = os.Getenv("GMAIL_PASSWORD")
	EmailNotify   = os.Getenv("EMAIL_NOTIFIY")

	SubjectNotify = "Notification Gologs!"
	Project       = "Gologs"
	Message       = "Send notification of gologs here when the error is critical"
)
