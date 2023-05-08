package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

func runMove(ch chan Comand) {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}

	var serverIP net.IP
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				serverIP = ipnet.IP
				break
			}
		}
	}

	if serverIP == nil {
		fmt.Println("Nepodarilo sa nájsť IP adresu servera")
		return
	} else {
		fmt.Printf("IP adresa servera: %s\n", serverIP.String())
	}
	r := gin.Default()

	r.GET("/direction", func(c *gin.Context) {
		d := c.Query("d")
		fmt.Println(d)
		switch d {
		case "up":
			ch <- UP
		case "down":
			ch <- DOWN
		case "right":
			ch <- RIGHT
		case "left":
			ch <- LEFT
		case "tuk":
			ch <- LEFTCLICK
		case "menu":
			ch <- RIGHTCLICK
		case "plus":
			ch <- PLUS
		case "minus":
			ch <- MINUS
		default:
			fmt.Println(d)
		}
		c.String(http.StatusOK, "Ahoj ok, %s", d)

	})

	err = r.Run(":8080")
	if err != nil {
		return
	}
}
