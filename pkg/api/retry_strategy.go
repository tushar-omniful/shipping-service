package api

import (
	"net/http"
	"time"
)

type RetryStrategy interface {
	GetMaxRetries() int
	GetRetryStatusCodes() []int
	ShouldRetry(resp *http.Response) bool
	GetRetryDelay(retryCount int) time.Duration
}

type DefaultRetryStrategy struct {
	MaxRetries            int
	RetryStatusCodes      []int
	BackoffDuration       time.Duration
	ServerDownStatusCodes []int
	RetryNumber           int
}

func (s *DefaultRetryStrategy) GetMaxRetries() int {
	if s.MaxRetries == 0 {
		return 1
	}

	return s.MaxRetries
}

func (s *DefaultRetryStrategy) GetRetryNumber() int {
	return s.RetryNumber
}

func (s *DefaultRetryStrategy) GetRetryStatusCodes() []int {
	return s.RetryStatusCodes
}

func (s *DefaultRetryStrategy) ShouldRetry(resp *http.Response) bool {
	if resp == nil {
		return false
	}

	return contains(s.RetryStatusCodes, resp.StatusCode) &&
		!(s.GetRetryNumber() > s.GetMaxRetries()) &&
		!s.IsServerDown(resp)
}

func (s *DefaultRetryStrategy) GetRetryDelay(retryCount int) time.Duration {
	return s.BackoffDuration * time.Duration(1<<retryCount)
}

func (s *DefaultRetryStrategy) IsServerDown(resp *http.Response) bool {
	if resp == nil {
		return false
	}

	return contains(s.ServerDownStatusCodes, resp.StatusCode)
}

// Helper function to check if a value is present in an int slice
func contains(s []int, e int) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
