package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/ellofae/DNS-searcher-Web-Server/src/domainProcess"

	"github.com/ellofae/DNS-searcher-Web-Server/src/dataProccess"
)

type GetNameServers struct {
	l *log.Logger
}

func NewGetNameServers(l *log.Logger) *GetNameServers {
	return &GetNameServers{l}
}

func (g *GetNameServers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Name Servers getter requested")

	fmt.Printf("Serving %s for %s\n", r.Host, r.URL.Path)
	myTemplate := template.Must(template.ParseFiles("getNameServerPage.html"))

	if r.Method != http.MethodPost {
		myTemplate.Execute(rw, dataProccess.NameServersData)
		return
	}

	domain := r.FormValue("domainName")

	servers, err := domainProcess.DomainProcess(domain, "", 'n')
	if err == nil {
		fmt.Printf("\n*recived domain name servers: %#v\n\n", servers)

		NSs := &dataProccess.NSData{Domain: domain, Servers: servers}
		dataProccess.NameServersData = append(dataProccess.NameServersData, *NSs)

		err := dataProccess.SaveData(dataProccess.DATAFILE1, dataProccess.NameServersData)
		if err != nil {
			fmt.Println(err)
			return
		}

		myTemplate.Execute(rw, dataProccess.NameServersData)
	} else {
		temp := &dataProccess.NSData{Domain: domain, Servers: []string{"no servers found"}}
		dataProccess.NameServersData = append(dataProccess.NameServersData, *temp)
		myTemplate.Execute(rw, dataProccess.NameServersData)
	}
}
