package GrpcSymbol

import orbit "github.com/coda247/orbit"

func (s *Symbol) ToSymbol() orbit.Symbol {
	return orbit.Symbol{BaseCurrency: s.BaseCurrency, QuoteCurrency: s.QuoteCurrency}
}
