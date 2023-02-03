package test

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
)

func test() {

	//cert, err := os.ReadFile("../certificate-generator/ca-cert.pem")
	cert, err := os.ReadFile("ca-cert.pem")

	if err != nil {
		fmt.Println(err)
	}
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(cert); !ok {
		fmt.Println(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: certPool,
			},
		},
	}

	r, err := client.Get("https://127.0.0.1:6663")
	if err != nil {
		fmt.Println(err)

	}
	defer r.Body.Close()

	html, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", r.Status)
	fmt.Printf(string(html))
}
