// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2018 Datadog, Inc.

package check

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// FIXTURE
type TestCheck struct{}

func (c *TestCheck) String() string                            { return "TestCheck" }
func (c *TestCheck) Stop()                                     {}
func (c *TestCheck) Configure(ConfigData, ConfigData) error    { return nil }
func (c *TestCheck) Interval() time.Duration                   { return 1 }
func (c *TestCheck) Run() error                                { return nil }
func (c *TestCheck) ID() ID                                    { return ID(c.String()) }
func (c *TestCheck) GetWarnings() []error                      { return []error{} }
func (c *TestCheck) GetMetricStats() (map[string]int64, error) { return make(map[string]int64), nil }

func TestIdentify(t *testing.T) {
	testCheck := &TestCheck{}

	instance1 := ConfigData("key1:value1\nkey2:value2")
	initConfig1 := ConfigData("key:value")
	instance2 := instance1
	initConfig2 := initConfig1
	assert.Equal(t, Identify(testCheck, instance1, initConfig1), Identify(testCheck, instance2, initConfig2))

	instance3 := ConfigData("key1:value1\nkey2:value3")
	initConfig3 := ConfigData("key:value")
	assert.NotEqual(t, Identify(testCheck, instance1, initConfig1), Identify(testCheck, instance3, initConfig3))
}

func TestIDToCheckName(t *testing.T) {
	testCases := []struct {
		in  string
		out string
	}{
		{
			in:  "valid:9505c316b4e4a028",
			out: "valid",
		},
		{
			in:  "",
			out: "",
		},
		{
			in:  "nocolon",
			out: "nocolon",
		},
		{
			in:  "nohash:",
			out: "nohash",
		},
		{
			in:  "multiple:colon:9505c316b4e4a028",
			out: "multiple",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case %d: %s", i, tc.in), func(t *testing.T) {
			assert.Equal(t, tc.out, IDToCheckName(ID(tc.in)))
		})
	}
}
