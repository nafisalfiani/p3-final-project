package entity

type Recipient struct {
	ToEmails  []string
	CCEmails  []string
	BCCEmails []string
}

type Email struct {
	Body        string
	BodyType    string
	Subject     string
	SenderName  string
	SenderEmail string
	Recipients  Recipient
	Attachments []string
	Headers     map[string]string
}
