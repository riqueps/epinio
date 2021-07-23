package acceptance_test

import (
	"fmt"
	"io/ioutil"
	"runtime"

	"github.com/epinio/epinio/helpers/epinio"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Epinio Installation", func() {
	var flags string
	epinioBinary := fmt.Sprintf("../../dist/epinio-%s-%s", runtime.GOOS, runtime.GOARCH)

	configFile := "./assets/tests/config.yaml"
	epinioHelper := epinio.NewEpinioHelper(epinioBinary)
	epinioUser := "epinio"
	epinioPassword := "password"

	BeforeEach(func() {
		flags = fmt.Sprintf("--config-file %s", configFile)
		flags = fmt.Sprintf("%s --skip-default-org", flags)
		flags = fmt.Sprintf("%s --user %s --password %s", flags, epinioUser, epinioPassword)

		epinioHelper.Flags = flags
	})

	AfterEach(func() {
		epinioHelper.Uninstall()
	})

	When("a epinio config file already exits", func() {
		It("should install epinio with new values and update the file", func() {
			By("Installing epinio")
			out, err := epinioHelper.Install()
			Expect(err).NotTo(HaveOccurred())

			By("Checking for updated values in epinio config file")
			data, err := ioutil.ReadFile(configFile)
			Expect(err).NotTo(HaveOccurred())
			dataString := string(data)

			// The values for checking are taken from ./assets/tests/config.yaml
			Expect(dataString).NotTo(ContainSubstring("pass: 05ec82a894940780"))
			Expect(dataString).NotTo(ContainSubstring("user: 996ee615fde2ceed"))
			Expect(dataString).To(ContainSubstring("pass: password"))
			Expect(dataString).To(ContainSubstring("user: epinio"))

			By("Checking the values in the stdout")
			Expect(out).To(ContainSubstring("API Password: password"))
			Expect(out).To(ContainSubstring("API User: epinio"))
		})
	})
})
