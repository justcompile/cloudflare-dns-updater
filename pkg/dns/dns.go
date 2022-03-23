package dns

import (
	"context"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

func RecordsForZone(zoneName string) ([]cloudflare.DNSRecord, error) {
	api, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))
	if err != nil {
		return nil, err
	}

	// Most API calls require a Context
	ctx := context.Background()

	zoneId, err := api.ZoneIDByName(zoneName)
	if err != nil {
		return nil, err
	}

	return api.DNSRecords(ctx, zoneId, cloudflare.DNSRecord{Type: "A"})
}
