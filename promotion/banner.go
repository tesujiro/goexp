package promotion

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

type Campaign struct {
	name      string
	startAt   time.Time
	expiresAt time.Time
	banner    string
}

var campaigns []Campaign
var mu *sync.Mutex

func init() {
	campaigns = make([]Campaign, 0)
	mu = new(sync.Mutex)
}

var ipAddrList = []net.IP{
	net.IPv4(10, 0, 0, 1),
	net.IPv4(10, 0, 0, 2),
}

func isAdminIP(ip net.IP) bool {
	for i := 0; i < len(ipAddrList); i++ {
		if ipAddrList[i].Equal(ip) {
			return true
		}
	}
	return false
}

// FromRequest extracts the user IP address from req, if present.
// https://blog.golang.org/context/userip/userip.go
func getIPFromRequest(req *http.Request) (net.IP, error) {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
	}
	return userIP, nil
}

func isFromAdmin(req *http.Request) (bool, error) {
	ip, err := getIPFromRequest(req)
	if err != nil {
		return false, err
	}
	if isAdminIP(ip) {
		return true, nil
	}
	return false, nil
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

	mu.Lock()
	campaigns = append(campaigns, *c)
	mu.Unlock()

	return nil
}

func (c *Campaign) duringThePeriod(t time.Time) bool {
	return (c.startAt.Before(t) || c.startAt.Equal(t)) && (c.expiresAt.After(t) || c.expiresAt.Equal(t))
}

// truncate campaigns for testing
func truncate() {
	campaigns = make([]Campaign, 0)
}

/*
func list() []Campaign {
	cs := make([]Campaign, 0)
	for _, c := range campaigns {
		cs = append(cs, c)
		fmt.Printf("campaign:%v\n", c)
	}
	return cs
}
*/

// number of campaigns for testing
func countCampaigns() int {
	return len(campaigns)
}

var now = time.Now

func Banner(r *http.Request) (string, error) {
	var banner string
	var expiresAt time.Time

	// Check ip address
	isAdmin, err := isFromAdmin(r)
	if err != nil {
		return "", fmt.Errorf("cannot check address:%v", err)
	}

	// Check periods of campaigns
	n := now()
	for _, c := range campaigns {
		if isAdmin || c.duringThePeriod(n) {
			if expiresAt.IsZero() || c.expiresAt.Before(expiresAt) {
				expiresAt = c.expiresAt
				banner = c.banner
			}
		}
	}
	return banner, nil
}
