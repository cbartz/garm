// Copyright 2025 Cloudbase Solutions SRL
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

package sql

import (
	"os"
	"testing"

	"github.com/cloudbase/garm/config"
	garmTesting "github.com/cloudbase/garm/internal/testing"
)

const (
	wrongPassphrase = "wrong-passphrase"
	webhookSecret   = "webhook-secret"
	falseString     = "false"
)

// testDBConfig returns the database config for a suite test case.
// Returns a PostgreSQL config when GARM_TEST_POSTGRES_DSN is set,
// otherwise a fresh per-test SQLite database.
func testDBConfig(t *testing.T) config.Database {
	if os.Getenv("GARM_TEST_POSTGRES_DSN") != "" {
		return garmTesting.GetTestPostgresDBConfig(t)
	}
	return garmTesting.GetTestSqliteDBConfig(t)
}
