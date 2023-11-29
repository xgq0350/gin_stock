package stock

type Tax struct {
}

func (e *Tax) CalcTax(types string, vol float64) float64 {
	//佣金
	if vol > 38500 {
		vol = vol - vol*0.00013
	} else {
		vol = vol - 5
	}
	//印花税
	if types == "1" {
	} else {
		vol = vol - vol*0.001
	}
	//过户费
	vol = vol - (vol * 0.00001)
	return vol
}
