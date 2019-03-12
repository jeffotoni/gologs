// Back-end in Go server
// @jeffotoni
// 2019-01-09

package email

import (
    "fmt"
    "net/mail"
    "net/smtp"
    "strings"
)

var (
    host     = "localhost"
    port     = "25"
    username = ""
    password = ""
    auth     = smtp.PlainAuth("", username, password, host)
    addr     = host + ":" + port
)

///////////
func SendPostifix(fromName, fromEmail string, toNames []string, toEmails []string, subject, body string) {

    // Build RFC-2822 email
    toAddresses := []string{}
    for i := range toEmails {
        to := mail.Address{toNames[i], toEmails[i]}
        toAddresses = append(toAddresses, to.String())
    }
    toHeader := strings.Join(toAddresses, ", ")
    from := mail.Address{fromName, fromEmail}
    fromHeader := from.String()
    subjectHeader := subject
    header := make(map[string]string)
    header["To"] = toHeader
    header["From"] = fromHeader
    header["Subject"] = subjectHeader
    header["Content-Type"] = `text/html; charset="UTF-8"`
    msg := ""
    for k, v := range header {
        msg += fmt.Sprintf("%s: %s\r\n", k, v)
    }
    msg += "\r\n" + body
    bMsg := []byte(msg)
    // Send using local postfix service
    c, err := smtp.Dial(addr)
    if err != nil {
        return
    }
    defer c.Close()
    if err = c.Mail(fromHeader); err != nil {
        return
    }

    ///////
    for _, addr := range toEmails {
        if err = c.Rcpt(addr); err != nil {
            return
        }
    }
    w, err := c.Data()
    if err != nil {
        return
    }
    _, err = w.Write(bMsg)
    if err != nil {
        return
    }
    err = w.Close()
    if err != nil {
        return
    }
    err = c.Quit()

    if err != nil {
        return
    }
}
