package promotion

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestPromotion(t *testing.T) {
	loc_tokyo, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatalf("LoadLocation loc1 failed:%v\n", err)
	}
	loc_newyork, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatalf("LoadLocation loc1 failed:%v\n", err)
	}

	/*
		loc1, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			t.Fatalf("LoadLocation loc1 failed:%v\n", err)
		}
				loc2, err := time.LoadLocation("America/New_York")
				if err != nil {
					t.Fatalf("LoadLocation loc2 failed:%v\n", err)
				}

			utc, err := time.LoadLocation("UTC")
			if err != nil {
				t.Fatalf("LoadLocation utc failed:%v\n", err)
			}
	*/

	// CAMPAIGN1   2018/10/01 00:00 JP - 2018/10/14 00:00 JP
	camp1_fr := time.Date(2018, time.October, 01, 0, 0, 0, 0, loc_tokyo)
	camp1_to := time.Date(2018, time.October, 14, 0, 0, 0, 0, loc_tokyo)

	// CAMPAIGN2   2018/10/07 09:00 EDT - 2018/10/21 18:00 EDT
	camp2_fr := time.Date(2018, time.October, 07, 9, 0, 0, 0, loc_newyork)
	camp2_to := time.Date(2018, time.October, 21, 18, 0, 0, 0, loc_newyork)

	camps := []struct {
		name     string
		from, to time.Time
		banner   string
	}{
		{name: "CAMPAIGN1", from: camp1_fr, to: camp1_to, banner: "<some>CAMPAIGN 1</some>"},
		{name: "CAMPAIGN2", from: camp2_fr, to: camp2_to, banner: "<some>CAMPAIGN 2</some>"},
	}

	for _, camp := range camps {
		err := NewCampaign(camp.name, camp.from, camp.to, camp.banner).Add()
		if err != nil {
			t.Fatalf("add campaign failed: %v", err)
		}
	}

	req_local, err := http.NewRequest("GET", "dummy", nil)
	if err != nil {
		fmt.Printf("NewRequest error:%v\n", err)
	}
	tests := []struct {
		now     time.Time
		request *http.Request
		banner  string
	}{
		// before campaigns
		{now: time.Date(2018, time.September, 30, 23, 59, 59, 0, loc_tokyo), request: req_local, banner: ""},
		{now: time.Date(2018, time.September, 30, 10, 59, 59, 0, loc_newyork), request: req_local, banner: ""},
		// campaign 1 started
		{now: time.Date(2018, time.October, 01, 00, 00, 0, 0, loc_tokyo), request: req_local, banner: "<some>CAMPAIGN 1</some>"},
		{now: time.Date(2018, time.September, 30, 11, 00, 0, 0, loc_newyork), request: req_local, banner: "<some>CAMPAIGN 1</some>"},
		// campaign 1 expired
		{now: time.Date(2018, time.October, 14, 00, 00, 0, 1, loc_tokyo), request: req_local, banner: "<some>CAMPAIGN 2</some>"},
		{now: time.Date(2018, time.October, 30, 11, 00, 0, 0, loc_newyork), request: req_local, banner: "<some>CAMPAIGN 2</some>"},
		// after campaigns
		{now: time.Date(2018, time.October, 22, 07, 00, 0, 1, loc_tokyo), request: req_local, banner: ""},
		{now: time.Date(2018, time.October, 21, 18, 00, 0, 1, loc_newyork), request: req_local, banner: ""},
	}
	list()

	for _, c := range tests {
		now = func() time.Time { return c.now }
		banner, err := Banner(c.request)
		if err != nil {
			t.Fatalf("Banner func failed:%v\n", err)
		}
		fmt.Printf("now=%v banner=%v\n", now(), banner)
		if c.banner != banner {
			t.Errorf("received: %v - expected: %v - case: %v", banner, c.banner, c)
		}
	}
}
