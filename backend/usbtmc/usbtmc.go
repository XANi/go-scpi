package usbtmc

import (
	"os"
	"fmt"
//	"io/ioutil"
	"bufio"
)


type USBTMC struct {
	dev string
}

func NewUSBTMCDevice (dev string) (*USBTMC) {
	var u USBTMC
	u.dev = dev
	return &u
}

func (u *USBTMC)SendAndRecvCmd(cmd string) (string, error) {
	f, err := os.OpenFile(u.dev, os.O_RDWR, 0644)
	if err != nil {
		return "", fmt.Errorf("Error opening %s: %s",u.dev,err)
	}
	defer f.Close()
	fmt.Fprintf(f, "%s\n", cmd)
	sc := bufio.NewScanner(f)
    sc.Scan()
	out := sc.Text()
	return string(out), err
}
func (u *USBTMC)SendCmd(cmd string) (error) {
	f, err := os.OpenFile(u.dev, os.O_RDWR, 0644)
	if err != nil {return err}
	defer f.Close()
	fmt.Fprintf(f, "%s\n", cmd)
	return err
}


func (u *USBTMC)Reset() (error) {
	return u.SendCmd("*RST")
}
