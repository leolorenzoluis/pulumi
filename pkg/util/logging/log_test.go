// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitLogging(t *testing.T) {
	// Just ensure we can initialize logging (and reset it afterwards).
	prevLog := LogToStderr
	prevV := Verbose
	prevFlow := LogFlow
	InitLogging(true, 9, true)
	InitLogging(prevLog, prevV, prevFlow)
	assert.Equal(t, prevLog, LogToStderr)
	assert.Equal(t, prevV, Verbose)
	assert.Equal(t, prevFlow, LogFlow)
}

func TestFilter(t *testing.T) {
	filter1 := CreateFilter([]string{"secret1", "secret2"}, "[secret]")
	msg1 := filter1.Filter("These are my secrets: secret1, secret2, secret3, secret10")
	assert.Equal(t, msg1, "These are my secrets: [secret], [secret], secret3, [secret]0")

	// Ensure htat special characters don't screw up the regex we create
	filter2 := CreateFilter([]string{"secret.*", "secre[t]3"}, "[creds]")
	msg2 := filter2.Filter("These are my secrets: secret1, secret2, secret3, secret.*, secre[t]3")
	assert.Equal(t, msg2, "These are my secrets: secret1, secret2, secret3, [creds], [creds]")
}
