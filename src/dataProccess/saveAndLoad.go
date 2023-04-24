package dataProccess

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	DATAFILE1 string = "/tmp/domain-name-server-data.json"
	DATAFILE2 string = "/tmp/mail-server-data.json"
	DATAFILE3 string = "/tmp/domain-name-data.json"
	DATAFILE4 string = "/tmp/ip-address-data.json"
)

// Data defenition
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
var DomainNamesData = make([]DomainData, 0)
var NameServersData = make([]NSData, 0)
var MailServersData = make([]MXData, 0)
var AddressData = make([]AddrData, 0)

// Saving and loading data
func SaveData(DATA string, SomeStruct interface{}) error {
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

func LoadData(DATA string, SomeStruct interface{}) error {
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

func LoadAllData() error {
	var err error

	err = LoadData(DATAFILE1, &NameServersData)
	if err != nil {
		fmt.Println(err)
	}

	err = LoadData(DATAFILE2, &MailServersData)
	if err != nil {
		fmt.Println(err)
	}

	err = LoadData(DATAFILE3, &DomainNamesData)
	if err != nil {
		fmt.Println(err)
	}

	err = LoadData(DATAFILE4, &AddressData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nLoading is completed!\n\n")

	return nil
}
