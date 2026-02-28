package uip

import (
	"path/filepath"
	"testing"
)

func TestSearchIPv4WithXdb(t *testing.T) {
	if testing.Short() {
		t.Skip("skip integration test in short mode")
	}

	dbPath := filepath.Join("..", "testdata", "ip2region_v4.xdb")
	MustInitIp2Region(dbPath)
	t.Cleanup(func() {
		CloseIp2Region()
		Ip2Region = nil
	})

	detail, err := SearchIPv4("114.114.114.114")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if detail == nil {
		t.Fatalf("expected detail")
	}
	if detail.Country == "" {
		t.Fatalf("country missing")
	}
	if detail.Province == "" {
		t.Fatalf("province missing")
	}
	if detail.City == "" {
		t.Fatalf("city missing")
	}
	if detail.CountryISO == "" {
		t.Fatalf("country iso missing")
	}
	t.Logf("SearchIPv4(\"114.114.114.114\") %+v", detail)
}
