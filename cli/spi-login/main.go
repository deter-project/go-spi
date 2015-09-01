package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/deter-project/go-spi/cli"
	"github.com/deter-project/go-spi/spi"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
)

func getCreds() (string, string) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("username: ")
	usr, _ := reader.ReadString('\n')
	usr = strings.TrimSuffix(usr, "\n")
	fmt.Print("password: ")
	pwd, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println("")

	return usr, string(pwd)

}

var usrF *string = nil
var pwdF *string = nil

func init() {

	usrF = flag.String("user", "", "SPI username")
	pwdF = flag.String("password", "", "SPI password")

}

func login() {

	usr, pwd := "", ""

	if *usrF != "" && *pwdF != "" {
		usr = *usrF
		pwd = *pwdF
	} else {
		usr, pwd = getCreds()
	}

	cert, err := spi.Login(usr, pwd)
	if err != nil {
		cli.Fatal(fmt.Sprintf("login failed: %v", err))
	}

	localUser, err := user.Current()
	if err != nil {
		cli.Fatal(fmt.Sprintf("could not get current user: %v", err))
	}

	spiDir := localUser.HomeDir + "/.spi"
	os.MkdirAll(spiDir, 0755)

	ioutil.WriteFile(spiDir+"/spi.cert", cert, 0644)
	ioutil.WriteFile(spiDir+"/spi.user", []byte(usr), 0644)
}

func main() {
	flag.Parse()
	login()
}
