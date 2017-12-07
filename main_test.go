package main_test

import (
	"testing"

	"github.com/tkitsunai/commango-slackbot/invoke"
)

func Test(t *testing.T) {
	invoke.CallCommand([]string{"tail", "10", "/data/log"})
}
