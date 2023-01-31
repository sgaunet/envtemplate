package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

func ReadTemplate(instream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(instream)
	return buf.String()
}

func ReadEnvVars(rawEnv []string) (environ map[string]string) {
	environ = make(map[string]string)
	for _, item := range rawEnv {
		parts := strings.SplitN(item, "=", 2)
		environ[parts[0]] = parts[1]
	}
	return
}

func WriteTemplateToStream(tplSource string, environ map[string]string, outStream io.Writer) {
	tpl := template.New("_root_")
	tpl.Funcs(template.FuncMap{
		"split":  TplSplitStr,
		"exists": TplCheckExists,
	})
	tpl.Option("missingkey=error")
	_, err := tpl.Parse(tplSource)
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.Execute(outStream, environ)
	if err != nil {
		log.Fatal(err)
	}
}

func TplSplitStr(args ...interface{}) ([]string, error) {
	rawValue := args[0].(string)
	sep := args[1].(string)
	limit := -1
	if len(args) > 2 {
		parsedLimit, ok := args[2].(int)
		if !ok {
			err := errors.New("limit arg (3rd) to `split` is not integer")
			return nil, err
		}
		limit = parsedLimit
	}
	return strings.SplitN(rawValue, sep, limit), nil
}

func TplCheckExists(args ...interface{}) (bool, error) {
	datamap, ok := args[0].(map[string]string)
	if !ok {
		return false, errors.New("data-map arg (1st) to `exists` should be a map[string]string, did you mean '.' or '$'?")
	}
	key, ok := args[1].(string)
	if !ok {
		return false, errors.New("lookup-key arg (2nd) to `exists` should be a string")
	}
	_, exists := datamap[key]
	return exists, nil
}

var version string = "development"

func printVersion() {
	fmt.Println(version)
}

func main() {
	var err error
	var f io.Reader
	var inputFile string
	var vOption, helpOption bool
	flag.StringVar(&inputFile, "i", "", "File to encrypt/decrypt")
	flag.BoolVar(&vOption, "v", false, "Get version")
	flag.BoolVar(&helpOption, "h", false, "Print help")
	flag.Parse()

	if vOption {
		printVersion()
		os.Exit(0)
	}

	if helpOption {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if inputFile != "" {
		f, err = os.Open(inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error when opening file %s: %s\n", inputFile, err.Error())
			os.Exit(1)
		}
	} else {
		f = os.Stdin
	}
	WriteTemplateToStream(ReadTemplate(f), ReadEnvVars(os.Environ()), os.Stdout)
}
