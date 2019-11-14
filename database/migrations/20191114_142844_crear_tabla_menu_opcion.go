package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaMenuOpcion_20191114_142844 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaMenuOpcion_20191114_142844{}
	m.Created = "20191114_142844"

	migration.Register("CrearTablaMenuOpcion_20191114_142844", m)
}

// Run the migrations
func (m *CrearTablaMenuOpcion_20191114_142844) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXIST configuracion.menu_opcion ( id serial NOT NULL, nombre character varying(60) NOT NULL, descripcion character varying(250) NOT NULL, url character varying(250) NOT NULL, layout character varying(60), aplicacion integer NOT NULL, tipo_opcion character varying, icono character varying(100), CONSTRAINT PK_MENU_OPCION PRIMARY KEY (id), CONSTRAINT FK_MENU_OPCION_APP FOREIGN KEY (aplicacion) REFERENCES configuracion.aplicacion (id), CONSTRAINT CHECK_TIPO_OPCION CHECK (tipo_opcion::text = ANY (ARRAY['Menú'::character varying::text, 'Acción'::character varying::text, 'Botón'::character varying::text])) );")
	m.SQL("ALTER TABLE configuracion.menu_opcion OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE configuracion.menu_opcion IS 'Tabla que contiene las diferentes opciones de los menus';")
	m.SQL("COMMENT ON COLUMN configuracion.menu_opcion.id IS 'Identificador de la tabla menu_opcion';")
	m.SQL("COMMENT ON COLUMN configuracion.menu_opcion.nombre IS 'Contiene el nombre de la opcion del menu';")
	m.SQL("COMMENT ON COLUMN configuracion.menu_opcion.descripcion IS 'Contiene la informacion de los menus desplegables';")
	m.SQL("COMMENT ON COLUMN configuracion.menu_opcion.url IS 'Url o ruta del menu';")
	m.SQL("COMMENT ON COLUMN configuracion.menu_opcion.aplicacion IS 'Llave foranea de la tabla aplicacion';")
	m.SQL("COMMENT ON COLUMN configuracion.menu_opcion.tipo_opcion IS 'Clasificación de Tipo opción';")
	m.SQL("COMMENT ON CONSTRAINT PK_MENU_OPCION ON configuracion.menu_opcion IS 'Llave primaria de la tabla menu_opcion';")
	m.SQL("COMMENT ON CONSTRAINT CHECK_TIPO_OPCION ON configuracion.menu_opcion IS 'Restringe el tipo de opción que se debe tener';")
}

// Reverse the migrations
func (m *CrearTablaMenuOpcion_20191114_142844) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS configuracion.menu_opcion")
}
