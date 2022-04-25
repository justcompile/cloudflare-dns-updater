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
	}

	log.Println("Current IP:", ip)

	for _, zone := range strings.Split(os.Getenv("ZONES_TO_UPDATE"), ",") {
		if err := dns.UpdateRecordsForZone(strings.TrimSpace(zone), ip); err != nil {
			log.Fatal(err)
		}
	}
}
