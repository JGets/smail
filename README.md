#Smail - Simple Mail
````
import "github.com/JGets/smail"
````

##About
Smail is meant to be a small package designed to simplify sending emails in Go.

##Usage
......

##Documentation
###Variables
````
var(
	ErrNoServer = errors.New("No server specified")
	ErrNoPort = errors.New("No port specified")
	ErrNoUsername = errors.New("No username specified")
	ErrNoPassword = errors.New("No password specified")
)
````

###Types
####Smail
````
type Smail struct {
	//contains unexported fields
}
````

####func NewSmail
````
func NewSmail(server, port, username, password string) (*Smail, error)
````
Makes a new `Smail` object for the given smtp `server` address, using `port`, for the given `username` and `password`.


####func (*Smail) SendPlaintextEmail
````
func (s *Smail) SendPlaintextEmail(recipients *AddrList, subject string, body string) error
````
Sends a plaintext email to the specified recipient email adress(es), with the given subject and body.
**NOTE:** it is assumed that the given email adresses are valid.


####func (*Smail) SendHTMLEmail
````
func (s *Smail) SendHTMLEmail(recipients *AddrList, subject string, body string) error
````
Sends an HTML email to the specified recipient email adress(es), with the given subject and body.
**NOTE:** it is assumend that the given email adresses are valid.


####AddrList
````
type AddrList struct {
	//contains unexported fields
}
````

####func NewAddreList
````
func NewAddrList() *AddrList
````
Makes an empty AddrList.


####func NewAddrListFromString
````
func NewAddrListFromString(list string) *AddrList
````
Makes an AddrList containing email addresses in the given string. Email addresses must be comma seperated, and any spaces in the string are ignored.


####func NewAddrListFromSlice
````
func NewAddrListFromSlice(slice []string) *AddrList
````
Makes an AddrList containing the email addresses in the given `slice`.


####func (*AddrList) AddAddress
````
func (al *AddrList) AddAddress(address string)
````
Adds a single address to the AddrList `al`.


####func (*AddrList) AddAddresses
````
func (al *AddrList) AddAddresses(addresses []string)
````
Adds all the email addresses in `addresses` to the AddrList `al`.


####func (*AddrList) RemoveAddress
````
func (al *AddrList) RemoveAddress(address string)
````
Removes a single address from the AddrList. If the address is not in the AddrList, does nothing.
**Complexity:** Worst Case: O(n), where n = # of addresses already in the list.


####func (*AddrList) RemoveAddresses
````
func (al *AddrList) RemoveAddresses(addresses []string)
````
Removes multiple addresses from the AddrList. If not all addresses given exist in the AddrList, it will only remove those that do.
**Complexity:** Worst Case: O(kn), where k = # of addresses given to remove, n = # of adresses in the list.


####func (*AddrList) Empty
````
func (al *AddrList) Empty() bool
````
Predicate function to check if an AddrList is empty (ie. it contains to addresses).


####func (*AddrList) String
````
func (al *AddrList) String() string
````
Get a string representation of the AddrList.






