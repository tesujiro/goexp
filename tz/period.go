package main

import (
	"fmt"
	"time"
)

// for testing purpose now func can be altered inside this package
var now func() time.Time = time.Now

type Period struct {
	from time.Time
	to   time.Time
	TZ   time.Location
}

func NewPeriod(from, to time.Time) *Period {
	return &Period{from: from, to: to}
}

func (p *Period) duringThePeriod() bool {
	now:=now()
	return (p.from.Before(now) || p.from.Equal(now)) && (p.to.After(now) || p.to.Equal(now))
}

func main() {
	//p := NewPeriod(time.Now(), time.Now())
	//fmt.Printf("p=%#v\n", p)
	now := time.Now()
	//oneHourBefor := now.Add(-time.Hour)
	//oneHourAfter := now.Add(time.Hour)
	fmt.Printf("now=%v\n", now)
	loc, _ := time.LoadLocation("UTC")
	fmt.Printf("now=%v\n", now.In(loc))
}
