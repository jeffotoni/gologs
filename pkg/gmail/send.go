// Back-end in Go server
// @jeffotoni
// 2019-01-09

package gmail

// you can parameterize this function as needed
func Send(to []string, subject, subjectTmpl, HtmlTmpl string) bool {
	gmail := NewSenderGmail(GmailUser, GmailPassword)
	Receiver := to
	Subject := subject
	message := TmplDefault(subjectTmpl, HtmlTmpl)
	bodyMessage := gmail.WriteHTML(Receiver, Subject, message)
	return gmail.SendMail(Receiver, Subject, bodyMessage)
}
