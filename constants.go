package main

import (
	"math/rand"
	"time"
)

// TotalOperations is the total number of distinct operation types
const (
	TotalOperations = 10
	NumOperations   = 3
	NumSteps        = 25
)

var (
	numbers = []string{"x", "y", "t", "n"}
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)
