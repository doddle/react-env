package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	flag "github.com/spf13/pflag"
)

type EnvVar struct {
	Key   string
	Value string
}

type EnvVarList []EnvVar

func generateEnvVarList(keys []string) EnvVarList {
	result := EnvVarList{}
	for _, key := range keys {
		value := os.Getenv(key)
		varObj := EnvVar{
			Key:   key,
			Value: value,
		}
		result = append(result, varObj)
	}
	return result

}

func main() {

	destPath := flag.String("dest", ".", "Destination path to write to (relative or absolute)")
	destFile := flag.String("file", ".env.js", "name of the file to write")
	prefix := flag.String("prefix", "REACT_APP_", "Prefixes of env vars to look for")
	flag.Parse()

	if *destPath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	absPath, err := filepath.Abs(*destPath)
	if err != nil {
		fmt.Printf("could not resolve a path for %s\n", *destPath)
		os.Exit(1)
	}

	// check dir exists (filepath.Abs already does this but better to be safe)
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		fmt.Printf("path %s does not exist\n", absPath)
		os.Exit(1)
	}

	fileName := fmt.Sprintf("%s/%s", absPath, *destFile)

	envs := envsWithPrefix(*prefix)
	fmt.Printf("found envs: %s\n", envs)
	fmt.Printf("writing to: %s\n", fileName)

	output := generateEnvVarList(envs)

	err = generateEnvJS(output, fileName)
	//fmt.Println(outputString)
	if err != nil {
		fmt.Println("error writing")
		fmt.Println(err)
		os.Exit(3)
	}

}

func generateEnvJS(input EnvVarList, file string) error {

	const tpl = `
window._env = { {{range $index, $item := .Items}} {{if ne $index 0}},{{end}}"{{ $item.Key }}":"{{ $item.Value }}" {{end}} }
`

	t, err := template.New("webpage").Parse(tpl)
	if err != nil {
		return err
	}

	data := struct {
		Items []EnvVar
	}{
		Items: input,
	}

	f, err := os.Create(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}

	// ensure the handler gets closed when its time
	defer f.Close()
	err = t.Execute(f, data)
	if err != nil {
		return err
	}

	return err

}
func envsWithPrefix(prefix string) []string {
	result := []string{}
	// returns the names of env vars matching a prefix
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if strings.HasPrefix(pair[0], prefix) {
			result = append(result, pair[0])
		}
	}
	return result
}
