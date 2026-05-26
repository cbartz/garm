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
// When dbCfgFn is set (non-nil), it returns that backend's config.
// Otherwise it creates a fresh SQLite database per test case, preserving
// the original per-test isolation for the default SQLite backend.
func testDBConfig(dbCfgFn func(*testing.T) config.Database, t *testing.T) config.Database {
	if dbCfgFn != nil {
		return dbCfgFn(t)
	}
	return garmTesting.GetTestSqliteDBConfig(t)
}

// pgTestDBConfig returns a PostgreSQL database config read from GARM_TEST_POSTGRES_DSN.
// Returns false if the environment variable is not set, so callers can skip the
// PostgreSQL run without calling t.Skip() from an outer test function.
func pgTestDBConfig(t *testing.T) (config.Database, bool) {
	t.Helper()
	if os.Getenv("GARM_TEST_POSTGRES_DSN") == "" {
		return config.Database{}, false
	}
	return garmTesting.GetTestPostgresDBConfig(t), true
}
