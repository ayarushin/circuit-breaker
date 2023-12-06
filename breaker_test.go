package main

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBreaker_Success(t *testing.T) {
	mockCircuit := func(ctx context.Context) (string, error) {
		return "success", nil
	}

	breaker := Breaker(mockCircuit, 3)
	response, err := breaker(context.Background())

	assert.Equal(t, "success", response)
	assert.Nil(t, err)
}

func TestBreaker_FailureOpen(t *testing.T) {
	mockCircuit := func(ctx context.Context) (string, error) {
		return "", errors.New("service unreachable")
	}

	breaker := Breaker(mockCircuit, 3)

	response, err := breaker(context.Background())

	assert.Equal(t, "", response)
	assert.Equal(t, errors.New("service unreachable"), err)
}

func TestBreaker_AttemptWhileOpen(t *testing.T) {
	mockCircuit := func(ctx context.Context) (string, error) {
		return "", errors.New("service unreachable")
	}

	breaker := Breaker(mockCircuit, 3)

	response, err := breaker(context.Background())

	assert.Equal(t, "", response)
	assert.Equal(t, errors.New("service unreachable"), err)

	time.Sleep(time.Second * 2) // Simulate waiting for the retry interval

	response, err = breaker(context.Background())

	assert.Equal(t, "", response)
	assert.Equal(t, errors.New("service unreachable"), err)
}
