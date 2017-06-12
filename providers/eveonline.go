package providers

import (
	"log"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bitly/oauth2_proxy/api"
)

type EveOnlineProvider struct {
	*ProviderData
}

func NewEveOnlineProvider(p *ProviderData) *EveOnlineProvider {

	log.Printf("Creating Eve Online provider")

	p.ProviderName = "EveOnline"
	if p.LoginURL == nil || p.LoginURL.String() == "" {
		p.LoginURL = &url.URL{
			Scheme: "https",
			Host:   "login.eveonline.com",
			Path:   "/oauth/authorize",
		}
	}
	if p.RedeemURL == nil || p.RedeemURL.String() == "" {
		p.RedeemURL = &url.URL{
			Scheme: "https",
			Host:   "login.eveonline.com",
			Path:   "/oauth/token",
		}
	}
	if p.ValidateURL == nil || p.ValidateURL.String() == "" {
		p.ValidateURL = &url.URL{
			Scheme: "https",
			Host:   "login.eveonline.com",
			Path:   "/oauth/verify",
		}
	}
	if p.Scope == "" {
		p.Scope = ""
	}
	return &EveOnlineProvider{ProviderData: p}
}

func (p *EveOnlineProvider) GetEmailAddress(s *SessionState) (string, error) {

    req, err := http.NewRequest("GET", p.ValidateURL.String(), nil)
    header := make(http.Header)
	header.Set("Authorization", fmt.Sprintf("Bearer %s", s.AccessToken))
    req.Header = header
	if err != nil {
		log.Printf("failed building request %s", err)
		return "", err
	}

	json, err := api.Request(req)
	if err != nil {
		log.Printf("failed making request %s", err)
		return "", err
	}

    charname, err := json.Get("CharacterName").String()
    if err != nil {
        return "", err
    }

	return charname, nil
}
