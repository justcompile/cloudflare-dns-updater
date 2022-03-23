package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/justcompile/cloudflare-dns-updater/pkg/dns"
	"github.com/rdegges/go-ipify"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	ip, err := ipify.GetIp()
	if err != nil {
		fmt.Println("Couldn't get my IP address:", err)
	} else {
		fmt.Println("My IP address is:", ip)
	}

	for _, zone := range strings.Split(os.Getenv("ZONES_TO_UPDATE"), ",") {
		records, err := dns.RecordsForZone(strings.TrimSpace(zone))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(zone)
		fmt.Println(records)
	}
}
