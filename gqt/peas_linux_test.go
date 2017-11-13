package gqt_test

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/gqt/runner"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Partially shared containers (peas)", func() {
	var (
		gdn           *runner.RunningGarden
		tmpDir        string
		peaRootfs     string
		ctr           garden.Container
		containerSpec garden.ContainerSpec
	)

	BeforeEach(func() {
		gdn = runner.Start(config)
		var err error
		tmpDir, err = ioutil.TempDir("", "peas-gqts")
		Expect(err).NotTo(HaveOccurred())

		Expect(exec.Command("cp", "-a", defaultTestRootFS, tmpDir).Run()).To(Succeed())
		Expect(os.Chmod(tmpDir, 0777)).To(Succeed())
		peaRootfs = filepath.Join(tmpDir, "rootfs")
		Expect(exec.Command("chown", "-R", "4294967294:4294967294", peaRootfs).Run()).To(Succeed())
		Expect(ioutil.WriteFile(filepath.Join(peaRootfs, "ima-pea"), []byte("pea!"), 0644)).To(Succeed())
		containerSpec = garden.ContainerSpec{}
	})

	JustBeforeEach(func() {
		var err error
		ctr, err = gdn.Create(containerSpec)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(gdn.DestroyAndStop()).To(Succeed())
		Expect(os.RemoveAll(tmpDir)).To(Succeed())
	})

	It("should not leak pipes", func() {
		initialPipes := numPipes(gdn.Pid)

		process, err := ctr.Run(garden.ProcessSpec{
			Path:  "echo",
			Args:  []string{"hello"},
			Image: garden.ImageRef{URI: "raw://" + peaRootfs},
		}, garden.ProcessIO{})
		Expect(err).NotTo(HaveOccurred())
		Expect(process.Wait()).To(Equal(0))

		Eventually(func() int { return numPipes(gdn.Pid) }).Should(Equal(initialPipes))
	})

	Context("bind mounts", func() {
		var testSrcFile *os.File
		destinationFile := "/tmp/file"
		output := gbytes.NewBuffer()

		BeforeEach(func() {
			var err error
			testSrcFile, err = ioutil.TempFile(tmpDir, "host-file")
			Expect(err).NotTo(HaveOccurred())
			_, err = testSrcFile.WriteString("test-mount")
			Expect(err).NotTo(HaveOccurred())
			Expect(exec.Command("chown", "4294967294", testSrcFile.Name()).Run()).To(Succeed())
		})

		Context("when we create a pea with bind mounts", func() {
			It("should have access to the mounts", func() {
				process, err := ctr.Run(garden.ProcessSpec{
					Path:  "cat",
					Args:  []string{destinationFile},
					Image: garden.ImageRef{URI: "raw://" + peaRootfs},
					BindMounts: []garden.BindMount{
						garden.BindMount{
							SrcPath: testSrcFile.Name(),
							DstPath: destinationFile,
						},
					},
				}, garden.ProcessIO{
					Stdout: io.MultiWriter(GinkgoWriter, output),
					Stderr: GinkgoWriter,
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(process.Wait()).To(Equal(0))
				Expect(output).To(gbytes.Say("test-mount"))
			})
		})

		Context("when there are already bind mounts in the container", func() {
			BeforeEach(func() {
				containerSpec = garden.ContainerSpec{
					BindMounts: []garden.BindMount{
						garden.BindMount{
							SrcPath: testSrcFile.Name(),
							DstPath: destinationFile,
						},
					},
				}
			})

			It("the pea should not have access to the mounts", func() {
				process, err := ctr.Run(garden.ProcessSpec{
					Path:  "cat",
					Args:  []string{destinationFile},
					Image: garden.ImageRef{URI: "raw://" + peaRootfs},
				}, garden.ProcessIO{
					Stdout: GinkgoWriter,
					Stderr: io.MultiWriter(GinkgoWriter, output),
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(process.Wait()).To(Equal(1))
				Expect(output).To(gbytes.Say("No such file or directory"))
			})
		})
	})
})
