package application

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Gerar retorna aplicação de linha de comando
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Retorna IP e nomes de servidores na internet"
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "Google.com.br",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IP de sites",
			Flags:  flags,
			Action: BuscarIps,
		},
		{
			Name:   "servers",
			Usage:  "Procura nome de servidores na internet",
			Flags:  flags,
			Action: BuscarServidores,
		},
	}
	return app
}

func BuscarIps(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro, "\n")
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func BuscarServidores(c *cli.Context) {
	host := c.String("host")
	servers, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, servidor := range servers {
		fmt.Println(servidor.Host)
	}
}
