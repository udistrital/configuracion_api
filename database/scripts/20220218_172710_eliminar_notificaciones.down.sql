-- Diff code generated with pgModeler (PostgreSQL Database Modeler)
-- pgModeler version: 0.9.1-beta
-- Diff date: 2022-02-18 20:39:39
-- Source model: test
-- Database: test
-- PostgreSQL version: 10.0

-- [ Diff summary ]
-- Dropped objects: 0
-- Created objects: 23
-- Changed objects: 16
-- Truncated tables: 0

SET search_path=public,configuracion;
-- ddl-end --


-- [ Created objects ] --
-- object: configuracion.metodo_http_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS configuracion.metodo_http_id_seq CASCADE;
CREATE SEQUENCE configuracion.metodo_http_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE;
	-- OWNED BY NONE;
-- ddl-end --
-- ALTER SEQUENCE configuracion.metodo_http_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: configuracion.metodo_http | type: TABLE --
-- DROP TABLE IF EXISTS configuracion.metodo_http CASCADE;
CREATE TABLE configuracion.metodo_http(
	id integer NOT NULL DEFAULT nextval('configuracion.metodo_http_id_seq'::regclass),
	nombre character varying NOT NULL,
	descripcion character varying NOT NULL,
	CONSTRAINT pk_metodo_http PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE configuracion.metodo_http IS 'Tipos de metodo http usados';
-- ddl-end --
COMMENT ON COLUMN configuracion.metodo_http.id IS 'Identificador único de la tabla';
-- ddl-end --
COMMENT ON COLUMN configuracion.metodo_http.nombre IS 'Identificador único de la configuracion generada';
-- ddl-end --
COMMENT ON COLUMN configuracion.metodo_http.descripcion IS 'Identificador único de la configuracion generada';
-- ddl-end --
-- ALTER TABLE configuracion.metodo_http OWNER TO desarrollooas;
-- ddl-end --

-- object: configuracion.notificacion_tipo_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS configuracion.notificacion_tipo_id_seq CASCADE;
CREATE SEQUENCE configuracion.notificacion_tipo_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE;
	-- OWNED BY NONE;
-- ddl-end --
-- ALTER SEQUENCE configuracion.notificacion_tipo_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: configuracion.notificacion_tipo | type: TABLE --
-- DROP TABLE IF EXISTS configuracion.notificacion_tipo CASCADE;
CREATE TABLE configuracion.notificacion_tipo(
	id integer NOT NULL DEFAULT nextval('configuracion.notificacion_tipo_id_seq'::regclass),
	nombre character varying(100),
	CONSTRAINT pk_notificacion_tipo PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE configuracion.notificacion_tipo IS 'Centralización de los tipo de las notificaciones de los sistemas';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion_tipo.id IS 'Identificador único del tipo de la notificación';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion_tipo.nombre IS 'Nombre del tipo de la notificación';
-- ddl-end --
-- ALTER TABLE configuracion.notificacion_tipo OWNER TO desarrollooas;
-- ddl-end --

-- object: configuracion.notificacion_configuracion_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS configuracion.notificacion_configuracion_id_seq CASCADE;
CREATE SEQUENCE configuracion.notificacion_configuracion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE;
	-- OWNED BY NONE;
-- ddl-end --
-- ALTER SEQUENCE configuracion.notificacion_configuracion_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: configuracion.notificacion_configuracion | type: TABLE --
-- DROP TABLE IF EXISTS configuracion.notificacion_configuracion CASCADE;
CREATE TABLE configuracion.notificacion_configuracion(
	id integer NOT NULL DEFAULT nextval('configuracion.notificacion_configuracion_id_seq'::regclass),
	end_point character varying NOT NULL,
	metodo_http integer NOT NULL,
	tipo integer NOT NULL,
	cuerpo_notificacion json NOT NULL,
	aplicacion integer NOT NULL,
	CONSTRAINT pk_notificacion_configuracion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE configuracion.notificacion_configuracion IS 'Configuracion de las notificaciones que seran emitidas';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion_configuracion.id IS 'Identificador único de la configuracion generada';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion_configuracion.end_point IS 'end_point del api que genera la notificacion';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion_configuracion.metodo_http IS 'Metodo http de la peticion que genera la notificación';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion_configuracion.tipo IS 'tipo de notificación';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion_configuracion.cuerpo_notificacion IS 'cuerpo de la notificacion (configuracion para mostrar en cliente y plantilla de mensaje)';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion_configuracion.aplicacion IS 'Aplicacion que genera la notificacion';
-- ddl-end --
-- ALTER TABLE configuracion.notificacion_configuracion OWNER TO desarrollooas;
-- ddl-end --

-- object: configuracion.notificacion_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS configuracion.notificacion_id_seq CASCADE;
CREATE SEQUENCE configuracion.notificacion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE;
	-- OWNED BY NONE;
-- ddl-end --
-- ALTER SEQUENCE configuracion.notificacion_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: configuracion.notificacion | type: TABLE --
-- DROP TABLE IF EXISTS configuracion.notificacion CASCADE;
CREATE TABLE configuracion.notificacion(
	id integer NOT NULL DEFAULT nextval('configuracion.notificacion_id_seq'::regclass),
	fecha_creacion timestamp with time zone NOT NULL DEFAULT now(),
	cuerpo_notificacion json,
	notificacion_configuracion integer NOT NULL,
	CONSTRAINT pk_notificacion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE configuracion.notificacion IS 'Notificaciones de sistema a usuario centralizadas para todos los sistemas';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion.id IS 'Identificador único de la notificación generada';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion.fecha_creacion IS 'Fecha y hora en la que se creó la notificación';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion.cuerpo_notificacion IS 'Cuerpo de la notificación, será un objeto JSON en el que se almacenará el mensaje, título y otra información relevante';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion.notificacion_configuracion IS 'Referencia a la configuracion de la notificacion generada';
-- ddl-end --
-- ALTER TABLE configuracion.notificacion OWNER TO desarrollooas;
-- ddl-end --

-- object: configuracion.notificacion_configuracion_perfil_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS configuracion.notificacion_configuracion_perfil_id_seq CASCADE;
CREATE SEQUENCE configuracion.notificacion_configuracion_perfil_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE;
	-- OWNED BY NONE;
-- ddl-end --
-- ALTER SEQUENCE configuracion.notificacion_configuracion_perfil_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: configuracion.notificacion_configuracion_perfil | type: TABLE --
-- DROP TABLE IF EXISTS configuracion.notificacion_configuracion_perfil CASCADE;
CREATE TABLE configuracion.notificacion_configuracion_perfil(
	id integer NOT NULL DEFAULT nextval('configuracion.notificacion_configuracion_perfil_id_seq'::regclass),
	notificacion_configuracion integer NOT NULL,
	perfil integer NOT NULL,
	CONSTRAINT pk_notificacion_configuracion_perfil PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE configuracion.notificacion_configuracion_perfil IS 'Tipos de metodo http usados';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion_configuracion_perfil.id IS 'Identificador único de la tabla';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion_configuracion_perfil.notificacion_configuracion IS 'Referencia a la configuracion de la notificacion';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion_configuracion_perfil.perfil IS 'referencia al perfil que debe recibir la notificacion';
-- ddl-end --
-- ALTER TABLE configuracion.notificacion_configuracion_perfil OWNER TO desarrollooas;
-- ddl-end --

-- object: configuracion.notificacion_estado_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS configuracion.notificacion_estado_id_seq CASCADE;
CREATE SEQUENCE configuracion.notificacion_estado_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE;
	-- OWNED BY NONE;
-- ddl-end --
-- ALTER SEQUENCE configuracion.notificacion_estado_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: configuracion.notificacion_estado | type: TABLE --
-- DROP TABLE IF EXISTS configuracion.notificacion_estado CASCADE;
CREATE TABLE configuracion.notificacion_estado(
	id integer NOT NULL DEFAULT nextval('configuracion.notificacion_estado_id_seq'::regclass),
	nombre character varying(100),
	codigo_abreviacion character varying(20),
	descripcion character varying(250),
	activo boolean DEFAULT true,
	numero_orden numeric(5,2),
	CONSTRAINT pk_notificacion_estado PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE configuracion.notificacion_estado IS 'Centralización de los estados de las notificaciones de los sistemas';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion_estado.id IS 'Identificador único del estado de la notificación';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion_estado.nombre IS 'Nombre del estado de la notificación';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion_estado.codigo_abreviacion IS 'Código de abreviación del estado de la notificación';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion_estado.descripcion IS 'Descripción del estado de notificación';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion_estado.activo IS 'indicador de si el registro se encuentra en estado activo';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion_estado.numero_orden IS 'Orden de los estados de notificación';
-- ddl-end --
-- ALTER TABLE configuracion.notificacion_estado OWNER TO desarrollooas;
-- ddl-end --

-- object: configuracion.notificacion_estado_usuario_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS configuracion.notificacion_estado_usuario_id_seq CASCADE;
CREATE SEQUENCE configuracion.notificacion_estado_usuario_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE;
	-- OWNED BY NONE;
-- ddl-end --
-- ALTER SEQUENCE configuracion.notificacion_estado_usuario_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: configuracion.notificacion_estado_usuario | type: TABLE --
-- DROP TABLE IF EXISTS configuracion.notificacion_estado_usuario CASCADE;
CREATE TABLE configuracion.notificacion_estado_usuario(
	id integer NOT NULL DEFAULT nextval('configuracion.notificacion_estado_usuario_id_seq'::regclass),
	notificacion integer NOT NULL,
	notificacion_estado integer NOT NULL,
	fecha timestamp NOT NULL,
	usuario character varying(100),
	activo boolean DEFAULT true,
	CONSTRAINT pk_notificacion_estado_usuario PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE configuracion.notificacion_estado_usuario IS 'Tabla que relaciona una notificación, un estado y un usuario. Permite así conservar el histórico de la notificación desde que se crea hasta que es leída';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion_estado_usuario.id IS 'Identificador único de notificacion_estado_usuario';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion_estado_usuario.notificacion IS 'Id de la notificación registrada';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion_estado_usuario.notificacion_estado IS 'Estado de la notificación registrada';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion_estado_usuario.fecha IS 'Fecha de la notificación ';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion_estado_usuario.usuario IS 'Usuario de WSO2 que interactúa con la notificación registrada';
-- ddl-end --
COMMENT ON COLUMN configuracion.notificacion_estado_usuario.activo IS 'indicador de si el registro se encuentra en estado activo';
-- ddl-end --
-- ALTER TABLE configuracion.notificacion_estado_usuario OWNER TO desarrollooas;
-- ddl-end --


-- [ Created foreign keys ] --
-- object: fk_configuracion_notificacion_metodo_http | type: CONSTRAINT --
-- ALTER TABLE configuracion.notificacion_configuracion DROP CONSTRAINT IF EXISTS fk_configuracion_notificacion_metodo_http CASCADE;
ALTER TABLE configuracion.notificacion_configuracion ADD CONSTRAINT fk_configuracion_notificacion_metodo_http FOREIGN KEY (metodo_http)
REFERENCES configuracion.metodo_http (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_notificacion_configuracion_aplicacion | type: CONSTRAINT --
-- ALTER TABLE configuracion.notificacion_configuracion DROP CONSTRAINT IF EXISTS fk_notificacion_configuracion_aplicacion CASCADE;
ALTER TABLE configuracion.notificacion_configuracion ADD CONSTRAINT fk_notificacion_configuracion_aplicacion FOREIGN KEY (aplicacion)
REFERENCES configuracion.aplicacion (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_notificacion_configuracion_notificacion_tipo | type: CONSTRAINT --
-- ALTER TABLE configuracion.notificacion_configuracion DROP CONSTRAINT IF EXISTS fk_notificacion_configuracion_notificacion_tipo CASCADE;
ALTER TABLE configuracion.notificacion_configuracion ADD CONSTRAINT fk_notificacion_configuracion_notificacion_tipo FOREIGN KEY (tipo)
REFERENCES configuracion.notificacion_tipo (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_notificacion_notificacion_configuracion | type: CONSTRAINT --
-- ALTER TABLE configuracion.notificacion DROP CONSTRAINT IF EXISTS fk_notificacion_notificacion_configuracion CASCADE;
ALTER TABLE configuracion.notificacion ADD CONSTRAINT fk_notificacion_notificacion_configuracion FOREIGN KEY (notificacion_configuracion)
REFERENCES configuracion.notificacion_configuracion (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_notificacion_configuracion_perfil_notificacion_configuracion | type: CONSTRAINT --
-- ALTER TABLE configuracion.notificacion_configuracion_perfil DROP CONSTRAINT IF EXISTS fk_notificacion_configuracion_perfil_notificacion_configuracion CASCADE;
ALTER TABLE configuracion.notificacion_configuracion_perfil ADD CONSTRAINT fk_notificacion_configuracion_perfil_notificacion_configuracion FOREIGN KEY (notificacion_configuracion)
REFERENCES configuracion.notificacion_configuracion (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_notificacion_configuracion_perfil_perfil | type: CONSTRAINT --
-- ALTER TABLE configuracion.notificacion_configuracion_perfil DROP CONSTRAINT IF EXISTS fk_notificacion_configuracion_perfil_perfil CASCADE;
ALTER TABLE configuracion.notificacion_configuracion_perfil ADD CONSTRAINT fk_notificacion_configuracion_perfil_perfil FOREIGN KEY (perfil)
REFERENCES configuracion.perfil (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_notificacion_estado_usuario_notificacion | type: CONSTRAINT --
-- ALTER TABLE configuracion.notificacion_estado_usuario DROP CONSTRAINT IF EXISTS fk_notificacion_estado_usuario_notificacion CASCADE;
ALTER TABLE configuracion.notificacion_estado_usuario ADD CONSTRAINT fk_notificacion_estado_usuario_notificacion FOREIGN KEY (notificacion)
REFERENCES configuracion.notificacion (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_notificacion_estado_usuario_notificacion_estado | type: CONSTRAINT --
-- ALTER TABLE configuracion.notificacion_estado_usuario DROP CONSTRAINT IF EXISTS fk_notificacion_estado_usuario_notificacion_estado CASCADE;
ALTER TABLE configuracion.notificacion_estado_usuario ADD CONSTRAINT fk_notificacion_estado_usuario_notificacion_estado FOREIGN KEY (notificacion_estado)
REFERENCES configuracion.notificacion_estado (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

