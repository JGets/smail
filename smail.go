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
	Makes a new Smail object with the given server, port, account username and password.
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
	Sends a plaintext email to the specified recipient email adress(es), with the given subject and body.
	NOTE: it is assumed that the given email adresses are valid.
*/
func (s *Smail) SendPlaintextEmail(recipients *AddrList, subject string, body string) error {

	if recipients.Empty() {
		return errors.New("Empty AddrList given")
	}

	emailBody := "To: "+ recipients.csv + "\r\nSubject: " + subject + "\r\n\r\n" + body

	err := smtp.SendMail(s.fullServerAddr_, s.auth_, s.username_, recipients.slice, []byte(emailBody))
	
	return err
}

/*
	Sends an HTML email to the specified recipient email adress(es), with the given subject and body.
	NOTE: it is assumend that the given email adresses are valid.
*/
func (s *Smail) SendHTMLEmail(recipients *AddrList, subject string, body string) error {

	if recipients.Empty() {
		return errors.New("Empty AddrList given")
	}

	mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n";

	emailBody := "To: "+ recipients.csv + "\r\nSubject: " + subject + "\r\n" + mime + body
	
	err := smtp.SendMail(s.fullServerAddr_, s.auth_, s.username_, recipients.slice, []byte(emailBody))
	
	return err
}

/*
	Takes a string of comma-seperated email addresses (with or without spaces), and splits it into a
	slice of email adresses
*/
func toAddrSlice(addresses string) []string {
	if addresses == "" {
		return make([]string, 0, 0)
	}

	//remove all spaces from the string, and then split the string on all commas
	return strings.Split(strings.Replace(addresses, " ", "", -1), ",")
}

/*
	Takes a slice of email adresses, and combines them into a comma seperated list of email adresses, 
	with spaces trailing each comma
*/
func toAddrString(addresses []string) string {

	if len(addresses) < 1 {
		return ""
	}

	var ret = ""
	//loop through all but the last element in the address slice, appending them to the ret string with a comma afterwards
	for _, v := range addresses[:len(addresses)-1] {
		ret += v + ", "
	}
	//append the last address to the ret string & return the list of addresses
	ret += addresses[len(addresses)-1]
	return ret
}





type AddrList struct {
	slice []string
	csv string
}

/*
	Makes an empty AddrList
*/
func NewAddrList() *AddrList {
	//return a reference to a new AddrList with an empty slice (with capacity 1), and an empty string
	return &AddrList{make([]string, 0, 1), ""}
}

/*
	Makes an AddrList containing the comma-seperated email addresses in the given string
*/
func NewAddrListFromString(list string) *AddrList {
	//If we're given an empty string, just return an empty AddrList
	if list == "" {
		return NewAddrList()
	}

	return &AddrList{toAddrSlice(list), list}
}

/*
	Makes an AddrList containing the email addresses in the given slice
*/
func NewAddrListFromSlice(slice []string) *AddrList {
	//if we're given an empty slice, just return an empty AddrList
	if len(slice) == 0 {
		return NewAddrList()
	}

	return &AddrList{slice, toAddrString(slice)}
}

/*
	Adds a single address to the AddrList
*/
func (al *AddrList) AddAddress(address string) {
	//append the address to the slice
	al.slice = append(al.slice, address)

	//if there is a previous address, append a comma and a space
	if al.csv != "" {
		al.csv += ", "
	}
	// then append the new address to the comma-seperated string
	al.csv += address
}

/*
	Adds multiple addresses to the AddrList
*/
func (al *AddrList) AddAddresses(addresses []string) {
	//append all the addresses to the slice
	al.slice = append(al.slice, addresses...)

	//re-build the comma-seperated string of addresses
	al.csv = toAddrString(al.slice)

}


/*
	Removes a single address from the AddrList. If the address is not in the AddrList, does nothing.
	Worst Case: O(n), where n = # of addresses already in the list
*/
func (al *AddrList) RemoveAddress(address string) {

	al.slice = removeStringFromSlice(address, al.slice)

	//re-build the comma-seperated string of addresses
	al.csv = toAddrString(al.slice)
	
}

/*
	Removes multiple addresses from the AddrList. If not all addresses given exist in the AddrList, 
	it will only remove those that do.
	Worst Case: O(kn), where k = # of addresses given to remove, n = # of adresses in the list
*/
func (al *AddrList) RemoveAddresses(addresses []string) {

	//go through each given address to remove, and attempt to remove them from the AddrList's internal slice
	for _, v := range addresses {
		al.slice = removeStringFromSlice(v, al.slice)
	}

	//re-build the comma-seperated string of addresses
	al.csv = toAddrString(al.slice)

}

/*
	Removes the first instance of the given string from the given slice. If not instance of the 
	string extist in the slice, the original slice given is returned
	Note: This function is case sensetive
	Worst Case: O(n), where n = # of elements in the slice
*/
func removeStringFromSlice(str string, slice []string) []string {
	for i, v := range slice {
		if v == str {
			//append the subslice of all elements after this one, to the sublice of all elements before this one
			return append(slice[:i], slice[i+1:]...)
		}
	}

	//if the string was not present, just return the slice back
	return slice
}

/*
	Predicate function to check if an AddrList is empty (ie. it contains to addresses)
*/
func (al *AddrList) Empty() bool {
	return len(al.slice) == 0 && al.csv == ""
}

/*
	Get a string representation of the AddrList
*/
func (al *AddrList) String() string {
	return al.csv
}

