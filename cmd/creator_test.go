package cmd

import (
	"os"
	"testing"
)

const (
	enumName      = "TestEnum"
	testGoPackage = "enumgenerator"
	goEnumFile    = "testenum.go"
)

var enumValues = []string{"firstValue", "secondVal", "lastOne"}

func testConfig() Config {
	return Config{
		name:        enumName,
		values:      enumValues,
		destination: ".",
		goPackage:   testGoPackage,
		style:       enumStyleCamelCase,
		createDir:   false,
	}
}

func testCreateGoEnum(t *testing.T, config Config) {
	defer os.Remove(goEnumFile)
	err := CreateGoFile(config)
	failIfErrNotNil(t, err)
	_, err = os.Stat(goEnumFile)
	failIfErrNotNil(t, err)
}

func failIfErrNotNil(t *testing.T, err error) {
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
}

func TestGenerateGoEnums(t *testing.T) {
	testCreateGoEnum(t, testConfig())
}

func TestGenerateGoAndGqlEnums(t *testing.T) {
	testCreateGoEnum(t, testConfig())
}
