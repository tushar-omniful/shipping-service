package utils

import (
	"context"
	"github.com/omniful/go_commons/config"
)

func GetNameSpace(ctx context.Context) string {
	return config.GetString(ctx, "service.name")
}

func ContainsEmptyString(slice []string) bool {
	for _, s := range slice {
		if len(s) == 0 {
			return true
		}
	}
	return false
}

func ContainsKeyInMap[K comparable, V any](key K, m map[K]V) bool {
	_, ok := m[key]
	return ok
}

func Contains[T comparable](array []T, value T) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
}
