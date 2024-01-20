package lc0295

import (
	"testing"
)

func TestMedianFinder_binarySearch(t *testing.T) {
	t.SkipNow() // TODO: enable this test
	type fields struct {
		arr []int
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: ``,
			fields: fields{
				arr: func() []int {
					arr := make([]int, 10, 10)
					for i := 0; i < 10; i++ {
						arr[i] = i * 2
					}
					return arr
				}(),
			},
			args: args{
				value: -1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mf := &MedianFinder{
				arr: tt.fields.arr,
			}
			if got := mf.binarySearch(tt.args.value); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (mf *MedianFinder) findMedianAssert(t *testing.T, want float64) {
	got := mf.FindMedian()
	if got != want {
		t.Errorf("FindMedian() = %v, want %v", got, want)
	}
}

func TestMedianFinder_FindMedian(t *testing.T) {
	t.SkipNow() // TODO: enable this test
	tests := []struct {
		name string
	}{
		{
			name: `all`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mf := Constructor()
			mf.AddNum(1)
			mf.AddNum(2)

			mf.findMedianAssert(t, 1.5)

			mf.AddNum(3)
			mf.AddNum(4)
			mf.AddNum(5)
			mf.AddNum(6)
			mf.AddNum(7)

			mf.findMedianAssert(t, 4)
		})

		t.Run(tt.name, func(t *testing.T) {
			mf := Constructor()
			mf.AddNum(-1)
			mf.findMedianAssert(t, -1)
			mf.AddNum(-2)
			mf.findMedianAssert(t, -1.5)
			mf.AddNum(-3)
			mf.findMedianAssert(t, -2)
			mf.AddNum(-4)
			mf.findMedianAssert(t, -2.5)
			mf.AddNum(-5)
			mf.findMedianAssert(t, -3)
		})
	}
}
