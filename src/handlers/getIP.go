package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/ellofae/DNS-searcher-Web-Server/src/domainProcess"

	"github.com/ellofae/DNS-searcher-Web-Server/src/dataProccess"
)

// Handler interface implementation

type GetIP struct {
	l *log.Logger
}

func NewGetIP(l *log.Logger) *GetIP {
	return &GetIP{l}
}

func (g *GetIP) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("IP service getter requested")

	fmt.Printf("Serving %s for %s\n", r.Host, r.URL.Path)
	myTemplate := template.Must(template.ParseFiles("getAddrPage.html"))

	if r.Method != http.MethodPost {
		myTemplate.Execute(rw, dataProccess.AddressData)
		return
	}

	hostName := r.FormValue("domainName")

	IPs, err := domainProcess.DomainProcess(hostName, "", 'i')
	if err == nil {
		fmt.Printf("\n*recived IPs: %#v\n\n", IPs)

		addr := &dataProccess.AddrData{IPaddresses: IPs, HostName: hostName}
		dataProccess.AddressData = append(dataProccess.AddressData, *addr)

		err := dataProccess.SaveData(dataProccess.DATAFILE4, dataProccess.AddressData)
		if err != nil {
			fmt.Println(err)
			return
		}

		myTemplate.Execute(rw, dataProccess.AddressData)
	} else {
		addr := &dataProccess.AddrData{IPaddresses: []string{"no IP addresses found"}, HostName: hostName}
		dataProccess.AddressData = append(dataProccess.AddressData, *addr)
		myTemplate.Execute(rw, dataProccess.AddressData)
	}
}
