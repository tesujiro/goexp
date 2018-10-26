package promotion

import (
	"fmt"
	"net"
	"net/http"
	"sort"
	"sync"
	"time"
)

type Campaign struct {
	name      string
	startAt   time.Time
	expiresAt time.Time
	banner    string
}

type camps []Campaign

// all the campaigns
var campaigns camps

// for lock while adding new campaigns
var mu *sync.Mutex

func init() {
	campaigns = make(camps, 0)
	mu = new(sync.Mutex)
}

// Administrators IP addresses
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

// Check the request from administrator(s)
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
func (a camps) Len() int           { return len(a) }
func (a camps) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a camps) Less(i, j int) bool { return a[i].expiresAt.Before(a[j].expiresAt) }
*/

// Return a new Campaign struct
func NewCampaign(name string, start, end time.Time, banner string) *Campaign {
	return &Campaign{
		name:      name,
		startAt:   start,
		expiresAt: end,
		banner:    banner,
	}
}

func addCampaign(c Campaign) camps {
	//return append(campaigns, c)
	// Binary Search
	i := sort.Search(len(campaigns), func(i int) bool { return campaigns[i].expiresAt.After(c.expiresAt) })
	campaigns = append(campaigns, Campaign{})
	copy(campaigns[i+1:], campaigns[i:])
	campaigns[i] = c
	return campaigns
}

// TODO: Campaign -> AddCampaign  ????
func (c *Campaign) Add() error {
	if c.startAt.After(c.expiresAt) {
		return fmt.Errorf("error: start date is after expire date.")
	}
	if c.banner == "" {
		return fmt.Errorf("error: banner is not set.")
	}

	mu.Lock()
	campaigns = addCampaign(*c)
	mu.Unlock()

	return nil
}

// truncate campaigns for testing
// TODO: truncate -> truncateCampaigns
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

// Return number of campaigns for testing
func countCampaigns() int {
	return len(campaigns)
}

// The function to get current time can be changed for testing.
var nowFunc = time.Now

func (c *Campaign) duringThePeriod(t time.Time) bool {
	return (c.startAt.Before(t) || c.startAt.Equal(t)) && (c.expiresAt.After(t) || c.expiresAt.Equal(t))
}

// TODO: COMMENT
func Banner(r *http.Request) (string, error) {
	// Check ip address
	isAdmin, err := isFromAdmin(r)
	if err != nil {
		return "", fmt.Errorf("cannot check address:%v", err)
	}

	// Check periods of campaigns
	now := nowFunc()
	/*
			var banner string
			var expiresAt time.Time
			for _, c := range campaigns {
				if (isAdmin && now.Before(c.expiresAt)) || c.duringThePeriod(now) {
					if expiresAt.IsZero() || c.expiresAt.Before(expiresAt) {
						expiresAt = c.expiresAt
						banner = c.banner
					}
				}
			}
		return banner, nil
	*/
	// Binary Seach for Performance
	i := sort.Search(len(campaigns), func(i int) bool { return campaigns[i].expiresAt.After(now) || campaigns[i].expiresAt.Equal(now) })
	for j := i; j < len(campaigns); j++ {
		if isAdmin || campaigns[i].startAt.Before(now) || campaigns[i].startAt.Equal(now) {
			return campaigns[i].banner, nil
		}
	}
	return "", nil
}
