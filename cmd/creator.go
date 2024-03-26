package cmd

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"go/format"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	GoViewPath = "enumtemplate.txt"
)

//go:embed enumtemplate.txt
var views embed.FS

func renameCreateFile(fileName string, createDir bool) (*os.File, error) {
	_, err := os.Stat(fileName)

	if err == nil {
		err = os.Rename(fileName, fmt.Sprintf("%s.bak", fileName))
		if err != nil {
			return nil, err
		}
	}

	if createDir {
		dirName := filepath.Dir(fileName)
		err = os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	return os.Create(fileName)
}

func handleBakFile(fileName string) {
	backupFilePath := fmt.Sprintf("%s.bak", fileName)
	_, err := os.Stat(backupFilePath)
	if err != nil {
		return // no backup file
	}
	_, genFileErr := os.Stat(fileName)
	if genFileErr == nil { // generated file exists
		err = os.Remove(backupFilePath)
	} else { // generated file does not exist, rename backup file back
		err = os.Rename(backupFilePath, fileName)
	}
	if err != nil {
		fmt.Println(err)
	}
}

func CreateGoFile(config Config) error {
	toPascalCasingFunc := func(s string) string { return fmt.Sprintf("%s%s", strings.ToUpper(string(s[0])), s[1:]) }
	toCamelCaseFunc := func(s string) string { return fmt.Sprintf("%s%s", strings.ToLower(string(s[0])), s[1:]) }
	type data struct {
		Type             string
		AllValuesVarName string
		TypeLower        string
		FirstLetter      string
		Package          string
		StylizedValues   []string
		Values           []string
	}
	config.name = toPascalCasingFunc(config.name)

	d := data{
		Type:             config.name,
		AllValuesVarName: fmt.Sprintf("All%ss", toPascalCasingFunc(config.name)),
		TypeLower:        toCamelCaseFunc(config.name),
		Package:          config.goPackage,
		Values:           config.values,
		StylizedValues:   applyEnumStyle(config.values, config.style),
		FirstLetter:      strings.ToLower(string(config.name[0])),
	}
	if config.private {
		d.Type = fmt.Sprintf(d.TypeLower)
		d.AllValuesVarName = toCamelCaseFunc(d.AllValuesVarName)
	}

	funcMap := template.FuncMap{
		"ToPascalCasing": toPascalCasingFunc,
	}

	goTemplate, err := template.New(GoViewPath).Funcs(funcMap).ParseFS(views, GoViewPath)
	var buf bytes.Buffer
	err = goTemplate.Execute(&buf, d)
	if err != nil {
		return err
	}
	p, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	goDestDirPath := filepath.Dir(config.destination)
	fileName := path.Join(goDestDirPath, fmt.Sprintf("%s.go", strings.ToLower(config.name)))
	goFile, err := renameCreateFile(fileName, config.createDir)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("destination directory does not exist: %s - use --createDir to create destination directories", goDestDirPath)
		}
		return err
	}
	defer handleBakFile(fileName)
	_, err = goFile.Write(p)
	if err != nil {
		return err
	}

	if err = goFile.Close(); err != nil && !errors.Is(err, os.ErrClosed) {
		return err
	}

	return nil
}
