package apprule

// rule expr
type  Expr interface {
	Match(meta []string,row []string)bool
}
