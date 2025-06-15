package main

import "encoding/json"

type DiscountedProduct struct {
	*Product `json:"product,omitempty"`
	Discount float64 `json:"-"`
}

func (dp *DiscountedProduct) MarshalJSON() (jsn []byte, err error) {
	if dp.Product != nil {
		m := map[string]interface{}{
			"product": dp.Name,
			"cost":    dp.Price - dp.Discount,
		}
		jsn, err = json.Marshal(m)
	}
	return
}
