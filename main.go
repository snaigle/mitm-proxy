package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type ProxyHandler struct {
	Tr *http.Transport
}

func (h *ProxyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	scheme := "http"
	if req.TLS != nil {
		scheme = "https"
	}
	urlStr := scheme + "://" + req.Host + req.URL.Path
	if req.URL.RawQuery != "" {
		urlStr += "?" + req.URL.RawQuery
	}
	req.URL, _ = url.Parse(urlStr)

	resp, err := h.Tr.RoundTrip(req)
	if err != nil {
		log.Println("url:", req.URL, "failed", err)
		w.WriteHeader(502)
		w.Write([]byte("proxy error"))
		return
	}
	copyHeaders(w.Header(), resp.Header, false)
	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	if err := resp.Body.Close(); err != nil {
		log.Println("proxy close resp body error", err)
	}
}
func copyHeaders(dst, src http.Header, keepDestHeaders bool) {
	if !keepDestHeaders {
		for k := range dst {
			dst.Del(k)
		}
	}
	for k, vs := range src {
		for _, v := range vs {
			dst.Add(k, v)
		}
	}
}

func main() {
	proxyHandler := &ProxyHandler{
		Tr: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy: func(r *http.Request) (*url.URL, error) {
				ip := GetIpFromRemoteAddr(r.RemoteAddr)
				proxyIp := GetProxy(ip)
				log.Println("url:", r.URL, ",remoteIp:", ip, ",proxyIp:", proxyIp)
				if proxyIp == "" {
					return nil, errors.New(fmt.Sprint(ip, "has no proxy"))
				} else {
					return url.Parse(fmt.Sprintf("http://%v", proxyIp))
				}
			},
		},
	}
	certs := initCerts()
	// listen 80
	// listen 443
	go func() {
		log.Println("start listen :80")
		err := http.ListenAndServe(":80", proxyHandler)
		log.Println("end listen :80")
		if err != nil {
			log.Println("listen :80 failed", err)
			return
		}
	}()
	go func() {
		server := http.Server{
			Addr:    ":443",
			Handler: proxyHandler,
			TLSConfig: &tls.Config{
				Certificates: certs,
			},
		}
		log.Println("start listen :443")
		err := server.ListenAndServeTLS("", "")
		log.Println("end listen :443")
		if err != nil {
			log.Println("listen :443 failed", err)
		}
	}()
	// listen admin port 8080
	go func() {
		startAdmin()
	}()
	log.Println("proxy server starting")
	c := make(chan int, 1)
	<-c
}

func initCerts() []tls.Certificate {
	certs := make([]tls.Certificate, 0)
	if len(os.Args) < 3 {
		return certs
	}
	log.Println("args:", os.Args)
	privKeys := strings.Split(os.Args[1], ",")
	pubKeys := strings.Split(os.Args[2], ",")
	if len(privKeys) != len(pubKeys) {
		log.Println("certs not valid")
		os.Exit(-1)
	}
	for i, certPath := range privKeys {
		keyPath := pubKeys[i]
		cert, err := ioutil.ReadFile(certPath)
		if err != nil {
			log.Println("load", certPath, "failed", err)
			continue
		}
		key, err := ioutil.ReadFile(keyPath)
		if err != nil {
			log.Println("load", keyPath, "failed", err)
			continue
		}
		certificate, err := newCaCert(cert, key)
		if err != nil {
			log.Println("new ca cert[", certPath, ",", keyPath, "]", "failed", err)
			continue
		}
		certs = append(certs, certificate)
	}
	return certs
}
func newCaCert(caCert []byte, caKey []byte) (tls.Certificate, error) {
	ca, err := tls.X509KeyPair(caCert, caKey)
	if err != nil {
		log.Println("NewCaCert error:", err)
		return ca, err
	}
	if ca.Leaf, err = x509.ParseCertificate(ca.Certificate[0]); err != nil {
		log.Println("NewCaCert error:", err)
		return ca, err
	}
	log.Println("NewCaCert Ok")
	return ca, nil
}
