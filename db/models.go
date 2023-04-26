package db

type IvaCompras struct {
	cuit_cliente                                    string
	Fecha_de_Emisión                                string
	Tipo_de_Comprobante                             string
	Punto_de_Venta                                  string
	Número_de_Comprobante                           string
	Nro_Doc_Vendedor                                string
	Denominación_Vendedor                           string
	Importe_de_Percepciones_de_Ingresos_Brutos      int32
	Importe_de_Percepciones_o_Pagos_a_Cuenta_de_IVA int32
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
	Total_Verificador                               int32
}

type Users struct {
	cuit        string
	razonSocial string
}
