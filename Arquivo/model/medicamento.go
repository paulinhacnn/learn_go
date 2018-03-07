package model

type Medicamento struct {
	Empresa          string `json:"empresa"`
	Tempo            int    `json:"tempo"`
	Faturamento      int    `json:"empresa"`
	Invertimento_PeD int    `json:"investimento_ped"`
	Func_total       int    `json:"func_total"`
	Func_total_PeD   int    `json:"FunctotalPeD"`
	Func_Dr_PeD      int    `json:"FuncDrPeD"`
	Func_Ms_PeD      int    `json:"FuncMsPeD"`
	Func_Grad_PeD    int    `json:"FuncGradPeD"`
	Origem_Capital   int    `json:"origemCapital"`
	Fonte_Financ     int    `json:"FonteFinanc"`
	Produtos         int    `json:"Produtos"`
	Tipos            string `json:"Tipos"`
	Prod_Lancados    int    `json:"Prod_Lancados"`
}
