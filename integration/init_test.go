package integration_test

import (
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

var rustBuildpack string

func TestIntegration(t *testing.T) {
	Expect := NewWithT(t).Expect

	output, err := exec.Command("bash", "-c", "../scripts/package.sh --version 1.2.3").CombinedOutput()
	Expect(err).NotTo(HaveOccurred(), string(output))

	rustBuildpack, err = filepath.Abs("../build/buildpackage.cnb")
	Expect(err).NotTo(HaveOccurred())

	format.MaxLength = 0
	SetDefaultEventuallyTimeout(10 * time.Second)
	SetDefaultEventuallyPollingInterval(100 * time.Millisecond)

	// Run the suites sequentially: each build pulls and caches a full Rust
	// toolchain, so running them in parallel can exhaust the runner's disk.
	suite := spec.New("Integration", spec.Sequential(), spec.Report(report.Terminal{}))
	suite("Simple App", testSimpleApp)
	suite("Multiple Binaries", testMultipleBinaries)
	suite("Procfile", testProcfile)
	suite("Toolchain File", testToolchainFile)
	suite("Rust Distribution", testRustDistribution)
	suite.Run(t)
}
