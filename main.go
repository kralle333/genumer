package main

import "genumer/cmd"

//go:generate go run main.go --name "enumStyle" --values "camelCase,pascalCase,upperCase,kebabCase,screamingKebabCase,snakeCase,screamingSnakeCase" --dest "cmd/" --private

func main() {
	cmd.Execute()
}
