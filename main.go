package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/k0kubun/pp"
	"github.com/ory/hydra/sdk/go/hydra"
	"github.com/ory/hydra/sdk/go/hydra/swagger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	clientID     = "my-client"
	clientSecret = "secret"
)

type transporter struct {
	*http.Transport
	FakeTLSTermination bool
}

func main() {
	start := time.Now()
	// getToken()
	introspectOAuth2Token()
	log.Println("%v", time.Since(start))
	// cli := getClient()
	// _ = cli
	// pp.Println(cli)
	// listClient()
	// createClient()
	// token := genToken(cli)
	// pp.Println(token)
}

func introspectOAuth2Token() {
	sdk := hydraNew()
	introspect, _, err := sdk.IntrospectOAuth2Token("kft0xeAJ1WK0NO2kx3bvL2AtYy3iUzULAu8OwmCnJyM.INE-H5jzxGDjzAPyOSA4Nmgrg5r9ZBSCO7ofbuYZeBE", "")
	if err != nil {
		log.Fatal("ERROR", err)
	}
	pp.Println(introspect)
	// return cli

}

func getToken() {
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{
		Transport: &transporter{
			FakeTLSTermination: false,
			Transport:          &http.Transport{},
		},
	})
	oauthConfig := clientcredentials.Config{
		ClientID:     "cdp-client-02",
		ClientSecret: clientSecret,
		TokenURL:     "http://localhost:4444/oauth2/token",
	}
	t, err := oauthConfig.Token(ctx)
	if err != nil {
		log.Fatal("ERROR", err)
	}
	pp.Println(t)
}

func getClient() *swagger.OAuth2Client {
	sdk := hydraNew()
	cli, _, err := sdk.GetOAuth2Client("cdp-client-01")
	if err != nil {
		log.Fatal("ERROR", err)
	}
	return cli
}

func listClient() {
	sdk := hydraNew()
	clients, _, err := sdk.ListOAuth2Clients(100, 0)
	if err != nil {
		log.Fatal("ERROR", err)
	}
	pp.Println(clients)
}

func createClient() {
	sdk := hydraNew()
	client, _, err := sdk.CreateOAuth2Client(swagger.OAuth2Client{
		ClientId:      "cdp-client-02",
		ClientName:    "Sample Client",
		ClientSecret:  clientSecret,
		GrantTypes:    []string{"client_credentials"},
		ResponseTypes: []string{"token"},
		Scope:         "openid,offline,hydra",
	})
	if err != nil {
		log.Fatal("ERROR", err)
	}
	pp.Println(client)
}

func hydraNew() *hydra.CodeGenSDK {
	sdk, err := hydra.NewSDK(&hydra.Configuration{
		AdminURL:     "http://localhost:4445",
		PublicURL:    "http://localhost:4444",
		ClientID:     clientID,
		ClientSecret: clientSecret,
	})
	if err != nil {
		log.Fatal("ERROR", err)
	}

	return sdk
}
