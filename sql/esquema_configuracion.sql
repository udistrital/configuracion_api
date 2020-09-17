-- Creación de esquema Configuración
CREATE SCHEMA configuracion;

-- Creación de secuencia y tabla Aplicacion
CREATE SEQUENCE configuracion.aplicacion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE configuracion.aplicacion
(
    id integer NOT NULL DEFAULT nextval('configuracion.aplicacion_id_seq'::regclass),
    nombre character varying(100) NOT NULL,
    descripcion text NOT NULL,
    dominio character varying(200) NOT NULL,
    estado boolean DEFAULT true,
    alias character varying(100),
    estilo_icono character varying(100),
    CONSTRAINT "PK_APLICACION" PRIMARY KEY (id),
    CONSTRAINT uq_aplicacion_dominio UNIQUE (dominio),
    CONSTRAINT uq_aplicacion_nombre UNIQUE (nombre)
);

COMMENT ON TABLE configuracion.aplicacion IS 'Tabla que se encarga de almacenar la informacion de cada aplicacion contenida en la base de datos';
COMMENT ON COLUMN configuracion.aplicacion.id IS 'Identificador unico de la tabla aplicacion';
COMMENT ON COLUMN configuracion.aplicacion.nombre IS 'Nombre que identifica al aplicativo';
COMMENT ON COLUMN configuracion.aplicacion.descripcion IS 'Descripcion detallada de las funcionalidades del aplicativo';
COMMENT ON COLUMN configuracion.aplicacion.dominio IS 'Dominio o url de acceso publico del aplicativo';
COMMENT ON COLUMN configuracion.aplicacion.alias IS 'Dato que se presentará a cambio del nombre de la aplicación';
COMMENT ON COLUMN configuracion.aplicacion.estilo_icono IS 'Valor de la clase css que referenciará el ícono';
COMMENT ON CONSTRAINT "PK_APLICACION" ON configuracion.aplicacion IS 'Llave primaria de la tabla aplicacion';


-- Creación de secuencia y tabla menu_opcion
CREATE SEQUENCE configuracion.menu_opcion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE configuracion.menu_opcion
(
    id integer NOT NULL DEFAULT nextval('configuracion.menu_opcion_id_seq'::regclass),
    nombre character varying(60) NOT NULL,
    descripcion character varying(250)  NOT NULL,
    url character varying(250)  NOT NULL,
    layout character varying(60),
    aplicacion integer NOT NULL,
    tipo_opcion character varying,
    icono character varying(100),
    CONSTRAINT "PK_MENU_OPCION" PRIMARY KEY (id),
    CONSTRAINT "FK_MENU_OPCION_APP" FOREIGN KEY (aplicacion) REFERENCES configuracion.aplicacion (id),      
    CONSTRAINT "CHECK_TIPO_OPCION" CHECK (tipo_opcion::text = ANY (ARRAY['Menú'::character varying::text, 'Acción'::character varying::text, 'Botón'::character varying::text]))
);   

COMMENT ON TABLE configuracion.menu_opcion IS 'Tabla que contiene las diferentes opciones de los menus';
COMMENT ON COLUMN configuracion.menu_opcion.id IS 'Identificador de la tabla menu_opcion';
COMMENT ON COLUMN configuracion.menu_opcion.nombre IS 'Contiene el nombre de la opcion del menu';
COMMENT ON COLUMN configuracion.menu_opcion.descripcion IS 'Contiene la informacion de los menus desplegables';
COMMENT ON COLUMN configuracion.menu_opcion.url IS 'Url o ruta del menu';
COMMENT ON COLUMN configuracion.menu_opcion.aplicacion IS 'Llave foranea de la tabla aplicacion';
COMMENT ON COLUMN configuracion.menu_opcion.tipo_opcion IS 'Clasificación de Tipo opción';
COMMENT ON CONSTRAINT "PK_MENU_OPCION" ON configuracion.menu_opcion IS 'Llave primaria de la tabla menu_opcion';
COMMENT ON CONSTRAINT "CHECK_TIPO_OPCION" ON configuracion.menu_opcion IS 'Restringe el tipo de opción que se debe tener';

-- Creación de secuencia y tabla menu_opcion_padre

CREATE SEQUENCE configuracion.menu_opcion_padre_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE configuracion.menu_opcion_padre
(
    id integer NOT NULL DEFAULT nextval('configuracion.menu_opcion_padre_id_seq'::regclass),
    padre integer NOT NULL,
    hijo integer,
    CONSTRAINT "PK_MENU_OPCION_PADRE" PRIMARY KEY (id),
    CONSTRAINT "UQ_PADRE_HIJO" UNIQUE (padre, hijo),   
    CONSTRAINT "FK_MENU_OPCION_HIJO_MENU_OPCION" FOREIGN KEY (hijo) REFERENCES configuracion.menu_opcion (id),
    CONSTRAINT "FK_MENU_OPCION_PADRE_MENU_OPCION" FOREIGN KEY (padre) REFERENCES configuracion.menu_opcion (id)
    
);

COMMENT ON TABLE configuracion.menu_opcion_padre IS 'Tabla que sirve para reemplazar la relacion reflexiva del padre de los menus_opcion.';
COMMENT ON COLUMN configuracion.menu_opcion_padre.id IS 'Identificador de la tabla menu_opcion_padre';
COMMENT ON COLUMN configuracion.menu_opcion_padre.padre IS 'Campo que contiene el id del padre de un menu.';
COMMENT ON COLUMN configuracion.menu_opcion_padre.hijo IS 'Campo que contiene el ID del menu hijo.';
COMMENT ON CONSTRAINT "PK_MENU_OPCION_PADRE" ON configuracion.menu_opcion_padre IS 'Llave primaria de la tabla menu_opcion_padre';

-- Creación de secuencia y tabla metodo_http

CREATE SEQUENCE configuracion.metodo_http_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
    
CREATE TABLE configuracion.metodo_http
(
    id integer NOT NULL DEFAULT nextval('configuracion.metodo_http_id_seq'::regclass),
    nombre character varying NOT NULL,
    descripcion character varying NOT NULL,
    CONSTRAINT pk_metodo_http PRIMARY KEY (id)
);

COMMENT ON TABLE configuracion.metodo_http IS 'Tipos de metodo http usados';
COMMENT ON COLUMN configuracion.metodo_http.id IS 'Identificador único de la tabla';
COMMENT ON COLUMN configuracion.metodo_http.nombre IS 'Identificador único de la configuracion generada';
COMMENT ON COLUMN configuracion.metodo_http.descripcion IS 'Identificador único de la configuracion generada';

-- Creación de secuencia y tabla configuracion.notificacion_tipo

CREATE SEQUENCE configuracion.notificacion_tipo_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;


CREATE TABLE configuracion.notificacion_tipo
(
    id integer NOT NULL DEFAULT nextval('configuracion.notificacion_tipo_id_seq'::regclass),
    nombre character varying(100),
    CONSTRAINT "PK_notificacion_tipo" PRIMARY KEY (id)
);

COMMENT ON TABLE configuracion.notificacion_tipo IS 'Centralización de los tipo de las notificaciones de los sistemas';
COMMENT ON COLUMN configuracion.notificacion_tipo.id IS 'Identificador único del tipo de la notificación';
COMMENT ON COLUMN configuracion.notificacion_tipo.nombre IS 'Nombre del tipo de la notificación';

-- Creación de secuencia y tabla notificacion_configuracion

CREATE SEQUENCE configuracion.notificacion_configuracion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE configuracion.notificacion_configuracion
(
    id integer NOT NULL DEFAULT nextval('configuracion.notificacion_configuracion_id_seq'::regclass),
    end_point character varying NOT NULL,
    metodo_http integer NOT NULL,
    tipo integer NOT NULL,
    cuerpo_notificacion json NOT NULL,
    aplicacion integer NOT NULL,
    CONSTRAINT pk_notificacion_configuracion PRIMARY KEY (id),
    CONSTRAINT fk_configuracion_notificacion_metodo_http FOREIGN KEY (metodo_http) REFERENCES configuracion.metodo_http (id),
    CONSTRAINT fk_notificacion_configuracion_aplicacion FOREIGN KEY (aplicacion) REFERENCES configuracion.aplicacion (id),
    CONSTRAINT fk_notificacion_configuracion_notificacion_tipo FOREIGN KEY (tipo) REFERENCES configuracion.notificacion_tipo (id) 
);

COMMENT ON TABLE configuracion.notificacion_configuracion IS 'Configuracion de las notificaciones que seran emitidas';
COMMENT ON COLUMN configuracion.notificacion_configuracion.id IS 'Identificador único de la configuracion generada';
COMMENT ON COLUMN configuracion.notificacion_configuracion.end_point IS 'end_point del api que genera la notificacion';
COMMENT ON COLUMN configuracion.notificacion_configuracion.metodo_http IS 'Metodo http de la peticion que genera la notificación';
COMMENT ON COLUMN configuracion.notificacion_configuracion.tipo IS 'tipo de notificación';
COMMENT ON COLUMN configuracion.notificacion_configuracion.cuerpo_notificacion IS 'cuerpo de la notificacion (configuracion para mostrar en cliente y plantilla de mensaje)';
COMMENT ON COLUMN configuracion.notificacion_configuracion.aplicacion IS 'Aplicacion que genera la notificacion';

-- Creación de secuencia y tabla notificacion

CREATE SEQUENCE configuracion.notificacion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE configuracion.notificacion
(
    id integer NOT NULL DEFAULT nextval('configuracion.notificacion_id_seq'::regclass),
    fecha_creacion timestamp with time zone NOT NULL DEFAULT now(),
    cuerpo_notificacion json,
    notificacion_configuracion integer NOT NULL,
    CONSTRAINT "PK_notificacion" PRIMARY KEY (id),
    CONSTRAINT fk_notificacion_notificacion_configuracion FOREIGN KEY (notificacion_configuracion) REFERENCES configuracion.notificacion_configuracion (id) 
);    

COMMENT ON TABLE configuracion.notificacion IS 'Notificaciones de sistema a usuario centralizadas para todos los sistemas';
COMMENT ON COLUMN configuracion.notificacion.id IS 'Identificador único de la notificación generada';
COMMENT ON COLUMN configuracion.notificacion.fecha_creacion IS 'Fecha y hora en la que se creó la notificación';
COMMENT ON COLUMN configuracion.notificacion.cuerpo_notificacion IS 'Cuerpo de la notificación, será un objeto JSON en el que se almacenará el mensaje, título y otra información relevante';
COMMENT ON COLUMN configuracion.notificacion.notificacion_configuracion IS 'Referencia a la configuracion de la notificacion generada';

-- Creación de secuencia y tabla perfil

CREATE SEQUENCE configuracion.perfil_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE configuracion.perfil
(
    id integer NOT NULL DEFAULT nextval('configuracion.perfil_id_seq'::regclass),
    nombre character varying(50) NOT NULL,
    aplicacion integer NOT NULL,
    CONSTRAINT "PK_PERFIL" PRIMARY KEY (id),
    CONSTRAINT "FK_PERFIL_APP" FOREIGN KEY (aplicacion) REFERENCES configuracion.aplicacion (id)
);

COMMENT ON TABLE configuracion.perfil IS 'Contiene la informacion de un perfil';
COMMENT ON COLUMN configuracion.perfil.id IS 'Identificador de un perfil';
COMMENT ON COLUMN configuracion.perfil.nombre IS 'Nombre del perfil';
COMMENT ON COLUMN configuracion.perfil.aplicacion IS 'Contiene el id de la aplicacion asociada al perfil';
COMMENT ON CONSTRAINT "PK_PERFIL" ON configuracion.perfil IS 'llave primaria de la tabla perfil';

CREATE SEQUENCE configuracion.perfil_x_menu_opcion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE configuracion.perfil_x_menu_opcion
(
    id integer NOT NULL DEFAULT nextval('configuracion.perfil_x_menu_opcion_id_seq'::regclass),
    perfil integer NOT NULL,
    opcion integer NOT NULL,
    CONSTRAINT "PK_PERFIL_X_MENU_OPCION" PRIMARY KEY (id),
    CONSTRAINT "UQ_PERFIL_X_MENU" UNIQUE (perfil, opcion),
    CONSTRAINT "FK_OPCION" FOREIGN KEY (opcion) REFERENCES configuracion.menu_opcion (id), 
    CONSTRAINT "FK_PERFIL" FOREIGN KEY (perfil) REFERENCES configuracion.perfil (id) 
);

COMMENT ON TABLE configuracion.perfil_x_menu_opcion IS 'Tabla que maneja los menus que le pertenecen a ciertos perfiles';
COMMENT ON COLUMN configuracion.perfil_x_menu_opcion.id IS 'Identificador de la tabla perfil_x_menu_opcion';
COMMENT ON COLUMN configuracion.perfil_x_menu_opcion.perfil IS 'Contiene la informacion del perfil';
COMMENT ON COLUMN configuracion.perfil_x_menu_opcion.opcion IS 'Contiene la informacion del menu_opcion';
COMMENT ON CONSTRAINT "PK_PERFIL_X_MENU_OPCION" ON configuracion.perfil_x_menu_opcion IS 'Llave primera de la tabla de rompimiento';
COMMENT ON CONSTRAINT "UQ_PERFIL_X_MENU" ON configuracion.perfil_x_menu_opcion IS 'Restricción que solo deja asociar un registro de menú por perfil, garantiza que no haya replica.';

-- Creación de secuencia y tabla notificacion_configuracion_perfil

CREATE SEQUENCE configuracion.notificacion_configuracion_perfil_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE configuracion.notificacion_configuracion_perfil
(
    id integer NOT NULL DEFAULT nextval('configuracion.notificacion_configuracion_perfil_id_seq'::regclass),
    notificacion_configuracion integer NOT NULL,
    perfil integer NOT NULL,
    CONSTRAINT pk_notificacion_configuracion_perfil PRIMARY KEY (id),
    CONSTRAINT fk_notificacion_configuracion_perfil_notificacion_configuracion FOREIGN KEY (notificacion_configuracion) REFERENCES configuracion.notificacion_configuracion (id),
    CONSTRAINT fk_notificacion_configuracion_perfil_perfil FOREIGN KEY (perfil)  REFERENCES configuracion.perfil (id) 
);

COMMENT ON TABLE configuracion.notificacion_configuracion_perfil IS 'Tipos de metodo http usados';
COMMENT ON COLUMN configuracion.notificacion_configuracion_perfil.id IS 'Identificador único de la tabla';
COMMENT ON COLUMN configuracion.notificacion_configuracion_perfil.notificacion_configuracion IS 'Referencia a la configuracion de la notificacion';
COMMENT ON COLUMN configuracion.notificacion_configuracion_perfil.perfil IS 'referencia al perfil que debe recibir la notificacion';

-- Creación de secuencia y tabla notificacion_estado

CREATE SEQUENCE configuracion.notificacion_estado_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE configuracion.notificacion_estado
(
    id integer NOT NULL DEFAULT nextval('configuracion.notificacion_estado_id_seq'::regclass),
    nombre character varying(100),
    codigo_abreviacion character varying(20),
    descripcion character varying(250),
    activo boolean DEFAULT true,
    numero_orden numeric(5,2),
    CONSTRAINT "PK_notificacion_estado" PRIMARY KEY (id)
);

COMMENT ON TABLE configuracion.notificacion_estado IS 'Centralización de los estados de las notificaciones de los sistemas';
COMMENT ON COLUMN configuracion.notificacion_estado.id IS 'Identificador único del estado de la notificación';
COMMENT ON COLUMN configuracion.notificacion_estado.nombre IS 'Nombre del estado de la notificación';
COMMENT ON COLUMN configuracion.notificacion_estado.codigo_abreviacion IS 'Código de abreviación del estado de la notificación';
COMMENT ON COLUMN configuracion.notificacion_estado.descripcion IS 'Descripción del estado de notificación';
COMMENT ON COLUMN configuracion.notificacion_estado.activo IS 'indicador de si el registro se encuentra en estado activo';
COMMENT ON COLUMN configuracion.notificacion_estado.numero_orden IS 'Orden de los estados de notificación';

-- Creación de secuencia y tabla notificacion_estado_usuario

CREATE SEQUENCE configuracion.notificacion_estado_usuario_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE configuracion.notificacion_estado_usuario
(
    id integer NOT NULL DEFAULT nextval('configuracion.notificacion_estado_usuario_id_seq'::regclass),
    notificacion integer NOT NULL,
    notificacion_estado integer NOT NULL,
    fecha timestamp without time zone NOT NULL,
    usuario character varying(100),
    activo boolean DEFAULT true,
    CONSTRAINT pk_notificacion_estado_usuario PRIMARY KEY (id),
    CONSTRAINT fk_notificacion_estado_usuario_notificacion FOREIGN KEY (notificacion) REFERENCES configuracion.notificacion (id),
    CONSTRAINT fk_notificacion_estado_usuario_notificacion_estado FOREIGN KEY (notificacion_estado) REFERENCES configuracion.notificacion_estado (id)
);

COMMENT ON TABLE configuracion.notificacion_estado_usuario IS 'Tabla que relaciona una notificación, un estado y un usuario. Permite así conservar el histórico de la notificación desde que se crea hasta que es leída';
COMMENT ON COLUMN configuracion.notificacion_estado_usuario.id IS 'Identificador único de notificacion_estado_usuario';
COMMENT ON COLUMN configuracion.notificacion_estado_usuario.notificacion IS 'Id de la notificación registrada';
COMMENT ON COLUMN configuracion.notificacion_estado_usuario.notificacion_estado IS 'Estado de la notificación registrada';
COMMENT ON COLUMN configuracion.notificacion_estado_usuario.fecha IS 'Fecha de la notificación ';
COMMENT ON COLUMN configuracion.notificacion_estado_usuario.usuario IS 'Usuario de WSO2 que interactúa con la notificación registrada';
COMMENT ON COLUMN configuracion.notificacion_estado_usuario.activo IS 'indicador de si el registro se encuentra en estado activo';

-- Creación de secuencia y tabla configuracion.parametro

CREATE SEQUENCE configuracion.parametro_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE configuracion.parametro
(
    id integer NOT NULL DEFAULT nextval('configuracion.parametro_id_seq'::regclass),
    nombre character varying(200) NOT NULL,
    valor text NOT NULL,
    aplicacion integer NOT NULL,
    CONSTRAINT "PK_PARAMETRO" PRIMARY KEY (id),
    CONSTRAINT "FK_PARAMETRO_APP" FOREIGN KEY (aplicacion) REFERENCES configuracion.aplicacion (id)
);

COMMENT ON TABLE configuracion.parametro IS 'Tabla que  describe todos los parametros asociados a una aplicacion';
COMMENT ON COLUMN configuracion.parametro.id IS 'Identificador de la tabla parametro';
COMMENT ON COLUMN configuracion.parametro.nombre IS 'Nombre o llave del parametro ';
COMMENT ON COLUMN configuracion.parametro.valor IS 'Valor del parametro';
COMMENT ON COLUMN configuracion.parametro.aplicacion IS 'Id de la aplicación';
COMMENT ON CONSTRAINT "PK_PARAMETRO" ON configuracion.parametro IS 'Llave primaria del parametro';