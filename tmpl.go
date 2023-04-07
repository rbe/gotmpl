package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "sort"
    "strings"
    "text/template"
)

type Todo struct {
    Name        string
    Description string
    Env         map[string]string
}

func main() {
    var templateFile string
    flag.StringVar(&templateFile, "template-file", "", "a template file path")
    flag.Parse()

    data := readTemplate(templateFile)
    t := template.Must(template.New("todos").Parse(string(data)))
    todo := Todo{"Test templates",
        "Let's test a template to see the magic.",
        getenv()}
    err := t.Execute(os.Stdout, todo)
    if err != nil {
        panic(err)
    }
}

func readTemplate(templateFile string) []byte {
    //filenames := []string{"tmpl/todo.gohtml"}
    filenames := []string{templateFile}
    data, err := ioutil.ReadFile(filenames[0])
    if err != nil {
        panic(err)
    }
    return data
}

func getenv() map[string]string {
    getenvironment := func(data []string, getkeyval func(item string) (key, val string)) map[string]string {
        items := make(map[string]string)
        for _, item := range data {
            key, val := getkeyval(item)
            items[key] = val
        }
        return items
    }
    environment := getenvironment(os.Environ(), func(item string) (key, val string) {
        splits := strings.Split(item, "=")
        key = splits[0]
        val = splits[1]
        return
    })
    return environment
}

func printEnvironment() {
    bla := os.Environ()
    sort.Strings(bla)
    for _, e := range bla {
        pair := strings.SplitN(e, "=", 2)
        fmt.Println(pair[0], "=", pair[1])
    }
}

func printCurrentWorkingDir() error {
    path, err := os.Getwd()
    if err != nil {
        log.Println(err)
    } else {
        fmt.Println(path)
    }
    return err
}
