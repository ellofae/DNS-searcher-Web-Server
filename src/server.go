package main

import (
	dns "dns/main/domainProcess"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

const (
	DATAFILE1 string = "/tmp/domain-name-server-data.json"
	DATAFILE2 string = "/tmp/mail-server-data.json"
	DATAFILE3 string = "/tmp/domain-name-data.json"
	DATAFILE4 string = "/tmp/ip-address-data.json"
)

// Type data definitions
type AddrData struct {
	HostName    string
	IPaddresses []string
}

type DomainData struct {
	IPaddress   string
	DomainNames []string
}

type NSData struct {
	Domain  string
	Servers []string
}

type MXData struct {
	Domain  string
	Servers []string
}

// Data storage variables implementations
var AddressData = make([]AddrData, 0)
var DomainNamesData = make([]DomainData, 0)
var NameServersData = make([]NSData, 0)
var MailServersData = make([]MXData, 0)

// Saving and loading data
func saveData(DATA string, SomeStruct interface{}) error {
	fmt.Println("Saving", DATA)
	err := os.Remove(DATA)
	if err != nil {
		log.Println(err)
	}

	saveTo, err := os.Create(DATA)
	if err != nil {
		log.Println("Didn't manage to create", DATA)
		return err
	}
	defer saveTo.Close()

	encoder := json.NewEncoder(saveTo)
	err = encoder.Encode(SomeStruct)
	if err != nil {
		log.Println("Didn't manage to save to", DATA)
		return err
	}
	return nil
}

func loadData(DATA string, SomeStruct interface{}) error {
	fmt.Println("Loading", DATA)
	loadFrom, err := os.Open(DATA)
	if err != nil {
		log.Println("No records have been found")
		return err
	}
	defer loadFrom.Close()

	decoder := json.NewDecoder(loadFrom)
	decoder.Decode(&SomeStruct)
	return nil
}

func loadAllData() error {
	var err error

	err = loadData(DATAFILE1, &NameServersData)
	if err != nil {
		fmt.Println(err)
	}

	err = loadData(DATAFILE2, &MailServersData)
	if err != nil {
		fmt.Println(err)
	}

	err = loadData(DATAFILE3, &DomainNamesData)
	if err != nil {
		fmt.Println(err)
	}

	err = loadData(DATAFILE4, &AddressData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nLoading is completed!\n\n")

	return nil
}

// Handle functions implementations
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving %s for %s\n", r.Host, r.URL.Path)
	myTemplate := template.Must(template.ParseGlob("homePage.html"))
	myTemplate.ExecuteTemplate(w, "homePage.html", nil)
}

func GetIP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving %s for %s\n", r.Host, r.URL.Path)
	myTemplate := template.Must(template.ParseFiles("getAddrPage.html"))

	if r.Method != http.MethodPost {
		myTemplate.Execute(w, AddressData)
		return
	}

	hostName := r.FormValue("domainName")

	IPs, err := dns.DomainProcess(hostName, "", 'i')
	if err == nil {
		fmt.Printf("\n*recived IPs: %#v\n\n", IPs)

		addr := &AddrData{IPaddresses: IPs, HostName: hostName}
		AddressData = append(AddressData, *addr)

		err := saveData(DATAFILE4, AddressData)
		if err != nil {
			fmt.Println(err)
			return
		}

		myTemplate.Execute(w, AddressData)
	} else {
		addr := &AddrData{IPaddresses: []string{"no IP addresses found"}, HostName: hostName}
		AddressData = append(AddressData, *addr)
		myTemplate.Execute(w, AddressData)
	}
}

func GetName(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving %s for %s\n", r.Host, r.URL.Path)
	myTemplate := template.Must(template.ParseFiles("getDomainPage.html"))

	if r.Method != http.MethodPost {
		myTemplate.Execute(w, DomainNamesData)
		return
	}

	ip := r.FormValue("ip")

	domains, err := dns.DomainProcess("", ip, 'd')
	if err == nil {
		fmt.Printf("\n*recived Hosts: %#v\n\n", domains)

		hosts := &DomainData{IPaddress: ip, DomainNames: domains}
		DomainNamesData = append(DomainNamesData, *hosts)

		err := saveData(DATAFILE3, DomainNamesData)
		if err != nil {
			fmt.Println(err)
			return
		}

		myTemplate.Execute(w, DomainNamesData)
	} else {
		hosts := &DomainData{IPaddress: ip, DomainNames: []string{"no domains found"}}
		DomainNamesData = append(DomainNamesData, *hosts)
		myTemplate.Execute(w, DomainNamesData)
	}
}

func GetNameServers(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving %s for %s\n", r.Host, r.URL.Path)
	myTemplate := template.Must(template.ParseFiles("getNameServerPage.html"))

	if r.Method != http.MethodPost {
		myTemplate.Execute(w, NameServersData)
		return
	}

	domain := r.FormValue("domainName")

	servers, err := dns.DomainProcess(domain, "", 'n')
	if err == nil {
		fmt.Printf("\n*recived domain name servers: %#v\n\n", servers)

		NSs := &NSData{Domain: domain, Servers: servers}
		NameServersData = append(NameServersData, *NSs)

		err := saveData(DATAFILE1, NameServersData)
		if err != nil {
			fmt.Println(err)
			return
		}

		myTemplate.Execute(w, NameServersData)
	} else {
		temp := &NSData{Domain: domain, Servers: []string{"no servers found"}}
		NameServersData = append(NameServersData, *temp)
		myTemplate.Execute(w, NameServersData)
	}
}

func GetMailServers(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving %s for %s\n", r.Host, r.URL.Path)
	myTemplate := template.Must(template.ParseFiles("getMailServerPage.html"))

	if r.Method != http.MethodPost {
		myTemplate.Execute(w, MailServersData)
		return
	}

	domain := r.FormValue("domainName")

	servers, err := dns.DomainProcess(domain, "", 'm')
	if err == nil {
		fmt.Printf("\n*recived mail servers: %#v\n\n", servers)

		MXs := &MXData{Domain: domain, Servers: servers}
		MailServersData = append(MailServersData, *MXs)

		err := saveData(DATAFILE2, MailServersData)
		if err != nil {
			fmt.Println(err)
			return
		}

		myTemplate.Execute(w, MailServersData)
	} else {
		temp := &MXData{Domain: domain, Servers: []string{"no servers found"}}
		MailServersData = append(MailServersData, *temp)
		myTemplate.Execute(w, MailServersData)
	}

}

func main() {
	err := loadAllData()
	if err != nil {
		fmt.Println(err)
		return
	}

	PORT := ":8080"
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Using a default PORT ", PORT)
	} else {
		PORT = ":" + arguments[1]
	}

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/getIP", GetIP)
	http.HandleFunc("/getName", GetName)
	http.HandleFunc("/nameServers", GetNameServers)
	http.HandleFunc("/mailServers", GetMailServers)

	err = http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
