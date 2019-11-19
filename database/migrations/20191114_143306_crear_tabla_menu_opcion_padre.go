package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaMenuOpcionPadre_20191114_143306 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaMenuOpcionPadre_20191114_143306{}
	m.Created = "20191114_143306"

	migration.Register("CrearTablaMenuOpcionPadre_20191114_143306", m)
}

// Run the migrations
func (m *CrearTablaMenuOpcionPadre_20191114_143306) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS configuracion.menu_opcion_padre ( id serial NOT NULL, padre integer NOT NULL, hijo integer, CONSTRAINT PK_MENU_OPCION_PADRE PRIMARY KEY (id), CONSTRAINT UQ_PADRE_HIJO UNIQUE (padre, hijo), CONSTRAINT FK_MENU_OPCION_HIJO_MENU_OPCION FOREIGN KEY (hijo) REFERENCES configuracion.menu_opcion (id), CONSTRAINT FK_MENU_OPCION_PADRE_MENU_OPCION FOREIGN KEY (padre) REFERENCES configuracion.menu_opcion (id) );")
	m.SQL("ALTER TABLE configuracion.menu_opcion_padre OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE configuracion.menu_opcion_padre IS 'Tabla que sirve para reemplazar la relacion reflexiva del padre de los menus_opcion.';")
	m.SQL("COMMENT ON COLUMN configuracion.menu_opcion_padre.id IS 'Identificador de la tabla menu_opcion_padre';")
	m.SQL("COMMENT ON COLUMN configuracion.menu_opcion_padre.padre IS 'Campo que contiene el id del padre de un menu.';")
	m.SQL("COMMENT ON COLUMN configuracion.menu_opcion_padre.hijo IS 'Campo que contiene el ID del menu hijo.';")
	m.SQL("COMMENT ON CONSTRAINT PK_MENU_OPCION_PADRE ON configuracion.menu_opcion_padre IS 'Llave primaria de la tabla menu_opcion_padre';")
}

// Reverse the migrations
func (m *CrearTablaMenuOpcionPadre_20191114_143306) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS configuracion.menu_opcion_padre")
}
