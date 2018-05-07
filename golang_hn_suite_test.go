package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGolangHn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GolangHn Suite")
}
