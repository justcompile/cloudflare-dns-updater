package dns

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

func ListRecordsForZone(zoneName string) ([]cloudflare.DNSRecord, error) {
	api, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))
	if err != nil {
		return nil, err
	}

	// Most API calls require a Context
	ctx := context.Background()

	zoneId, err := api.ZoneIDByName(zoneName)
	if err != nil {
		return nil, fmt.Errorf("zone by name: %s", err.Error())
	}

	return api.DNSRecords(ctx, zoneId, cloudflare.DNSRecord{Type: "A"})
}

func UpdateRecordsForZone(zoneName, ipAddress string) error {
	api, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))
	if err != nil {
		return errors.Wrap(err, "cloudflare auth")
	}

	// Most API calls require a Context
	ctx := context.Background()

	zoneId, err := api.ZoneIDByName(zoneName)
	if err != nil {
		return errors.Wrap(err, "zone lookup")
	}

	records, err := api.DNSRecords(ctx, zoneId, cloudflare.DNSRecord{Type: "A"})
	if err != nil {
		return errors.Wrap(err, "list records")
	}

	for _, r := range records {
		if r.Content == ipAddress {
			log.Printf("%s is already up-to-date", r.Name)
			continue
		}

		r.Content = ipAddress

		if err := api.UpdateDNSRecord(ctx, zoneId, r.ID, r); err != nil {
			return errors.Wrap(err, "record update")
		}

		log.Printf("updated %s => %s", r.Name, r.Content)
	}

	return nil
}
