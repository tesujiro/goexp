package promotion

/*
Package promotion provides html banners for Mercari app.

What is Banner function

Banner function provides a banner.
	A banner is expired when the display period is over.
	A banner’s display period is the duration the banner is active on the screen.
	A banner is active during the display period.
	Requests from IP address (10.0.0.1, 10.0.0.2) can display a banner even if current time is before the display peiod.
	When more than one banner are active, return the banner with the earlier expiration.

How to set display period

AddCamapaign function adds a new campaign with a banner and a period to display it.

*/

import (
	"fmt"
	"net"
	"net/http"
	"sort"
	"sync"
	"time"
)

type campaign struct {
	name      string
	startAt   time.Time
	expiresAt time.Time
	banner    string
}

// a global object to store all the campaigns
var campaigns []campaign

// for lock while adding new campaigns
var mu *sync.Mutex

func init() {
	// Initialize the campaigns object.
	// If campaigns are stored in the persistence devices, retreive them from devices here.
	campaigns = make([]campaign, 0)

	// Mutex for lock campaigns
	mu = new(sync.Mutex)
}

// Administrators IP addresses
var ipAddrList = []net.IP{
	net.IPv4(10, 0, 0, 1),
	net.IPv4(10, 0, 0, 2),
}

// Add a new campaign with a banner and a period to display it.
func AddCampaign(name string, start, end time.Time, banner string) error {

	c := &campaign{
		name:      name,
		startAt:   start,
		expiresAt: end,
		banner:    banner,
	}

	// TODO: Check current time, if expire time is older than current time, return error
	if c.startAt.After(c.expiresAt) {
		return fmt.Errorf("error: start date is after expire date.")
	}
	if c.banner == "" {
		return fmt.Errorf("error: banner is not set.")
	}

	// Campaigns are ordered by expiresAt, so at first search the new expire date in the campaigns.
	i := sort.Search(len(campaigns), func(i int) bool { return campaigns[i].expiresAt.After(c.expiresAt) })

	mu.Lock()

	// Insert new campaign at the location.
	campaigns = append(campaigns, campaign{})
	copy(campaigns[i+1:], campaigns[i:])
	campaigns[i] = *c

	mu.Unlock()

	return nil
}

// The function to get current time can be changed for testing.
var nowFunc = time.Now

// TODO: COMMENT
//
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
			for _, c := range campaigns {  // low performance
				if (isAdmin && now.Before(c.expiresAt)) || c.duringThePeriod(now) {
					if expiresAt.IsZero() || c.expiresAt.Before(expiresAt) {
						expiresAt = c.expiresAt
						banner = c.banner
					}
				}
			}
		return banner, nil
	*/
	// Search the expired time. Campaigns are ordered by expired date.
	i := sort.Search(len(campaigns), func(i int) bool { return campaigns[i].expiresAt.After(now) || campaigns[i].expiresAt.Equal(now) })
	for j := i; j < len(campaigns); j++ {
		if isAdmin || campaigns[i].startAt.Before(now) || campaigns[i].startAt.Equal(now) {
			return campaigns[i].banner, nil
		}
	}
	return "", nil
}

// Returns t is in the campaign period.
func (c *campaign) duringThePeriod(t time.Time) bool {
	return (c.startAt.Before(t) || c.startAt.Equal(t)) && (c.expiresAt.After(t) || c.expiresAt.Equal(t))
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

// FromRequest extracts the user IP address from req, if present.
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

// Check administrator IP address
func isAdminIP(ip net.IP) bool {
	for i := 0; i < len(ipAddrList); i++ {
		if ipAddrList[i].Equal(ip) {
			return true
		}
	}
	return false
}
