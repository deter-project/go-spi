package spi

import (
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"encoding/xml"
	"errors"
	"log"
	"net/http"
	"os"
	"regexp"
)

//Common variables shared amongst the binding components -----------------------

const (
	//note: you must set your /etc/hosts for this entry
	API_HOST = "spi.deterlab.net"

	API_PORT = "52323"

	API_HTTPS = "https://" + API_HOST + ":" + API_PORT + "/axis2/services"
)

//This has to be here to allow for self-signed certifacates to work
var certPool = x509.NewCertPool()
var tr = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true, RootCAs: certPool},
}
var client = &http.Client{Transport: tr}

//Common messaging structs -----------------------------------------------------
type Envelope struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`
	Body    Body
}

type Body struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Body"`
}

//Common functionality ---------------------------------------------------------

/*
spiCall encapsulates most of the minutia associated with making soap POST calls.
The addr paramter specifies a URL address at which the soap message is directed.
The message parmeter is the message of the POST request. The result of the
message is read from the io.Reader interface in the message response and handed
back into the result parameter. Thus the result parameter should be a pointer
to the desired unmarshalled data type.
*/
func spiCall(addr string, message interface{}, result interface{}) (
	*http.Response, string, error) {

	//create the envelope
	msg, err := xml.Marshal(&message)
	if err != nil {
		log.Println(err)
		return nil, "", err
	}

	//make the request
	req, err := http.NewRequest("POST", addr, bytes.NewBuffer(msg))
	if err != nil {
		log.Println(err)
		return nil, "", err
	}
	req.Header.Add("Content-Type", "application/soap+xml")

	//send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return resp, "", err
	}

	//read the result
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	bs := buf.String()

	xml.Unmarshal([]byte(bs), result)

	return resp, bs, nil

}

var certRx = regexp.MustCompile(
	`-----BEGIN CERTIFICATE-----(.|\n|\r)*-----END CERTIFICATE-----`)

var keyRx = regexp.MustCompile(
	`-----BEGIN RSA PRIVATE KEY-----(.|\n|\r)*-----END RSA PRIVATE KEY-----`)

func addSpiCert() error {

	var spicertBits = `-----BEGIN CERTIFICATE-----
MIIDQjCCAiqgAwIBAgIEUikbXDANBgkqhkiG9w0BAQUFADBjMQswCQYDVQQGEwJVUzELMAkGA1UE
CBMCQ0ExDDAKBgNVBAcTA01kUjERMA8GA1UEChMIREVURVJMYWIxEDAOBgNVBAsTB1Vua25vd24x
FDASBgNVBAMTC3ZpbS5pc2kuZWR1MB4XDTEzMDkwNjAwMDEzMloXDTIzMDkwNDAwMDEzMlowYzEL
MAkGA1UEBhMCVVMxCzAJBgNVBAgTAkNBMQwwCgYDVQQHEwNNZFIxETAPBgNVBAoTCERFVEVSTGFi
MRAwDgYDVQQLEwdVbmtub3duMRQwEgYDVQQDEwt2aW0uaXNpLmVkdTCCASIwDQYJKoZIhvcNAQEB
BQADggEPADCCAQoCggEBAIvCETRdViQtygsbAvdMdjmlRPglUskV7C60gxLEebfIGNxeuHCh0hS4
mbpsJGPO+vJXSJScZQFPrd07vK5M5Zk3kUvYTp0TG0bavwqcYLc5J695gBPDa8DskTtHAiUC29hz
JG5yihkeIiozzbjMqGJUaLXmJK1U917QEU2MXVPLP1S6j4i18GGlWk/ouKkuEkPX/jUPPQf6Na5n
T5G7rMwGtIncbqRFqk2FJcnzbqHStlfol/d7sLx6tKSTMeRk/wzjvVo/AtJHNJUJzOO6BPL4KYNu
L21awtjWp+zmyUbgLpz1aZyB+hoxxlMqjVbQNytGsSVQ/RWP/UGkSKOND2cCAwEAATANBgkqhkiG
9w0BAQUFAAOCAQEAFe70D0jDz2nBM3ppnzn6CzvJN6XOpvdEak861WDpjzGEdblopweJLrja80Di
HY/RGztcwgZnCYY3Vnu9mjoZUopJ6gO+b3Uzb/nUij2bsUs8tl4+Gn+8rTNAbzhErHTvp1MkN9yq
qcCaV1nRVKWYIpyBCvjNQD5QmDO7N3mMWKCn+5hwujim8GiY9Gmpyrt5fJbIGz+5m1kyWY1iGfPL
m+HZRfsPB5qo7jx6lGI1Y7+VOxCYHYjsCNrUt+bIO+geR4WfVK9idz8kVLaAIH0mJG6LxqIfe+gQ
mRhSy2Hpoey/h99fZJbRTQ1cUhKRvodImOvdp7b0V55ybm6FDMWxoQ==
-----END CERTIFICATE-----`

	pemBlock, _ := pem.Decode([]byte(spicertBits))
	if pemBlock == nil {
		err := errors.New("could not decode PEM block")
		log.Println(err)
		return err
	}

	spicert, err := x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		log.Println(err)
		return err
	}
	certPool.AddCert(spicert)

	return nil

}

func init() {
	err := addSpiCert()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

/*
setCertificate sets the certificate that is used when communicating with the
SPI server
*/
func SetCertificate(comboCert []byte) error {

	certBits := certRx.Find(comboCert)
	keyBits := keyRx.Find(comboCert)

	x509, err := tls.X509KeyPair(certBits, keyBits)
	if err != nil {
		log.Println(err)
		return err
	}

	tr.TLSClientConfig.Certificates = []tls.Certificate{x509}
	tr.TLSClientConfig.Rand = rand.Reader
	tr.TLSClientConfig.BuildNameToCertificate()
	tr.TLSClientConfig.NameToCertificate = nil

	tr.CloseIdleConnections()

	return nil
}
