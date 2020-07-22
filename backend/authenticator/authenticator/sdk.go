package authenticator

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	name = "authenticator"
)

type Interface interface {
	SignUp(SignUpRequest) error
	SignOut(SignOutRequest) error
	SignIn(SignInRequest) (SignInResponse, error)
	ValidateToken(string) (string, string, error)
}

func NewAuthenticatorSDK(host string, httpClient *http.Client) Interface {
	return &authenticatorSDK{
		host:   host,
		client: httpClient,
	}
}

var (
	NotAuthorized        = errors.New("notAuthorized")
	UnexpectedStatusCode = errors.New("unexpectedStatusCode")
)

type authenticatorSDK struct {
	host   string
	client *http.Client
}

func (c *authenticatorSDK) SignUp(sur SignUpRequest) error {
	body, err := json.Marshal(&sur)
	if err != nil {
		return err
	}
	r, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/authenticator/v1/user", c.host), bytes.NewReader(body))
	if err != nil {
		return err
	}
	res, err := c.client.Do(r)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusNoContent {
		body, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		if res.StatusCode != http.StatusNoContent {
			return fmt.Errorf("%w: %d: %s", UnexpectedStatusCode, res.StatusCode, string(body))
		}
	}
	return nil
}

func (c *authenticatorSDK) SignOut(sor SignOutRequest) error {
	body, err := json.Marshal(&sor)
	if err != nil {
		return err
	}
	r, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/authenticator/v1/user/session", c.host), bytes.NewReader(body))
	if err != nil {
		return err
	}
	res, err := c.client.Do(r)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusNoContent {
		body, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		if res.StatusCode != http.StatusNoContent {
			return fmt.Errorf("%w: %d: %s", UnexpectedStatusCode, res.StatusCode, string(body))
		}
	}
	return nil
}

func (c *authenticatorSDK) SignIn(sir SignInRequest) (SignInResponse, error) {
	body, err := json.Marshal(&sir)
	if err != nil {
		return SignInResponse{}, err
	}
	r, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/authenticator/v1/user/session", c.host), bytes.NewReader(body))
	if err != nil {
		return SignInResponse{}, err
	}
	res, err := c.client.Do(r)
	if err != nil {
		return SignInResponse{}, err
	}
	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return SignInResponse{}, err
	}
	if res.StatusCode != http.StatusNoContent {
		return SignInResponse{}, fmt.Errorf("%w: %d: %s", UnexpectedStatusCode, res.StatusCode, string(body))
	}
	sires := SignInResponse{}
	err = json.Unmarshal(body, &sires)
	if err != nil {
		return SignInResponse{}, err
	}
	return sires, nil
}

func (c *authenticatorSDK) ValidateToken(token string) (string, string, error) {
	r, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/authenticator/v1/user/session", c.host), nil)
	if err != nil {
		return "", "", err
	}
	res, err := c.client.Do(r)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()
	switch res.StatusCode {
	case http.StatusNoContent:
	case http.StatusForbidden:
		return "", "", NotAuthorized
	default:
		return "", "", fmt.Errorf("%w: got %d", UnexpectedStatusCode, res.StatusCode)
	}
	return r.Header.Get(HttpUserIDHeader), r.Header.Get(HttpInternalUserIDHeader), nil
}
