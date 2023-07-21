package mygopackage

import "fmt"

type Cotacao struct {
  Moeda string
}

// Converte r em reais para a moeda configurada na struct
func (c *Cotacao) Converter(r float32)  (float32, error) {
  if c.Moeda == "dolar" {
    return c.converterRealParaDolar(r), nil
  } else if c.Moeda == "euro" {
    return c.converterRealParaEuro(r), nil
  }
  return 0, fmt.Errorf("%s: %s", "moeda desconhecida", c.Moeda)
}

func (Cotacao) converterRealParaDolar(r float32) float32 {
  valorDolarEmReal := float32(5) // valor ficticio para comprar 1 dolar com reais
  return r * valorDolarEmReal
}

func (Cotacao) converterRealParaEuro(r float32) float32 {
  valorEuroEmReal := float32(6) // valor ficticio para comprar 1 euro com reais
  return r * valorEuroEmReal
}
