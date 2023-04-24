package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/ellofae/DNS-searcher-Web-Server/src/domainProcess"

	"github.com/ellofae/DNS-searcher-Web-Server/src/dataProccess"
)

type GetName struct {
	l *log.Logger
}

func NewGetName(l *log.Logger) *GetName {
	return &GetName{l}
}

func (g *GetName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Domain name service getter requested")

	fmt.Printf("Serving %s for %s\n", r.Host, r.URL.Path)
	myTemplate := template.Must(template.ParseFiles("getDomainPage.html"))

	if r.Method != http.MethodPost {
		myTemplate.Execute(rw, dataProccess.DomainNamesData)
		return
	}

	ip := r.FormValue("ip")

	domains, err := domainProcess.DomainProcess("", ip, 'd')
	if err == nil {
		fmt.Printf("\n*recived Hosts: %#v\n\n", domains)

		hosts := &dataProccess.DomainData{IPaddress: ip, DomainNames: domains}
		dataProccess.DomainNamesData = append(dataProccess.DomainNamesData, *hosts)

		err := dataProccess.SaveData(dataProccess.DATAFILE3, dataProccess.DomainNamesData)
		if err != nil {
			fmt.Println(err)
			return
		}

		myTemplate.Execute(rw, dataProccess.DomainNamesData)
	} else {
		hosts := &dataProccess.DomainData{IPaddress: ip, DomainNames: []string{"no domains found"}}
		dataProccess.DomainNamesData = append(dataProccess.DomainNamesData, *hosts)
		myTemplate.Execute(rw, dataProccess.DomainNamesData)
	}

}
