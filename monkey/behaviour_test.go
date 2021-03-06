package monkey

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItGetsRandomBehavioursAccordingToFrequency(t *testing.T) {
	behaviour1 := behaviour{
		Frequency: 0.15,
		Status:    123,
	}

	behaviour2 := behaviour{
		Frequency: 0.2,
		Status:    456,
	}

	allBehaviours := []behaviour{behaviour1, behaviour2}

	result := getBehaviour(allBehaviours, fakeRandomiser{0.9})

	assert.Nil(t, result)

	assert.Equal(t, getBehaviour(allBehaviours, fakeRandomiser{0.13}).Status, behaviour1.Status)

	assert.Equal(t, getBehaviour(allBehaviours, fakeRandomiser{0.19}).Status, behaviour2.Status)
}

func TestItReturnsAnErrorForBadYAML(t *testing.T) {
	_, err := monkeyConfigFromYAML([]byte("lol not YAML"))
	assert.NotNil(t, err, "Should get an error for bad YAML")
}

func TestItParsesYAMLIntoBehaviour(t *testing.T) {
	yaml := `
---
# Writes a different body 50% of the time
- body: "This is wrong :( "
  frequency: 0.5

# Delays initial writing of response by a second 20% of the time
- delay: 1000
  frequency: 0.2

# Returns a 404 30% of the time
- status: 404
  frequency: 0.3

# Write 10,000,000 garbage bytes 10% of the time
- garbage: 10000000
  frequency: 0.09
  `
	behaviours, _ := monkeyConfigFromYAML([]byte(yaml))
	assert.Len(t, behaviours, 4)
}

type fakeRandomiser struct {
	value float64
}

func (f fakeRandomiser) getFloat() float64 {
	return f.value
}
