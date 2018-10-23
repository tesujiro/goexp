package main

import (
	"fmt"
	"time"
)

type Period struct {
	from time.Time
	to   time.Time
	TZ   time.Location
}

func NewPeriod(from, to time.Time) *Period {
	return &Period{from: from, to: to}
}

func (p *Period) duringThePeriod(t time.Time) bool {
	return (p.from.Before(t) || p.from.Equal(t)) && (p.to.After(t) || p.to.Equal(t))
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
