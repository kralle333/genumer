package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type Config struct {
	name        string
	values      []string
	destination string
	goPackage   string
	style       enumStyle
	createDir   bool
	private     bool
}

func splitValuesString(s string) ([]string, error) {

	valuesSlice := strings.Split(s, ",")
	if len(valuesSlice) == 0 || s == "" {
		return nil, fmt.Errorf("no enum values found, usage: \"an,array,of,strings\"")
	}
	return valuesSlice, nil
}

func getPackage(destination string) string {
	if strings.HasSuffix(destination, "/") {
		destination = destination[:len(destination)-1]
	}
	splitted := strings.Split(destination, "/")
	if len(splitted) > 1 {
		return splitted[len(splitted)-1]
	}
	return destination
}

func getGoPackage(inputGoPackage string, destination string) string {
	if inputGoPackage != "" {
		if strings.HasSuffix(inputGoPackage, "/") {
			inputGoPackage = inputGoPackage[:len(inputGoPackage)-1]
		}
		return inputGoPackage
	}
	return getPackage(destination)
}

// Root command for the generator
var rootCmd = &cobra.Command{
	Use:   "genumer",
	Short: "Generate go enums",
	Long: `A tool for generating go enums in various styles. 

Either use it from the command line:
genumer --name="MyEnum"" --values="firstValue,secondValue" --dest="cmd/"
which will generate a file called myenum.go in the cmd/ directory

or use it as a go generate directive:
go:generate genumer --name "enumStyle" --values "camelCase,pascalCasing,screamingSnakeCase" --dest "cmd/test/" --createDir
which will generate a file called enumstyle.go in the cmd/test directory, creating the directories if they do not exist
`,
	Example: "genumer --name=myEnum --values=\"firstValue,secondValue\" --dest=generated/",
	Run: func(cmd *cobra.Command, args []string) {

		values, err := splitValuesString(values)
		if err != nil {
			cmd.PrintErrln("parsing values failed:", err)
			return
		}

		style, err := getStyleOrDefault(style, enumStyleCamelCase)
		if err != nil {
			cmd.PrintErrln("parsing style failed:", err)
			return
		}
		generatorConfig := Config{
			name:        name,
			values:      values,
			destination: destination,
			goPackage:   getGoPackage(goPackage, destination),
			style:       style,
			createDir:   createDir,
			private:     private,
		}

		if generatorConfig.goPackage == "" {
			cmd.PrintErrln("could not deduce package name from destination")
			return
		}

		err = CreateGoFile(generatorConfig)
		if err != nil {
			cmd.PrintErrln("generating enum failed:", err)
		}
	},
}

// Flags for the generator
var (
	name        string
	values      string
	destination string
	goPackage   string
	style       string
	createDir   bool
	private     bool
)

func markRequiredOrPanic(name string) {
	err := rootCmd.MarkFlagRequired(name)
	if err != nil {
		panic(err)
	}

}
func stylesHelp() string {
	var options []string
	for _, s := range allEnumStyles {
		options = append(options, s.String())
	}
	return strings.Join(options, ", ")
}

func init() {
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the enum in camelCase, e.g. \"myEnum\"")
	rootCmd.Flags().StringVarP(&values, "values", "v", "", "Target enum values in camelCase, e.g. \"firstValue,thenSecond,theVeryLastValue\"")
	rootCmd.Flags().StringVarP(&destination, "dest", "d", "", "Destination of generated go enum file")
	rootCmd.Flags().StringVarP(&goPackage, "package", "p", "", "Package of generated go enum file. If not set, the destination will be used to deduct package name")
	rootCmd.Flags().StringVarP(&style, "style", "s", "camelCase", fmt.Sprintf("Style of the generated go enum names: %s", stylesHelp()))
	rootCmd.Flags().BoolVar(&createDir, "createDir", false, "Create the destination directory/directories if they do not exist")
	rootCmd.Flags().BoolVar(&private, "private", false, "Make the generated enum private")

	markRequiredOrPanic("name")
	markRequiredOrPanic("values")
	markRequiredOrPanic("dest")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
