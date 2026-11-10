package model

type GroupReport struct {
	ID            int
	Institute     string
	Group_name    string
	Starosta      string
	Students_list int
	Students_fact int
	Missed_dln    int
	Quality_pct   float64
	Measured_at   string
	Starosta_was  string
}

type PersonalMetrics struct {
	ID      int     // id
	FIO     string  // fio
	Dadt    string  // Disadaptation
	Smar    float64 // smar
	Smap    float64 // smap
	Smsd    float64 // smsd
	PctSmar float64 // pct_smar
	PctSmap float64 // pct_smap
	Rvp     float64 // rvp
	Tpi     float64 // tpi
	Rvpd    float64 // rvpd
	Tpid    float64 // tpid
	Gdp     float64 // gdp
	Pss     float64 // pss
	Psa     float64 // psa
	Iis     float64 // iis
	Ipe     float64 // ipe
	L04     float64 // l0_4
	L12     float64 // l1_2
}
