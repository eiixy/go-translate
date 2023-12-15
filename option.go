package translate

import "golang.org/x/time/rate"

type Option func(client *Client)

func WithTargetLang(lang string) Option {
	return func(client *Client) {
		client.targetLang = lang
	}
}

func WithLimiter(limiter *rate.Limiter) Option {
	return func(client *Client) {
		client.limiter = limiter
	}
}
