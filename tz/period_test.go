package main

import (
	"testing"
	"time"
)

func TestPeriod(t *testing.T) {
	loc1, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatalf("LoadLocation loc1 failed:%v\n", err)
	}

	/*
			loc2, err := time.LoadLocation("America/New_York")
			if err != nil {
				t.Fatalf("LoadLocation loc2 failed:%v\n", err)
			}

		utc, err := time.LoadLocation("UTC")
		if err != nil {
			t.Fatalf("LoadLocation utc failed:%v\n", err)
		}
	*/

	now_loc1 := time.Now().In(loc1)
	twoHourBefore_loc1 := now_loc1.Add(-2 * time.Hour)
	oneHourBefore_loc1 := now_loc1.Add(-time.Hour)
	oneHourAfter_loc1 := now_loc1.Add(time.Hour)
	//twoHourAfter_loc1 := now_loc1.Add(2 * time.Hour)

	/*
		now_loc2 := time.Now().In(loc1)
		twoHourBefore_loc2 := now_loc2.Add(-2 * time.Hour)
		oneHourBefore_loc2 := now_loc2.Add(-time.Hour)
		oneHourAfter_loc2 := now_loc2.Add(time.Hour)
		twoHourAfter_loc2 := now_loc2.Add(2 * time.Hour)
	*/

	//camp1_us_from :=time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC)

	cases := []struct {
		location *time.Location
		from, to time.Time
		nowFunc      func() time.Time
		during   bool
	}{
		{location: loc1, nowFunc: time.Now, from: twoHourBefore_loc1, to: oneHourBefore_loc1, during: false},
		{location: loc1, nowFunc: time.Now, from: oneHourBefore_loc1, to: oneHourAfter_loc1, during: true},
	}

	for _, c := range cases {
		now = c.nowFunc
		time.Local = c.location
		p := NewPeriod(c.from, c.to)
		during := p.duringThePeriod()
		if during != c.during {
			t.Errorf("received: %v - expected: %v - case: %v", during, c.during, c)
		}
	}
}
