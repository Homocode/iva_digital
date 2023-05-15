package services

import "github.com/homocode/libro_iva_afip/db"

type comprobanteFront struct {
	fechaEmision               string
	mesImputacion              string
	tipoComprobante            string
	numeroComprobante          string
	nombreProveedor            string
	cuitProvedor               string
	importeTotal               int32
	toalNetosGravados          int32
	totalIva                   int32
	conceptosExentosNoGravados int32
	percepcionesIva            int32
	percepcionesIibb           int32
}

func FilterColumns(comprobante db.IvaCompras) comprobanteFront {

	cf := comprobanteFront{
		fechaEmision:               comprobante.Fecha_de_Emision,
		mesImputacion:              comprobante.Mes_de_imputacion,
		tipoComprobante:            comprobante.Tipo_de_Comprobante,
		numeroComprobante:          comprobante.Número_de_Comprobante,
		nombreProveedor:            comprobante.Denominación_Vendedor,
		cuitProvedor:               comprobante.Nro_Doc_Vendedor,
		importeTotal:               comprobante.Importe_Total,
		toalNetosGravados:          comprobante.Total_Neto_Gravado,
		totalIva:                   comprobante.Total_IVA,
		conceptosExentosNoGravados: 4,
		percepcionesIva:            5334,
		percepcionesIibb:           434,
	}

	return cf
}
