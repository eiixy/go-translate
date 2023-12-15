package translate

import (
	"context"
	"encoding/json"
	"golang.org/x/time/rate"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const translateURL = "https://translate.googleapis.com/translate_a/single"

const (
	LangZhCn = "zh-CN"
)

type Client struct {
	limiter    *rate.Limiter
	targetLang string
}

// NewClient new a Client
func NewClient(opts ...Option) *Client {
	c := Client{
		targetLang: LangZhCn,
		limiter:    rate.NewLimiter(rate.Every(time.Second*10), 5),
	}
	for _, opt := range opts {
		opt(&c)
	}
	return &c
}

// Translates translate texts
func (r Client) Translates(texts []string) ([]string, error) {
	return r.TranslatesWithTargetLang(texts, r.targetLang)
}

// Translate translate text
func (r Client) Translate(text string) (string, error) {
	return r.TranslateWithTargetLang(text, r.targetLang)
}

type kv struct {
	key string
	val string
}

// TranslatesWithTargetLang translate texts with target lang
func (r Client) TranslatesWithTargetLang(texts []string, targetLang string) ([]string, error) {
	if len(texts) == 0 {
		return []string{}, nil
	}
	_ = r.limiter.Wait(context.Background())
	var contents []string
	for _, item := range texts {
		if item != "" {
			contents = append(contents, item)
		}
	}
	result, err := r.request(strings.Join(contents, "\n"), targetLang)
	if err != nil {
		return nil, err
	}

	var trans []kv
	if len(result) > 0 {
		if items, ok := result[0].([]any); ok {
			for _, item := range items {
				if values, ok := item.([]any); ok && len(values) >= 2 {
					trans = append(trans, kv{
						strings.Trim(values[1].(string), " \n"),
						strings.Trim(values[0].(string), " \n"),
					})
				}
			}
		}
	}

	for i, text := range texts {
		var temp []kv
		for _, item := range trans {
			if strings.Contains(text, item.key) {
				text = strings.Replace(text, item.key, item.val, 1)
			} else {
				temp = append(temp, item)
			}
		}
		trans = temp
		texts[i] = text
	}
	return texts, nil
}

// TranslateWithTargetLang translate text with target lang
func (r Client) TranslateWithTargetLang(text, targetLang string) (string, error) {
	_ = r.limiter.Wait(context.Background())
	result, err := r.request(text, targetLang)
	if err != nil {
		return "", err
	}
	var content string
	if len(result) > 0 {
		if items, ok := result[0].([]any); ok {
			for _, item := range items {
				if values, ok := item.([]any); ok && len(values) >= 2 {
					content += values[0].(string)
				}
			}
		}
	}
	return content, nil
}

func (r Client) request(text string, targetLang string) ([]any, error) {
	data := url.Values{}
	data.Add("client", "gtx")
	data.Add("sl", "auto")
	data.Add("tl", targetLang)
	data.Add("dt", "t")
	data.Add("q", text)

	resp, err := http.Post(translateURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result []any
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
