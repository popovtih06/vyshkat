package domain

import "testing"

func TestBeginnerPolicyAllows(t *testing.T) {
	policy := NewBeginnerPolicy(6)

	cases := []struct {
		name     string
		easiness int
		want     bool
	}{
		{"ниже порога", 5, false},
		{"ровно порог", 6, true},
		{"выше порога", 9, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			scooter := NewScooter("Самокат", "VK-0001", c.easiness)

			got := policy.Allows(scooter)

			if got != c.want {
				t.Errorf("простота %d: получили %v, ожидали %v", c.easiness, got, c.want)
			}
		})
	}
}

func TestBeginnerPolicyRespectsThreshold(t *testing.T) {
	strict := NewBeginnerPolicy(8)
	scooter := NewScooter("Самокат", "VK-0001", 7)

	if strict.Allows(scooter) {
		t.Error("при пороге 8 самокат с простотой 7 не должен подходить новичкам")
	}
}
