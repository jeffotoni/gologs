// Back-end in Go server
// @jeffotoni
// 2019-01-09

package gmail

import (
	"bytes"
	"fmt"
	"log"
	"mime/quotedprintable"
	"net/smtp"
	"strings"
)

const (
	SMTPSERVER = "smtp.gmail.com"
)

type SenderGmail struct {
	User     string
	Password string
}

func NewSenderGmail(Username, Password string) SenderGmail {

	return SenderGmail{Username, Password}
}

func (SenderGmail SenderGmail) WriteGmail(dest []string, contentType, subject, bodyMessage string) string {

	header := make(map[string]string)
	header["From"] = SenderGmail.User

	receipient := ""

	for _, user := range dest {
		receipient = receipient + user
	}

	header["To"] = receipient
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", contentType)
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	message := ""

	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	var encodedMessage bytes.Buffer

	finalMessage := quotedprintable.NewWriter(&encodedMessage)
	finalMessage.Write([]byte(bodyMessage))
	finalMessage.Close()

	message += "\r\n" + encodedMessage.String()

	return message
}

func (SenderGmail *SenderGmail) WriteHTML(dest []string, subject, bodyMessage string) string {

	return SenderGmail.WriteGmail(dest, "text/html", subject, bodyMessage)
}

func (SenderGmail *SenderGmail) WritePlainEmail(dest []string, subject, bodyMessage string) string {

	return SenderGmail.WriteGmail(dest, "text/plain", subject, bodyMessage)
}

func (SenderGmail SenderGmail) SendMail(Dest []string, Subject, bodyMessage string) bool {

	msg := "From: " + SenderGmail.User + "\n" +
		"To: " + strings.Join(Dest, ",") + "\n" +
		"Subject: " + Subject + "\n" + bodyMessage

	err := smtp.SendMail(SMTPSERVER+":587",
		smtp.PlainAuth("", SenderGmail.User, SenderGmail.Password, SMTPSERVER),
		SenderGmail.User, Dest, []byte(msg))

	if err != nil {

		log.Printf("smtp error: %s", err)
		return false
	}

	return true
}
