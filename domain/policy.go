package domain

type BeginnerPolicy struct {
	minEasiness int
}

func NewBeginnerPolicy(minEasiness int) BeginnerPolicy {
	return BeginnerPolicy{minEasiness: minEasiness}
}

func (bp BeginnerPolicy) Allows(r Rentable) bool {
	return r.Easiness() >= bp.minEasiness
}
