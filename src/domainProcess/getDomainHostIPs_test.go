package domainProcess

import (
	"sort"
	"testing"
)

func TestGetDomainHost(t *testing.T) {
	tempStruct := []struct {
		ipAddress      string
		expectedValues []string
	}{
		{"127.0.0.1", []string{"DESKTOP-3CLVD04"}},
	}

	for _, testData := range tempStruct {
		testSlice, err := getDomainHosts(testData.ipAddress)
		if err != nil {
			t.Errorf("Dind't manage to run the tested function")
		}

		if len(testSlice) != len(testData.expectedValues) {
			t.Errorf("Amount of data recived is not satisfied!")
		}

		for ind, value := range testSlice {
			if value != testData.expectedValues[ind] {
				t.Errorf("Data that was recived is not as expected!")
			}
		}
	}
}

func TestGetIPAddresses(t *testing.T) {
	tempStruct := []struct {
		ipAddress      string
		expectedValues []string
	}{
		{"www.github.com", []string{"140.82.121.4"}},
		{"www.google.com", []string{"216.58.209.164"}},
		{"www.kubernetes.io", []string{"147.75.40.148"}},
	}

	for _, testData := range tempStruct {
		testSlice, err := getDomainIPs(testData.ipAddress)
		if err != nil {
			t.Errorf("Dind't manage to run the tested function")
		}

		if len(testSlice) != len(testData.expectedValues) {
			t.Errorf("Amount of data recived is not satisfied!")
		}

		for ind, value := range testSlice {
			if value != testData.expectedValues[ind] {
				t.Errorf("Data that was recived is not as expected!")
			}
		}
	}
}

func TestGetNameServers(t *testing.T) {
	tempStruct := []struct {
		domainName     string
		expectedValues []string
	}{
		{"www.github.com", []string{"dns2.p08.nsone.net.", "dns3.p08.nsone.net.", "dns4.p08.nsone.net.", "ns-1283.awsdns-32.org.", "ns-1707.awsdns-21.co.uk.", "ns-421.awsdns-52.com.", "ns-520.awsdns-01.net.", "dns1.p08.nsone.net."}},
	}

	for _, testData := range tempStruct {
		testSlice, err := getNameServers(testData.domainName)
		sort.Strings(testSlice)
		if err != nil {
			t.Errorf("Dind't manage to run the tested function")
		}

		if len(testSlice) != len(testData.expectedValues) {
			t.Errorf("Amount of data recived is not satisfied!")
		}

		sort.Strings(testData.expectedValues)
		for ind, value := range testSlice {
			if value != testData.expectedValues[ind] {
				t.Errorf("Data that was recived is not as expected! exptected: %s, recived: %s", testData.expectedValues[ind], value)
			}
		}
	}
}

func TestGetMailServers(t *testing.T) {
	tempStruct := []struct {
		domainName     string
		expectedValues []string
	}{
		{"www.github.com", []string{"aspmx.l.google.com.", "alt1.aspmx.l.google.com.", "alt2.aspmx.l.google.com.", "alt3.aspmx.l.google.com.", "alt4.aspmx.l.google.com."}},
	}

	for _, testData := range tempStruct {
		testSlice, err := getMailServers(testData.domainName)
		sort.Strings(testSlice)
		if err != nil {
			t.Errorf("Dind't manage to run the tested function")
		}

		if len(testSlice) != len(testData.expectedValues) {
			t.Errorf("Amount of data recived is not satisfied!")
		}

		sort.Strings(testData.expectedValues)
		for ind, value := range testSlice {
			if value != testData.expectedValues[ind] {
				t.Errorf("Data that was recived is not as expected! exptected: %s, recived: %s", testData.expectedValues[ind], value)
			}
		}
	}
}
