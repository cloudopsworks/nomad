// This package exists to wrap our e2e provisioning and test framework so that it
// can be run via 'go test ./e2e'. See './framework/framework.go'
package e2e

import (
	"os"
	"testing"

	"github.com/hashicorp/nomad/e2e/framework"

	_ "github.com/hashicorp/nomad/e2e/affinities"
	_ "github.com/hashicorp/nomad/e2e/clientstate"

	_ "github.com/hashicorp/nomad/e2e/connect"
	_ "github.com/hashicorp/nomad/e2e/consul"
	_ "github.com/hashicorp/nomad/e2e/consultemplate"
	_ "github.com/hashicorp/nomad/e2e/csi"
	_ "github.com/hashicorp/nomad/e2e/deployment"
	_ "github.com/hashicorp/nomad/e2e/events"
	_ "github.com/hashicorp/nomad/e2e/example"
	_ "github.com/hashicorp/nomad/e2e/isolation"
	_ "github.com/hashicorp/nomad/e2e/license"
	_ "github.com/hashicorp/nomad/e2e/lifecycle"
	_ "github.com/hashicorp/nomad/e2e/metrics"
	_ "github.com/hashicorp/nomad/e2e/namespaces"
	_ "github.com/hashicorp/nomad/e2e/networking"
	_ "github.com/hashicorp/nomad/e2e/nodedrain"
	_ "github.com/hashicorp/nomad/e2e/nomad09upgrade"
	_ "github.com/hashicorp/nomad/e2e/nomadexec"
	_ "github.com/hashicorp/nomad/e2e/oversubscription"
	_ "github.com/hashicorp/nomad/e2e/parameterized"
	_ "github.com/hashicorp/nomad/e2e/periodic"
	_ "github.com/hashicorp/nomad/e2e/podman"
	_ "github.com/hashicorp/nomad/e2e/quotas"
	_ "github.com/hashicorp/nomad/e2e/rescheduling"
	_ "github.com/hashicorp/nomad/e2e/scaling"
	_ "github.com/hashicorp/nomad/e2e/scalingpolicies"
	_ "github.com/hashicorp/nomad/e2e/spread"
	_ "github.com/hashicorp/nomad/e2e/systemsched"
	_ "github.com/hashicorp/nomad/e2e/taskevents"
	_ "github.com/hashicorp/nomad/e2e/vaultsecrets"
	_ "github.com/hashicorp/nomad/e2e/volumes"
)

func TestE2E(t *testing.T) {
	if os.Getenv("NOMAD_E2E") == "" {
		t.Skip("Skipping e2e tests, NOMAD_E2E not set")
	} else {
		framework.Run(t)
	}
}
