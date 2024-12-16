package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	tld "github.com/jpillora/go-tld"
)

var build string

func main() {
	var url string
	var get string
	var v bool

	flag.StringVar(&url, "url", "", "Enter the URL to parse. Example:\n"+"https://example.com")

	flag.StringVar(&get, "get", "", "Get value needed from the URL. Can be:\n"+
		"host, hostWithoutPort, domain, subdomain, tld, port, path, fragment, scheme, registeredDomain, query.<paramName>")

	flag.BoolVar(&v, "v", false, "Prints the current build version and exits")

	flag.Parse()

	if v {
		fmt.Println(build)
		return
	}

	log.SetFlags(0)

	if len(url) == 0 || len(get) == 0 {
		log.Fatal("Both --url and --get flags are required.")
	}

	p, err := tld.Parse(url)

	if err != nil {
		log.Fatal("Error parsing url: ", err)
	}

	if strings.HasPrefix(get, "query.") {
		param := strings.Split(get, ".")[1]
		val := p.Query().Get(param)
		if len(val) == 0 {
			log.Fatalf("\"%s\" query param not found in url", param)
		}
		fmt.Print(val)
		return
	}

	switch get {
	case "host":
		fmt.Print(p.Host)

	case "hostWithoutPort":
		fmt.Print(p.Hostname())

	case "domain":
		fmt.Print(p.Domain)

	case "subdomain":
		fmt.Print(p.Subdomain)

	case "tld":
		fmt.Print(p.TLD)

	case "port":
		fmt.Print(p.Port)

	case "path":
		fmt.Print(p.Path)

	case "fragment":
		fmt.Print(p.Fragment)

	case "scheme":
		fmt.Print(p.Scheme)

	case "registeredDomain":
		fmt.Printf("%s.%s", p.Domain, p.TLD)

	default:
		log.Fatal("Invalid --get option")
	}

}
