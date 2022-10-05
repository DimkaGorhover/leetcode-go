package p1146

type SnapshotArray struct {
	arr      []*value
	snapshot int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		arr:      make([]*value, length, length),
		snapshot: 0,
	}
}

func (a *SnapshotArray) Set(index int, val int) {
	v := a.arr[index]
	if v == nil {
		v = newValue()
		a.arr[index] = v
	}
	v.set(a.snapshot, val)
}

func (a *SnapshotArray) Snap() int {
	snapshot := a.snapshot
	a.snapshot++
	return snapshot
}

func (a *SnapshotArray) Get(index int, snap_id int) int {
	v := a.arr[index]
	if v == nil {
		return 0
	}
	return v.get(snap_id)
}

type value struct {
	arr            []int
	latestSnapshot int
}

func newValue() *value {
	return &value{
		arr:            make([]int, 0),
		latestSnapshot: 0,
	}
}

func (v *value) get(snapshot int) int {
	if len(v.arr) == 0 {
		return 0
	}
	if snapshot >= v.latestSnapshot {
		return v.arr[v.latestSnapshot] - 1
	}
	for i := snapshot; i >= 0; i-- {
		if val := v.arr[i]; val > 0 {
			return val - 1
		}
	}
	return 0
}

func (v *value) set(snapshot, value int) {
	if snapshot > v.latestSnapshot {
		v.latestSnapshot = snapshot
	}
	v.grow(snapshot + 1)
	v.arr[snapshot] = value + 1
}

func (v *value) grow(minCapacity int) {
	oldCapacity := len(v.arr)
	if oldCapacity > minCapacity {
		return
	}
	newCapacity := oldCapacity + (oldCapacity >> 1)
	if newCapacity < minCapacity {
		newCapacity = minCapacity + 8
	}
	newArr := make([]int, newCapacity)
	copy(newArr, v.arr)
	v.arr = newArr
}
