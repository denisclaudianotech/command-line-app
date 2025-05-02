package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e nomes de Servidores na internet"

	// Flags
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "biodoc.com.br",
		},
	}

	// Adicionando os camondos
	app.Commands = []cli.Command{
		// Comando 1 buscar IPs
		{
			Name:  "ip",
			Usage: "Busca IPs de endereços na internet",
			// Flag é um slice
			Flags:  flags,
			Action: buscarIps,
		},
		// Comando 2 busca por servidores
		{
			Name:   "servidores",
			Usage:  "Busca pelo nome dos servidores",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	// Aqui usamos o pacote net para buscar os IPs
	// LookupIp vai retornar um slice porque dependendo do redirecionamento do site pode ser que ele tenha mais de um IP publico
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	// Chamando um função do pacote NET para retornar os servidores
	servidores, erro := net.LookupNS(host) // NS -> name server
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}

}
