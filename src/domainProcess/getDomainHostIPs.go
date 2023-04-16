package domainProcess

import (
	"errors"
	"fmt"
	"log"
	"net"
)

// DNS search execution
func DomainProcess(domain string, ip string, flagSpec byte) ([]string, error) {
	dnsResult := make([]string, 0)
	var err error

	switch flagSpec {
	case 'd':
		dnsResult, err = getDomainHosts(ip)
		return dnsResult, err
	case 'i':
		dnsResult, err = getDomainIPs(domain)
		return dnsResult, err
	case 'n':
		dnsResult, err = getNameServers(domain)
		return dnsResult, err
	case 'm':
		dnsResult, err = getMailServers(domain)
		return dnsResult, err
	default:
		return nil, errors.New("Incorrect flag passed")
	}
}

// Returns names mapping to the passed IP address
func getDomainHosts(ip string) ([]string, error) {
	hostsSlice := make([]string, 0)

	ipAddr := net.ParseIP(ip)
	if ipAddr != nil {
		hosts, err := getHosts(ip)

		if err == nil {

			for _, hostName := range hosts {
				hostsSlice = append(hostsSlice, hostName)
				fmt.Printf("\tFor IP address %#v found host name: %#v\n", ip, hostName)
			}
		}
	}

	return hostsSlice, nil
}

// Returns IP addresses mapping to the domain host names
func getDomainIPs(host string) ([]string, error) {
	ipsSlice := make([]string, 0)

	hostCheck := net.ParseIP(host)
	if hostCheck == nil {
		hosts, err := getIPs(host)

		if err == nil {

			for _, ip := range hosts {
				ipsSlice = append(ipsSlice, ip)
				fmt.Printf("\tFor domain host name %#v found IP: %#v\n", host, ip)
			}
		}
	}

	return ipsSlice, nil
}

// Getting a domain name server using NS recods of a domain
func getNameServers(domain string) ([]string, error) {
	nameServersSlice := make([]string, 0)

	NSs, err := net.LookupNS(domain)
	if err != nil {
		log.Println("Didn't manage to get a domain name server from domain NS records occured ")
		return nil, err
	}

	for _, ns := range NSs {
		nameServersSlice = append(nameServersSlice, ns.Host)
		fmt.Printf("\tFor domain host name %#v found domain name server: %#v\n", domain, ns.Host)
	}
	fmt.Println()

	return nameServersSlice, nil
}

func getMailServers(domain string) ([]string, error) {
	mailServersSlice := make([]string, 0)

	MXs, err := net.LookupMX(domain)
	if err != nil {
		log.Println("Didn't manage to get a domain mail server from domain MX records occured")
		return nil, err
	}

	for _, mx := range MXs {
		mailServersSlice = append(mailServersSlice, mx.Host)
		fmt.Printf("\tFor domain host name %#v found domain mail server: %#v\n", domain, mx.Host)
	}
	fmt.Println()

	return mailServersSlice, nil
}

// Getting names mapped to an IP address helping function
func getHosts(ip string) ([]string, error) {
	hosts, err := net.LookupAddr(ip)
	if err != nil {
		return nil, err
	}

	return hosts, nil
}

// Getting IP addresses mapped to domain hosts names helping function
func getIPs(host string) ([]string, error) {
	IPs, err := net.LookupHost(host)
	if err != nil {
		return nil, err
	}

	return IPs, nil
}
