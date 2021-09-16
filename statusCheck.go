package main

import (
	// "github.com/astaxie/beego/logs"
	"github.com/udistrital/configuracion_api/models"
)

func statusCheck() (statusError interface{}) {
	// A MEJORAR:
	// - Se podrían agregar más validaciones en este callback/handler
	// - Se podría considerar implementar un timeout

	// Se llama el mismo helper usado en GET/perfil
	// logs.Info("statusCheck")
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 1
	var offset int64
	if _, err := models.GetAllPerfil(query, fields, sortby, order, offset, limit); err != nil {
		// logs.Error(err)
		panic(err)
		// statusError = err // altenativa al panic(), funciona igual
	}
	return
}
