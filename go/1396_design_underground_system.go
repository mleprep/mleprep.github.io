//go:build ignore

// 1396. Design Underground System (Medium)
// https://leetcode.com/problems/design-underground-system/
// Run: go run 1396_design_underground_system.go lc.go
package main

import (
	"fmt"
	"os"
)

type checkIn struct {
	station string
	t       int
}

type UndergroundSystem struct {
	active map[int]checkIn      // id -> where and when they entered
	stats  map[[2]string][2]int // (from, to) -> (total time, trips)
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{active: map[int]checkIn{}, stats: map[[2]string][2]int{}}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.active[id] = checkIn{stationName, t}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	in := this.active[id]
	delete(this.active, id)
	key := [2]string{in.station, stationName}
	s := this.stats[key]
	s[0] += t - in.t
	s[1]++
	this.stats[key] = s
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	s := this.stats[[2]string{startStation, endStation}]
	return float64(s[0]) / float64(s[1])
}

func main() {
	u := Constructor()
	u.CheckIn(45, "Leyton", 3)
	u.CheckIn(32, "Paradise", 8)
	u.CheckIn(27, "Leyton", 10)
	u.CheckOut(45, "Waterloo", 15)
	u.CheckOut(27, "Waterloo", 20)
	u.CheckOut(32, "Cambridge", 22)
	assertEq([]float64{u.GetAverageTime("Paradise", "Cambridge"), u.GetAverageTime("Leyton", "Waterloo")}, []float64{14, 11})
	u2 := Constructor()
	u2.CheckIn(10, "A", 1)
	u2.CheckOut(10, "B", 5)
	u2.CheckIn(11, "A", 2)
	u2.CheckOut(11, "B", 4)
	assertEq(u2.GetAverageTime("A", "B"), 3.0)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
