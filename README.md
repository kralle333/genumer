# genumer - A go enum generator

Enums are not a first class citizen in Go, and are often implemented as a set of constants.
This tool aims to make it easier to create enums in Go and has various options for generating the enum in different
styles and languages, which is especially useful when working with a codebase with manually written enums that already
have a certain style.

## Enum Styles

Available Styles:

- `camelCase`: thisIsCamelCase
- `pascalCase`: ThisIsPascalCase
- `snakeCase`: this_is_snake_case
- `kebabCase`: this-is-kebab-case
- `upperCase`: THISISUPPERCASE
- `screamingKebabCase` THIS-IS-SCREAMING-KEBAB-CASE
- `screamingSnakeCase` THIS_IS_SCREAMING_SNAKE_CASE

## Usage

Either use it from the command line:

```
genumer --name="MyEnum"" --values="firstValue,secondValue" --dest="cmd/"
```
which will generate a file called myenum.go in the cmd/ directory

or use it as a go generate directive:
```
//go:generate genumer --name "enumStyle" --values "camelCase,pascalCasing,screamingSnakeCase" --dest "cmd/test/"
--createDir
```
which will generate a file called enumstyle.go in the cmd/test directory, creating the directories if they do not exist

### Help text

```
genumer [flags]

Examples:
genumer --name=myEnum --values="firstValue,secondValue" --dest=generated/

Flags:
--createDir        Create the destination directory/directories if they do not exist
-d, --dest string      Destination of generated go enum file
-h, --help             help for genumer
-n, --name string      Name of the enum in camelCase, e.g. "myEnum"
-p, --package string   Package of generated go enum file. If not set, the destination will be used to deduct package name
--private          Make the generated enum private
-s, --style string     Style of the generated go enum names: camelCase, pascalCase, upperCase, kebabCase, screamingKebabCase, snakeCase, screamingSnakeCase (default "camelCase")
-v, --values string    Target enum values in camelCase, e.g. "firstValue,thenSecond,theVeryLastValue"

```

## Installation

### go install
```
go install github.com/kralle333/genumer
```

### homebrew

```
brew tap kralle333/genumer
brew install genumer
```

### using releases
install the binary of your choice from [the git releases](https://github.com/kralle333/genumer/releases)
