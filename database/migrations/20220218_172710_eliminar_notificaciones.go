package main

import (
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type EliminarNotificaciones_20220218_172710 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &EliminarNotificaciones_20220218_172710{}
	m.Created = "20220218_172710"

	migration.Register("EliminarNotificaciones_20220218_172710", m)
}

// Run the migrations
func (m *EliminarNotificaciones_20220218_172710) Up() {
	file, err := ioutil.ReadFile("../scripts/20220218_172710_eliminar_notificaciones.up.sql")

	if err != nil {
		// handle error
		logs.Error(err)
		panic(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		// fmt.Println(request)
		m.SQL(request)
	}
}

// Reverse the migrations
func (m *EliminarNotificaciones_20220218_172710) Down() {
	file, err := ioutil.ReadFile("../scripts/20220218_172710_eliminar_notificaciones.down.sql")

	if err != nil {
		// handle error
		logs.Error(err)
		panic(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		// fmt.Println(request)
		m.SQL(request)
	}
}
