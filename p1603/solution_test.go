package p1603

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParkingSystem_AddCar(t *testing.T) {
	t.Parallel()

	t.Run(`test_001`, func(t *testing.T) {
		parking := Constructor(2, 2, 2)
		assert.True(t, parking.AddCar(1))
		assert.True(t, parking.AddCar(1))
		assert.False(t, parking.AddCar(1))
	})

	t.Run(`test_002`, func(t *testing.T) {
		parking := Constructor(2, 2, 2)
		assert.True(t, parking.AddCar(2))
		assert.True(t, parking.AddCar(2))
		assert.False(t, parking.AddCar(2))
	})

	t.Run(`test_003`, func(t *testing.T) {
		parking := Constructor(2, 2, 2)
		assert.True(t, parking.AddCar(3))
		assert.True(t, parking.AddCar(3))
		assert.False(t, parking.AddCar(3))
	})
}
