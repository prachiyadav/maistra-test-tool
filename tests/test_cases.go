// Copyright 2021 Red Hat, Inc.
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

package tests

import (
	"testing"

	"github.com/maistra/maistra-test-tool/pkg/tasks/security/certificate"
	"github.com/maistra/maistra-test-tool/pkg/tasks/traffic"
	"github.com/maistra/maistra-test-tool/pkg/tasks/traffic/egress"
	"github.com/maistra/maistra-test-tool/pkg/tasks/traffic/ingress"

	"github.com/maistra/maistra-test-tool/pkg/ossm"
)

var t = &testing.T{}
var testCases = []testing.InternalTest{
	testing.InternalTest{
		Name: "01",
		F:    traffic.TestRequestRouting,
	},
	testing.InternalTest{
		Name: "02",
		F:    traffic.TestFaultInjection,
	},
	testing.InternalTest{
		Name: "03",
		F:    traffic.TestTrafficShifting,
	},
	testing.InternalTest{
		Name: "04",
		F:    traffic.TestTCPShifting,
	},
	testing.InternalTest{
		Name: "05",
		F:    traffic.TestRequestTimeouts,
	},
	testing.InternalTest{
		Name: "06",
		F:    traffic.TestCircuitBreaking,
	},
	testing.InternalTest{
		Name: "07",
		F:    traffic.TestMirroring,
	},
	testing.InternalTest{
		Name: "08",
		F:    ingress.TestIngressGateways,
	},
	testing.InternalTest{
		Name: "09",
		F:    ingress.TestSecureGateways,
	},
	testing.InternalTest{
		Name: "10",
		F:    ingress.TestIngressWithoutTLS,
	},
	testing.InternalTest{
		Name: "11",
		F:    egress.TestAccessExternalServices,
	},
	testing.InternalTest{
		Name: "12",
		F:    egress.TestEgressTLSOrigination,
	},
	testing.InternalTest{
		Name: "13",
		F:    egress.TestEgressGateways,
	},

	testing.InternalTest{
		Name: "15",
		F:    egress.TestTLSOriginationFileMount,
	},
	testing.InternalTest{
		Name: "16",
		F:    egress.TestEgressWildcard,
	},
	testing.InternalTest{
		Name: "17",
		F:    certificate.TestExternalCert,
	},

	testing.InternalTest{
		Name: "31",
		F:    ossm.TestExtensionInstall,
	},
	testing.InternalTest{
		Name: "32",
		F:    ossm.TestTLSVersionSMCP,
	},
	testing.InternalTest{
		Name: "33",
		F:    ossm.TestSSL,
	},
	testing.InternalTest{
		Name: "34",
		F:    ossm.TestRateLimiting,
	},
	testing.InternalTest{
		Name: "35",
		F:    ossm.TestProxyEnv,
	},
}
