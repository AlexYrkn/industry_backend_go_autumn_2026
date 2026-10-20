package main

type Stats struct {
	Count         int
	Sum, Min, Max int64
}

func Calc(nums []int64) Stats {
	if len(nums) < 2 {
		return Stats{}
	}

	first := nums[1] - nums[0]
	res := Stats{
		Count: 1,
		Sum:   first,
		Min:   first,
		Max:   first,
	}
	for i := 2; i < len(nums); i++ {
		dif := nums[i] - nums[i-1]
		if dif < res.Min {
			res.Min = dif
		}
		if dif > res.Max {
			res.Max = dif
		}
		res.Sum += dif
		res.Count++
	}

	return res
}
