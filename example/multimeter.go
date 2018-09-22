package main

import (
	"github.com/XANi/go-scpi"
	"fmt"
)


func main() {
	scpi := scpi.NewUSBTMCDevice(`/dev/usbtmc0`)
	err := scpi.SendCmd("CONF:VOLT:DC 10,0.01")
	err = scpi.SendCmd(`DISP:TEXT "MEASURING"`)
	out, _ := scpi.SendAndRecvCmd("SAMP:COUN? MAX")
	fmt.Printf("max sample count: %s\n",out)
	err = scpi.SendCmd("SAMP:COUN 50")
	out, err = scpi.SendAndRecvCmd("READ?")
	if err != nil {
		fmt.Printf("Error running command: %s\n", err)
	}
	fmt.Printf("out: %s\n",out)
}
