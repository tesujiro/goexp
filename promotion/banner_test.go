package promotion

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestConcurrentAddCampaign(t *testing.T) {
	conc := 10000 // Concurrency
	camp_fr := time.Date(2018, time.October, 01, 0, 0, 0, 0, time.UTC)
	camp_to := time.Date(2018, time.October, 14, 0, 0, 0, 0, time.UTC)
	wg := &sync.WaitGroup{}
	for i := 0; i < conc; i++ {
		wg.Add(1)
		go func() {
			err := NewCampaign(
				fmt.Sprintf("Camp%v", i),
				camp_fr,
				camp_to,
				fmt.Sprintf("Camp%v", i),
			).Add()
			if err != nil {
				t.Fatalf("add campaign failed: %v", err)
			}

			wg.Done()
		}()
	}
	wg.Wait()

	count := countCampaigns()
	if count != conc {
		t.Errorf("campaigns number:%v - expected: %v ", count, conc)
	}
	truncate()
}

func setCampaignsCase1(t *testing.T) {
	loc_tokyo, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatalf("LoadLocation loc1 failed:%v\n", err)
	}
	loc_newyork, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatalf("LoadLocation loc1 failed:%v\n", err)
	}

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

	for i, camp := range camps {
		t.Logf("promotion:%v\t%v(%v-%v) banner:%v", i, camp.name, camp.from, camp.to, camp.banner)
		err := NewCampaign(camp.name, camp.from, camp.to, camp.banner).Add()
		if err != nil {
			t.Fatalf("add campaign failed: %v", err)
		}
	}

}

func TestCampaignPeriod(t *testing.T) {
	// set campaigns case 1
	setCampaignsCase1(t)

	loc_tokyo, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatalf("LoadLocation loc1 failed:%v\n", err)
	}
	loc_newyork, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatalf("LoadLocation loc1 failed:%v\n", err)
	}

	// request from user
	ReqFromUser, err := http.NewRequest("GET", "DUMMY URL", nil)
	if err != nil {
		fmt.Printf("NewRequest error:%v\n", err)
	}
	ReqFromUser.RemoteAddr = "127.0.0.1:12345"

	// request from admin1
	ReqFromAdmin1, err := http.NewRequest("GET", "DUMMY URL", nil)
	if err != nil {
		fmt.Printf("NewRequest error:%v\n", err)
	}
	ReqFromAdmin1.RemoteAddr = "10.0.0.1:12345"

	// request from admin2
	ReqFromAdmin2, err := http.NewRequest("GET", "DUMMY URL", nil)
	if err != nil {
		fmt.Printf("NewRequest error:%v\n", err)
	}
	ReqFromAdmin2.RemoteAddr = "10.0.0.2:12345"

	tests := []struct {
		now     time.Time
		request *http.Request
		banner  string
	}{
		// before campaigns
		{now: time.Date(2018, time.September, 30, 23, 59, 59, 0, loc_tokyo), request: ReqFromUser, banner: ""},
		{now: time.Date(2018, time.September, 30, 10, 59, 59, 0, loc_newyork), request: ReqFromUser, banner: ""},
		// campaign 1 started
		{now: time.Date(2018, time.October, 01, 00, 00, 0, 0, loc_tokyo), request: ReqFromUser, banner: "<some>CAMPAIGN 1</some>"},
		{now: time.Date(2018, time.September, 30, 11, 00, 0, 0, loc_newyork), request: ReqFromUser, banner: "<some>CAMPAIGN 1</some>"},
		// before campaign 2 start
		{now: time.Date(2018, time.October, 07, 21, 59, 59, 999, loc_tokyo), request: ReqFromUser, banner: "<some>CAMPAIGN 1</some>"},
		{now: time.Date(2018, time.October, 07, 8, 59, 59, 999, loc_newyork), request: ReqFromUser, banner: "<some>CAMPAIGN 1</some>"},
		// campaign 2 start
		{now: time.Date(2018, time.October, 07, 22, 00, 0, 0, loc_tokyo), request: ReqFromUser, banner: "<some>CAMPAIGN 1</some>"},
		{now: time.Date(2018, time.October, 07, 9, 00, 0, 0, loc_newyork), request: ReqFromUser, banner: "<some>CAMPAIGN 1</some>"},
		// campaign 1 last moment
		{now: time.Date(2018, time.October, 14, 00, 00, 0, 0, loc_tokyo), request: ReqFromUser, banner: "<some>CAMPAIGN 1</some>"},
		{now: time.Date(2018, time.October, 13, 11, 00, 0, 0, loc_newyork), request: ReqFromUser, banner: "<some>CAMPAIGN 1</some>"},
		// campaign 1 expired
		{now: time.Date(2018, time.October, 14, 00, 00, 0, 1, loc_tokyo), request: ReqFromUser, banner: "<some>CAMPAIGN 2</some>"},
		{now: time.Date(2018, time.October, 13, 11, 00, 0, 1, loc_newyork), request: ReqFromUser, banner: "<some>CAMPAIGN 2</some>"},
		// before campaign 2 expires
		{now: time.Date(2018, time.October, 22, 07, 00, 0, 0, loc_tokyo), request: ReqFromUser, banner: "<some>CAMPAIGN 2</some>"},
		{now: time.Date(2018, time.October, 21, 18, 00, 0, 0, loc_newyork), request: ReqFromUser, banner: "<some>CAMPAIGN 2</some>"},
		// after campaigns
		{now: time.Date(2018, time.October, 22, 07, 00, 0, 1, loc_tokyo), request: ReqFromUser, banner: ""},
		{now: time.Date(2018, time.October, 21, 18, 00, 0, 1, loc_newyork), request: ReqFromUser, banner: ""},

		// request from administrator
		{now: time.Date(2018, time.September, 30, 23, 59, 59, 0, loc_tokyo), request: ReqFromAdmin1, banner: "<some>CAMPAIGN 1</some>"},
		{now: time.Date(2018, time.September, 30, 23, 59, 59, 0, loc_tokyo), request: ReqFromAdmin2, banner: "<some>CAMPAIGN 1</some>"},
		{now: time.Date(2018, time.September, 30, 10, 59, 59, 0, loc_newyork), request: ReqFromUser, banner: "<some>CAMPAIGN 1</some>"},
	}

	for i, c := range tests {
		t.Logf("test case:%v\tnow:%v", i, c.now)
		now = func() time.Time { return c.now }
		banner, err := Banner(c.request)
		if err != nil {
			t.Fatalf("Banner func failed:%v\n", err)
		}
		//fmt.Printf("now=%v banner=%v\n", now(), banner)
		if c.banner != banner {
			t.Errorf("case:%v received: %v - expected: %v - case: %v", i, banner, c.banner, c)
		}
	}
	truncate()
}

/*
func TestIpAddress(t *testing.T) {
	// set campaigns case 1
	setCampaignsCase1(t)

	loc_tokyo, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatalf("LoadLocation loc1 failed:%v\n", err)
	}
	loc_newyork, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatalf("LoadLocation loc1 failed:%v\n", err)
	}

}
*/

func BenchmarkBanner(b *testing.B) {
	benchBanner(b, 100)
	benchBanner(b, 10000)
	benchBanner(b, 100000)
}

func benchBanner(b *testing.B, camps int) {
	randPeriod := func() (start, end time.Time) {
		start = time.Date(2018, time.October, rand.Intn(30), rand.Intn(24), 0, 0, 0, time.UTC)
		end = start.Add(time.Duration(rand.Intn(30*24)) * time.Hour)
		return
	}
	randTime := func() time.Time {
		return time.Date(2018, time.October, rand.Intn(60), rand.Intn(24), rand.Intn(60), 0, 0, time.UTC)
	}
	b.Run(fmt.Sprintf("%vCampaigns", camps), func(b *testing.B) {
		for i := 0; i < camps; i++ {
			from, to := randPeriod()
			err := NewCampaign(fmt.Sprintf("Camp%v", i), from, to, fmt.Sprintf("<some>Camp %v</some>", i)).Add()
			if err != nil {
				b.Fatalf("add campaign failed: %v", err)
			}
		}
		ReqFromUser, err := http.NewRequest("GET", "dummy", nil)
		if err != nil {
			b.Fatalf("NewRequest error:%v\n", err)
		}
		now = func() time.Time { return randTime() }
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			banner, err := Banner(ReqFromUser)
			if err != nil {
				b.Fatalf("Banner func failed:%v\n", err)
			}
			_ = banner
		}
	})
}
