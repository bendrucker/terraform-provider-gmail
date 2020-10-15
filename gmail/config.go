package gmail

import (
	"context"
	"encoding/json"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type Config struct {
	client *oauth2.Config
	token  *oauth2.Token
}

func (c *Config) Client(credentials []byte) error {
	client, err := google.ConfigFromJSON(credentials, gmail.GmailSettingsBasicScope)
	c.client = client
	return err
}

func (c *Config) Token(credentials []byte) error {
	token := &oauth2.Token{}

	if err := json.Unmarshal(credentials, token); err != nil {
		return err
	}

	c.token = token
	return nil
}

func (c *Config) NewService(ctx context.Context) (*gmail.Service, error) {
	return gmail.NewService(ctx, option.WithTokenSource(c.client.TokenSource(ctx, c.token)))
}
