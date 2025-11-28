package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearSchemaConfiguracion_20191114_141832 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearSchemaConfiguracion_20191114_141832{}
	m.Created = "20191114_141832"

	migration.Register("CrearSchemaConfiguracion_20191114_141832", m)
}

// Run the migrations
func (m *CrearSchemaConfiguracion_20191114_141832) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE SCHEMA IF NOT EXISTS configuracion;")
	m.SQL("ALTER SCHEMA configuracion OWNER TO desarrollooas;")
	m.SQL("SET search_path TO pg_catalog,public,configuracion;")
}

// Reverse the migrations
func (m *CrearSchemaConfiguracion_20191114_141832) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP SCHEMA IF EXISTS configuracion")
}
