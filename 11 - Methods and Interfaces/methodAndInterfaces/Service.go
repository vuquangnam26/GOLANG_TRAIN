package main

type Service struct {
	description    string
	durationMonths int
	monthlyFee     float64
	feature        []string
}

func (c Service) getName() string {
	return c.description
}

func (c Service) getCost(recur bool) float64 {
	if recur {
		return c.monthlyFee * float64(c.durationMonths)

	}
	return c.monthlyFee
}
