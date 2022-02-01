package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Proceso_20211120_212333 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Proceso_20211120_212333{}
	m.Created = "20211120_212333"

	migration.Register("Proceso_20211120_212333", m)
}

// Run the migrations
func (m *Proceso_20211120_212333) Up() {
	file, err := ioutil.ReadFile("../scripts/20211120_212333_procesos.up.sql")

	if err != nil {
		// handle error
		logs.Error(err)
		panic(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}

// Reverse the migrations
func (m *Proceso_20211120_212333) Down() {
	file, err := ioutil.ReadFile("../scripts/20211120_212333_procesos.down.sql")

	if err != nil {
		// handle error
		logs.Error(err)
		panic(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}
