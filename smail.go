package smail

import(
	"net/smtp"
	"strings"
	"errors"
)


type Smail struct {
	username_ string
	auth_ smtp.Auth
	fullServerAddr_ string
}


/*
	Makes a new Smail object with the given
*/
func NewSmail(server, port, username, password string) (*Smail, error) {
	//make sure parameters given were not empty
	if server == "" { return nil, errors.New("No server specified") }
	if port == "" { return nil, errors.New("No port specified") }
	if username == "" { return nil, errors.New("No username specified") }
	if password == "" { return nil, errors.New("No password specified") }

	//make a PlainAuth object
	auth := smtp.PlainAuth("", username, password, server)

	//make a smail object and return a pointer to it
	return &Smail{username, auth, server + ":" + port}, nil
}

/*
	Takes a string of comma-seperated email addresses (with or without spaces), and splits it into a
	slice of email adresses
*/
func ToAddrSlice(addresses string) []string {
	//remove all spaces from the string, and then split the string on all commas
	return strings.Split(strings.Replace(addresses, " ", "", -1), ",")
}

/*
	Takes a slice of email adresses, and combines them into a comma seperated list of email adresses, 
	with spaces trailing each comma
*/
func ToAddrList(addresses []string) string {
	var ret = ""
	//loop through all but the last element in the address slice, appending them to the ret string with a comma afterwards
	for _, v := range addresses[:len(addresses)-1] {
		ret += v + ", "
	}
	//append the last address to the ret string & return the list of addresses
	ret += addresses[len(addresses)-1]
	return ret
}


/*
	Sends a plaintext email to the specified recipient email adress(es), with the given subject and body.
	NOTE: it is assumed that the given email adresses are valid.
*/
func (s *Smail) SendPlaintextEmail(recipients []string, subject string, body string) error {

	emailBody := "To: "+ ToAddrList(recipients) + "\r\nSubject: " + subject + "\r\n\r\n" + body

	err := smtp.SendMail(s.fullServerAddr_, s.auth_, s.username_, recipients, []byte(emailBody))
	
	return err
}

/*
	Sends an HTML email to the specified recipient email adress(es), with the given subject and body.
	NOTE: it is assumend that the given email adresses ware valid.
*/
func (s *Smail) SendHTMLEmail(recipients []string, subject string, body string) error {

	mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n";

	emailBody := "To: "+ ToAddrList(recipients) + "\r\nSubject: " + subject + "\r\n" + mime + body
	
	err := smtp.SendMail(s.fullServerAddr_, s.auth_, s.username_, recipients, []byte(emailBody))
	
	return err
}
