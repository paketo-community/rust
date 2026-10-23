package integration_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/paketo-buildpacks/occam"
	. "github.com/paketo-buildpacks/occam/matchers"
	"github.com/sclevine/spec"
)

func testMultipleBinaries(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect     = NewWithT(t).Expect
		Eventually = NewWithT(t).Eventually

		pack   occam.Pack
		docker occam.Docker

		image           occam.Image
		container       occam.Container
		workerContainer occam.Container

		name   string
		source string
	)

	it.Before(func() {
		pack = occam.NewPack()
		docker = occam.NewDocker()

		var err error
		name, err = occam.RandomName()
		Expect(err).NotTo(HaveOccurred())

		source, err = occam.Source(filepath.Join("testdata", "multi_bin_app"))
		Expect(err).NotTo(HaveOccurred())
	})

	it.After(func() {
		if container.ID != "" {
			Expect(docker.Container.Remove.Execute(container.ID)).To(Succeed())
		}
		if workerContainer.ID != "" {
			Expect(docker.Container.Remove.Execute(workerContainer.ID)).To(Succeed())
		}
		Expect(docker.Image.Remove.Execute(image.ID)).To(Succeed())
		Expect(docker.Volume.Remove.Execute(occam.CacheVolumeNames(name))).To(Succeed())
		Expect(os.RemoveAll(source)).To(Succeed())
	})

	it("builds every binary in a Cargo workspace and runs each process type", func() {
		var err error
		var logs fmt.Stringer
		image, logs, err = pack.WithNoColor().Build.
			WithBuildpacks(rustBuildpack).
			WithPullPolicy("never").
			Execute(name, source)
		Expect(err).NotTo(HaveOccurred(), logs.String())

		Expect(logs.String()).To(ContainSubstring("--path=/workspace/web-app"))
		Expect(logs.String()).To(ContainSubstring("--path=/workspace/worker-app"))

		container, err = docker.Container.Run.
			WithEnv(map[string]string{"PORT": "8080"}).
			WithPublish("8080").
			Execute(image.ID)
		Expect(err).NotTo(HaveOccurred())

		Eventually(container).Should(BeAvailable())
		Eventually(container).Should(Serve(ContainSubstring("Hello, World!")).OnPort(8080))

		workerContainer, err = docker.Container.Run.
			WithEntrypoint("worker").
			Execute(image.ID)
		Expect(err).NotTo(HaveOccurred())

		Eventually(func() string {
			workerLogs, err := docker.Container.Logs.Execute(workerContainer.ID)
			Expect(err).NotTo(HaveOccurred())
			return workerLogs.String()
		}).Should(ContainSubstring("worker started"))
	})
}
