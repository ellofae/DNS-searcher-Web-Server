package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomePage(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Println(err)
		return
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HomePage)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("handler returned %v", status)
	}
}

func TestGetIP(t *testing.T) {
	t.Run("TestGetIPGetMethod", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/getIP", nil)
		if err != nil {
			log.Println(err)
			return
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HomePage)
		handler.ServeHTTP(rr, req)

		status := rr.Code
		if status != http.StatusOK {
			t.Errorf("handler returned %v", status)
		}
	})

	t.Run("TestGetIPPostMethod", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/getIP", nil)
		if err != nil {
			log.Println(err)
			return
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HomePage)
		handler.ServeHTTP(rr, req)

		status := rr.Code
		if status != http.StatusOK {
			t.Errorf("handler returned %v", status)
		}
	})
}

func TestGetName(t *testing.T) {
	t.Run("TestGetNameGetMethod", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/getName", nil)
		if err != nil {
			log.Println(err)
			return
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HomePage)
		handler.ServeHTTP(rr, req)

		status := rr.Code
		if status != http.StatusOK {
			t.Errorf("handler returned %v", status)
		}
	})

	t.Run("TestGetNamePostMethod", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/getName", nil)
		if err != nil {
			log.Println(err)
			return
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HomePage)
		handler.ServeHTTP(rr, req)

		status := rr.Code
		if status != http.StatusOK {
			t.Errorf("handler returned %v", status)
		}
	})
}

func TestNameServer(t *testing.T) {
	t.Run("TestNameServerGetMethod", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/nameServers", nil)
		if err != nil {
			log.Println(err)
			return
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HomePage)
		handler.ServeHTTP(rr, req)

		status := rr.Code
		if status != http.StatusOK {
			t.Errorf("handler returned %v", status)
		}
	})

	t.Run("TestNameServerPostMethod", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/nameServers", nil)
		if err != nil {
			log.Println(err)
			return
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HomePage)
		handler.ServeHTTP(rr, req)

		status := rr.Code
		if status != http.StatusOK {
			t.Errorf("handler returned %v", status)
		}
	})
}

func TestMailServers(t *testing.T) {
	t.Run("TestMailServersGetMethod", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/mailServers", nil)
		if err != nil {
			log.Println(err)
			return
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HomePage)
		handler.ServeHTTP(rr, req)

		status := rr.Code
		if status != http.StatusOK {
			t.Errorf("handler returned %v", status)
		}
	})

	t.Run("TestMailServersPostMethod", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/mailServers", nil)
		if err != nil {
			log.Println(err)
			return
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HomePage)
		handler.ServeHTTP(rr, req)

		status := rr.Code
		if status != http.StatusOK {
			t.Errorf("handler returned %v", status)
		}
	})
}
