package promotion

import (
	"fmt"
	"net/http"
	"time"
)

type Campaign struct {
	name      string
	startAt   time.Time
	expiresAt time.Time
	banner    string
}

var campaigns []Campaign

func init() {
	campaigns = make([]Campaign, 0)
}

/*
type by func(c1, c2 *Campagin) bool
func (b by) Sort(cs []Campagin) {
}
*/

func NewCampaign(name string, start, end time.Time, banner string) *Campaign {
	return &Campaign{
		name:      name,
		startAt:   start,
		expiresAt: end,
		banner:    banner,
	}
}

func (c *Campaign) Add() error {
	if c.startAt.After(c.expiresAt) {
		return fmt.Errorf("error: start date is after expire date.")
	}
	if c.banner == "" {
		return fmt.Errorf("error: banner is not set.")
	}

	campaigns = append(campaigns, *c)
	return nil
}

func (c *Campaign) duringThePeriod(t time.Time) bool {
	return (c.startAt.Before(t) || c.startAt.Equal(t)) && (c.expiresAt.After(t) || c.expiresAt.Equal(t))
}

func truncate() {
	campaigns = make([]Campaign, 0)
}

func list() []Campaign {
	cs := make([]Campaign, 0)
	for _, c := range campaigns {
		cs = append(cs, c)
		fmt.Printf("campaign:%v\n", c)
	}
	return cs
}

var now = time.Now

func Banner(r *http.Request) (string, error) {
	var banner string
	var expiresAt time.Time
	n := now()
	for _, c := range campaigns {
		if c.duringThePeriod(n) {
			if expiresAt.IsZero() || c.expiresAt.Before(expiresAt) {
				expiresAt = c.expiresAt
				banner = c.banner
			}
		}
	}
	return banner, nil
}
