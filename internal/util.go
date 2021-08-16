package internal

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/drone/drone-go/drone"

	"golang.org/x/oauth2"
)

func NewClient(token, server string) (drone.Client, error) {
	server = strings.TrimRight(server, "/")

	// if no server url is provided we can default
	// to the hosted Drone service.
	if len(server) == 0 {
		return nil, fmt.Errorf("error: you must provide the Drone server address")
	}
	if len(token) == 0 {
		return nil, fmt.Errorf("error: you must provide your Drone access token")
	}

	config := new(oauth2.Config)
	auther := config.Client(
		oauth2.NoContext,
		&oauth2.Token{
			AccessToken: token,
		},
	)

	auther.CheckRedirect = func(*http.Request, []*http.Request) error {
		return fmt.Errorf("attempting to redirect the requests, did you configure the correct drone server address")
	}
	return drone.NewClient(server, auther), nil
}
