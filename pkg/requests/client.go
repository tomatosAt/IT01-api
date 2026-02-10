package requests

import (
	"bytes"
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"net/http/httptrace"
	"time"

	"github.com/sirupsen/logrus"
)

type HttpClient struct {
	log *logrus.Entry
}

func NewHttpClient(logger *logrus.Entry) *HttpClient {
	return &HttpClient{
		log: logger.Dup().WithField("package", packageName),
	}
}

func (c *HttpClient) Request(
	ctx context.Context,
	method string,
	url string,
	headers map[string]string,
	body io.Reader,
	timeout int,
) (*Response, error) {
	return c.RequestWithTLSConfig(ctx, method, url, headers, body, timeout, nil)
}

func (c *HttpClient) RequestWithTLSConfig(
	ctx context.Context,
	method string,
	url string,
	headers map[string]string,
	body io.Reader,
	timeout int,
	tlsCfg *tls.Config,
) (*Response, error) {

	start := time.Now()

	if timeout == 0 {
		timeout = Timeout
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if tlsCfg == nil {
		tlsCfg = &DefaultTLSConfig
	}

	var remoteAddr string
	trace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) {
			if info.Conn != nil {
				remoteAddr = info.Conn.RemoteAddr().String()
			}
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsCfg,
		},
		Timeout: time.Duration(timeout) * time.Second,
	}

	c.log.WithFields(logrus.Fields{
		"method": method,
		"url":    url,
	}).Info("HTTP request started")

	resp, err := client.Do(req)
	if err != nil {
		c.log.WithError(err).Error("HTTP request failed")
		return nil, err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(resp.Body)

	c.log.WithFields(logrus.Fields{
		"method":     method,
		"url":        url,
		"status":     resp.StatusCode,
		"remote":     remoteAddr,
		"latency_ms": time.Since(start).Milliseconds(),
	}).Info("HTTP request finished")

	return &Response{
		Code:   resp.StatusCode,
		Status: resp.Status,
		Body:   buf.Bytes(),
		Header: resp.Header,
	}, nil
}

func (c *HttpClient) Get(ctx context.Context, url string, headers map[string]string, body io.Reader, timeout int) (*Response, error) {
	return c.Request(ctx, http.MethodGet, url, headers, body, timeout)
}

func (c *HttpClient) Post(ctx context.Context, url string, headers map[string]string, body io.Reader, timeout int) (*Response, error) {
	return c.Request(ctx, http.MethodPost, url, headers, body, timeout)
}

func (c *HttpClient) Put(ctx context.Context, url string, headers map[string]string, body io.Reader, timeout int) (*Response, error) {
	return c.Request(ctx, http.MethodPut, url, headers, body, timeout)
}

func (c *HttpClient) Delete(ctx context.Context, url string, headers map[string]string, body io.Reader, timeout int) (*Response, error) {
	return c.Request(ctx, http.MethodDelete, url, headers, body, timeout)
}
