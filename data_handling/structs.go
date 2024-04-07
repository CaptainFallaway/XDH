package data_handling

import (
	"strconv"
)

type MetalValue struct {
	Value float64
	IsLod bool
}

func (n *MetalValue) UnmarshalCSV(val string) (err error) {
	if val, err := strconv.ParseFloat(val, 64); err == nil {
		n.Value = val
		n.IsLod = false
	} else {
		n.IsLod = true
	}
	return err
}

func (n *MetalValue) MarshalCSV() (string, error) {
	panic("Not implemented")
	// if n.IsLod {
	// 	return "<LOD", nil
	// } else {
	// 	return strconv.FormatFloat(n.Value, 'f', -1, 64), nil
	// }
}

// Index,Reading No,Time,Type,Duration,Units,Sigma Value,Sequence,User1,Flags,Boat,Operator,User Login,Pb,Pb Error,Hg,Hg Error,W,W Error,Rb,Rb Error,Se,Se Error,Zn,Zn Error,Ag,Ag Error,Bk1,Bk1 Error,Mo,Mo Error,Sr,Sr Error,As,As Error,Cu,Cu Error,Ni,Ni Error,Fe,Fe Error,Cr,Cr Error,Bk2,Bk2 Error,Bk3,Bk3 Error,Bk4,Bk4 Error,Sn,Sn Error,
type Scan struct {
	Index      uint16     `csv:"Index"`
	Reading    uint16     `csv:"Reading No"`
	Time       string     `csv:"Time"`
	Type       string     `csv:"Type"`
	Duration   float32    `csv:"Duration"`
	Units      string     `csv:"Units"`
	SigmaValue int16      `csv:"Sigma Value"`
	Sequence   string     `csv:"Sequence"`
	User1      string     `csv:"User1"`
	Flags      string     `csv:"Flags"`
	Boat       string     `csv:"Boat"`
	Operator   string     `csv:"Operator"`
	UserLogin  string     `csv:"User Login"`
	Pb         MetalValue `csv:"Pb"`
	PbError    float64    `csv:"Pb Error"`
	Zn         MetalValue `csv:"Zn"`
	ZnError    float64    `csv:"Zn Error"`
	Cu         MetalValue `csv:"Cu"`
	CuError    float64    `csv:"Cu Error"`
	Sn         MetalValue `csv:"Sn"`
	SnError    float64    `csv:"Sn Error"`
}
