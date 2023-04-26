create table "comprabantes_compras_csv"(
    "Fecha_de_Emisión" text,
    "Tipo_de_Comprobante" text,
    "Punto_de_Venta" text,
    "Número_de_Comprobante" text,
    "Tipo_Doc._Vendedor" text,
    "Nro._Doc._Vendedor" text,
    "Denominación_Vendedor" text,
    "Importe_Total" text,
    "Moneda_Original" text,
    "Tipo_de_Cambio" text,
    "Importe_No_Gravado" text,
    "Importe_Exento" text,
    "Crédito_Fiscal_Computable" text,
    "Importe_de_Per._o_Pagos_a_Cta._de_Otros_Imp._Nac." text,
    "Importe_de_Percepciones_de_Ingresos_Brutos" text,
    "Importe_de_Impuestos_Municipales" text,
    "Importe_de_Percepciones_o_Pagos_a_Cuenta_de_IVA" text,
    "Importe_de_Impuestos_Internos" text,
    "Importe_Otros_Tributos" text,
    "Neto_Gravado_IVA_0%" text,
    "Neto_Gravado_IVA_2,5%" text, 
    "Importe_IVA_2,5%" text,
    "Neto_Gravado_IVA_5%" text,
    "Importe_IVA_5%" text,
    "Neto_Gravado_IVA_10,5%" text,
    "Importe_IVA_10,5%" text,
    "Neto_Gravado_IVA_21%" text,
    "Importe_IVA_21%" text,
    "Neto_Gravado_IVA_27%" text,
    "Importe_IVA_27%" text,
    "Total_Neto_Gravado" text,
    "Total_IVA" text
);

create table "iva_compras" (
    "cuit_cliente" varchar,
    "mes_imputación" varchar,
    like comprabantes_compras_csv including all
);

create table "clientes" (
    "cuit" varchar primary key,
    "razon_social" varchar
);

create index on "iva_compras" ("cuit_cliente");

create index on "clientes" ("cuit");

alter table "iva_compras" add foreign key ("cuit_cliente") references "clientes" ("cuit");