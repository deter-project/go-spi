package cli

import (
	"flag"
	"fmt"
	"github.com/deter-project/go-spi/spi"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"os/user"
)

var yellow = color.New(color.FgYellow).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var culprit = color.New(color.Bold).Add(color.Underline).SprintFunc()
var User = ""

func Fatal(message string) {
	fmt.Fprintf(os.Stderr, "%s %s\n", red("error:"), message)
	os.Exit(1)
}

func LoadCert() {

	localUser, err := user.Current()
	if err != nil {
		Fatal(fmt.Sprintf("could not get current user: %v", err))
	}

	spiDir := localUser.HomeDir + "/.spi"

	cert, err := ioutil.ReadFile(spiDir + "/spi.cert")
	if err != nil {
		Fatal("certificate read failed")
	}

	spi.SetCertificate(cert)

}

func LoadUser() {

	localUser, err := user.Current()
	if err != nil {
		Fatal(fmt.Sprintf("could not get current user: %v", err))
	}

	spiDir := localUser.HomeDir + "/.spi"

	_user, err := ioutil.ReadFile(spiDir + "/spi.user")
	if err != nil {
		Fatal("user read failed")
	}

	User = string(_user)

}

var debugF = flag.Bool("debug", false, "show xml messages")

func PreRun() {
	LoadCert()
	LoadUser()

	if *debugF {
		spi.Debug = true
	}
}
