package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/ellofae/DNS-searcher-Web-Server/src/domainProcess"

	"github.com/ellofae/DNS-searcher-Web-Server/src/dataProccess"
)

type GetMailServers struct {
	l *log.Logger
}

func NewGetMailServers(l *log.Logger) *GetMailServers {
	return &GetMailServers{l}
}

func (g *GetMailServers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Mail Servers getter requested")

	fmt.Printf("Serving %s for %s\n", r.Host, r.URL.Path)
	myTemplate := template.Must(template.ParseFiles("getMailServerPage.html"))

	if r.Method != http.MethodPost {
		myTemplate.Execute(rw, dataProccess.MailServersData)
		return
	}

	domain := r.FormValue("domainName")

	servers, err := domainProcess.DomainProcess(domain, "", 'm')
	if err == nil {
		fmt.Printf("\n*recived mail servers: %#v\n\n", servers)

		MXs := &dataProccess.MXData{Domain: domain, Servers: servers}
		dataProccess.MailServersData = append(dataProccess.MailServersData, *MXs)

		err := dataProccess.SaveData(dataProccess.DATAFILE2, dataProccess.MailServersData)
		if err != nil {
			fmt.Println(err)
			return
		}

		myTemplate.Execute(rw, dataProccess.MailServersData)
	} else {
		temp := &dataProccess.MXData{Domain: domain, Servers: []string{"no servers found"}}
		dataProccess.MailServersData = append(dataProccess.MailServersData, *temp)
		myTemplate.Execute(rw, dataProccess.MailServersData)
	}
}
