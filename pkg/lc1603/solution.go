package lc1603

type ParkingSystem [3]int

func Constructor(big, medium, small int) ParkingSystem {
	return ParkingSystem{big, medium, small}
}

func (s *ParkingSystem) AddCar(carType int) bool {
	value := s[carType-1]
	s[carType-1] = max(s[carType-1]-1, 0)
	return value > 0
}
