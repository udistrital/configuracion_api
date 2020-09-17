package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaMetodoHttp_20191114_143538 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaMetodoHttp_20191114_143538{}
	m.Created = "20191114_143538"

	migration.Register("CrearTablaMetodoHttp_20191114_143538", m)
}

// Run the migrations
func (m *CrearTablaMetodoHttp_20191114_143538) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS configuracion.metodo_http ( id serial NOT NULL, nombre character varying NOT NULL, descripcion character varying NOT NULL, CONSTRAINT pk_metodo_http PRIMARY KEY (id) );")
	m.SQL("ALTER TABLE configuracion.metodo_http OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE configuracion.metodo_http IS 'Tipos de metodo http usados';")
	m.SQL("COMMENT ON COLUMN configuracion.metodo_http.id IS 'Identificador único de la tabla';")
	m.SQL("COMMENT ON COLUMN configuracion.metodo_http.nombre IS 'Identificador único de la configuracion generada';")
	m.SQL("COMMENT ON COLUMN configuracion.metodo_http.descripcion IS 'Identificador único de la configuracion generada';")
}

// Reverse the migrations
func (m *CrearTablaMetodoHttp_20191114_143538) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS configuracion.metodo_http")
}
