/* "Fecha de Emisión"
"Tipo de Comprobante"
"Punto de Venta"
"Número de Comprobante"
"Tipo Doc. Vendedor"
"Nro. Doc. Vendedor"
"Denominación Vendedor"
"Importe Total"
"Moneda Original"
"Tipo de Cambio"
"Importe No Gravado"
"Importe Exento"
"Crédito Fiscal Computable"
"Importe de Per. o Pagos a Cta. de Otros Imp. Nac."
"Importe de Percepciones de Ingresos Brutos"
"Importe de Impuestos Municipales"
"Importe de Percepciones o Pagos a Cuenta de IVA"
"Importe de Impuestos Internos"
"Importe Otros Tributos"
"Neto Gravado IVA 0%"
"Neto Gravado IVA 2,5%"
"Importe IVA 2,5%"
"Neto Gravado IVA 5%"
"Importe IVA 5%"
"Neto Gravado IVA 10,5%"
"Importe IVA 10,5%"
"Neto Gravado IVA 21%"
"Importe IVA 21%"
"Neto Gravado IVA 27%"
"Importe IVA 27%"
"Total Neto Gravado"
"Total IVA" */

package db

type IvaCompras struct {
	Cuit_clientes                                   string
	Fecha_de_Emision                                string
	Mes_de_imputacion                               string
	Tipo_de_Comprobante                             string
	Punto_de_Venta                                  string
	Número_de_Comprobante                           string
	Tipo_Doc_Vendedor                               string
	Nro_Doc_Vendedor                                string
	Denominación_Vendedor                           string
	Importe_Total                                   int32
	Moneda_Original                                 string
	Tipo_de_Cambio                                  string
	Importe_No_Gravado                              int32
	Importe_Exento                                  int32
	Crédito_Fiscal_Computable                       int32
	Importe_de_Per_o_Pagos_a_Cta_de_Otros_Imp_Nac   int32
	Importe_de_Percepciones_de_Ingresos_Brutos      int32
	Importe_de_Impuestos_Municipales                int32
	Importe_de_Percepciones_o_Pagos_a_Cuenta_de_IVA int32
	Importe_de_Impuestos_Internos                   int32
	Importe_Otros_Tributos                          int32
	Neto_Gravado_IVA_0                              int32
	Neto_Gravado_IVA_25                             int32
	Importe_IVA_25                                  int32
	Neto_Gravado_IVA_5                              int32
	Importe_IVA_5                                   int32
	Neto_Gravado_IVA_105                            int32
	Importe_IVA_105                                 int32
	Neto_Gravado_IVA_21                             int32
	Importe_IVA_21                                  int32
	Neto_Gravado_IVA_27                             int32
	Importe_IVA_27                                  int32
	Total_Neto_Gravado                              int32
	Total_IVA                                       int32
}

type Clientes struct {
	Cuit        string `db:"cuit" json:"cuit"`
	RazonSocial string `db:"razon_social" json:"razonSocial"`
}
