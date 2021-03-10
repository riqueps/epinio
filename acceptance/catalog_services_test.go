package acceptance_test

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/suse/carrier/helpers"
)

var _ = Describe("Catalog Services", func() {
	var org = "apps-org"
	var serviceName string
	BeforeEach(func() {
		serviceName = "service-" + strconv.Itoa(int(time.Now().Nanosecond()))

		out, err := Carrier("create-org "+org, "")
		Expect(err).ToNot(HaveOccurred(), out)
		out, err = Carrier("target "+org, "")
		Expect(err).ToNot(HaveOccurred(), out)

		out, err = Carrier("enable services-incluster", "")
		Expect(err).ToNot(HaveOccurred(), out)

		// Wait until plans appear
		Eventually(func() error {
			_, err = helpers.Kubectl("get clusterserviceclass mariadb")
			return err
		}, "5m").ShouldNot(HaveOccurred())
	})

	Describe("create-service", func() {
		AfterEach(func() {
			out, err := Carrier(fmt.Sprintf("delete-service %s", serviceName), "")
			Expect(err).ToNot(HaveOccurred(), out)
		})

		It("creates a catalog based service", func() {
			out, err := Carrier(fmt.Sprintf("create-service %s mariadb 10-3-22", serviceName), "")
			Expect(err).ToNot(HaveOccurred(), out)
			out, err = Carrier("services", "")
			Expect(err).ToNot(HaveOccurred(), out)
			Expect(out).To(MatchRegexp(serviceName))
		})
	})

	Describe("delete service", func() {
		BeforeEach(func() {
			out, err := Carrier(fmt.Sprintf("create-service %s mariadb 10-3-22", serviceName), "")
			Expect(err).ToNot(HaveOccurred(), out)
		})

		It("deletes a catalog based service", func() {
			out, err := Carrier("delete-service "+serviceName, "")
			Expect(err).ToNot(HaveOccurred(), out)

			Eventually(func() string {
				out, err = Carrier("services", "")
				Expect(err).ToNot(HaveOccurred(), out)
				return out
			}, "10m").ShouldNot(MatchRegexp(serviceName))
		})

		PIt("doesn't delete a bound service", func() {
		})
	})

	Describe("bind-service", func() {
		var appName string
		BeforeEach(func() {
			appName = "apps-" + strconv.Itoa(int(time.Now().Nanosecond()))

			out, err := Carrier(fmt.Sprintf("create-service %s mariadb 10-3-22", serviceName), "")
			Expect(err).ToNot(HaveOccurred(), out)

			currentDir, err := os.Getwd()
			Expect(err).ToNot(HaveOccurred())
			appDir := path.Join(currentDir, "../sample-app")
			out, err = Carrier(fmt.Sprintf("push %s --verbosity 1", appName), appDir)
			Expect(err).ToNot(HaveOccurred(), out)
		})

		AfterEach(func() {
			out, err := Carrier(fmt.Sprintf("unbind-service %s %s", serviceName, appName), "")
			if err != nil {
				fmt.Printf("unbinding service failed: %s\n%s", err.Error(), out)
			}

			out, err = Carrier("delete "+appName, "")
			if err != nil {
				fmt.Printf("deleting app failed : %s\n%s", err.Error(), out)
			}

			out, err = Carrier("delete-service "+serviceName, "")
			if err != nil {
				fmt.Printf("deleting service failed : %s\n%s", err.Error(), out)
			}
		})

		It("binds a service to the application deployment", func() {
			out, err := Carrier(fmt.Sprintf("bind-service %s %s", serviceName, appName), "")
			Expect(err).ToNot(HaveOccurred(), out)
			out, err = helpers.Kubectl(fmt.Sprintf("get deployment -n carrier-workloads %s.%s -o=jsonpath='{.spec.template.spec.volumes}'", org, appName))
			Expect(err).ToNot(HaveOccurred(), out)
			Expect(out).To(MatchRegexp(serviceName))

			out, err = helpers.Kubectl(fmt.Sprintf("get deployment -n carrier-workloads %s.%s -o=jsonpath='{.spec.template.spec.containers[0].volumeMounts}'", org, appName))
			Expect(err).ToNot(HaveOccurred(), out)
			Expect(out).To(MatchRegexp("/services/" + serviceName))
		})
	})

	Describe("unbind-service", func() {
		var appName string
		BeforeEach(func() {
			appName = "apps-" + strconv.Itoa(int(time.Now().Nanosecond()))

			out, err := Carrier(fmt.Sprintf("create-service %s mariadb 10-3-22", serviceName), "")
			Expect(err).ToNot(HaveOccurred(), out)

			currentDir, err := os.Getwd()
			Expect(err).ToNot(HaveOccurred())
			appDir := path.Join(currentDir, "../sample-app")
			out, err = Carrier(fmt.Sprintf("push %s --verbosity 1", appName), appDir)
			Expect(err).ToNot(HaveOccurred(), out)

			out, err = Carrier(fmt.Sprintf("bind-service %s %s", serviceName, appName), "")
			Expect(err).ToNot(HaveOccurred(), out)
		})

		AfterEach(func() {
			out, err := Carrier("delete "+appName, "")
			if err != nil {
				fmt.Printf("deleting apps failed : %s\n%s", err.Error(), out)
			}

			out, err = Carrier("delete-service "+serviceName, "")
			if err != nil {
				fmt.Printf("deleting service failed : %s\n%s", err.Error(), out)
			}
		})

		It("unbinds a service from the application deployment", func() {
			out, err := Carrier(fmt.Sprintf("unbind-service %s %s", serviceName, appName), "")
			Expect(err).ToNot(HaveOccurred(), out)
			out, err = helpers.Kubectl(fmt.Sprintf("get deployment -n carrier-workloads %s.%s -o=jsonpath='{.spec.template.spec.volumes}'", org, appName))
			Expect(err).ToNot(HaveOccurred(), out)
			Expect(out).ToNot(MatchRegexp(serviceName))

			out, err = helpers.Kubectl(fmt.Sprintf("get deployment -n carrier-workloads %s.%s -o=jsonpath='{.spec.template.spec.containers[0].volumeMounts}'", org, appName))
			Expect(err).ToNot(HaveOccurred(), out)
			Expect(out).ToNot(MatchRegexp("/services/" + serviceName))
		})
	})
})