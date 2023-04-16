package main

import (
	dns "dns/main/domainProcess"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// Type data definitions
type AddrData struct {
	IPaddresses []string
	HostName    string
}

type DomainData struct {
	IPaddress   string
	DomainNames []string
}

// Data storage variables implementations
var AddressData = make([]AddrData, 0)
var DomainNamesData = make([]DomainData, 0)

// Handle functions implementations
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving %s for %s\n", r.Host, r.URL.Path)
	myTemplate := template.Must(template.ParseGlob("homePage.html"))
	myTemplate.ExecuteTemplate(w, "homePage.html", nil)
}

func getIP(w http.ResponseWriter, r *http.Request) {
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
		myTemplate.Execute(w, AddressData)
	}
}

func getName(w http.ResponseWriter, r *http.Request) {
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
		myTemplate.Execute(w, DomainNamesData)
	}
}

func getNameServers(w http.ResponseWriter, r *http.Request) {
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
		myTemplate.Execute(w, DomainNamesData)
	}
}

func main() {
	PORT := ":8080"
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Using a default PORT ", PORT)
	} else {
		PORT = ":" + arguments[1]
	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/getIP", getIP)
	http.HandleFunc("/getName", getName)
	http.HandleFunc("/nameServers", getNameServers)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

}
