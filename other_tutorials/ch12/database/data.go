package database

type Point struct {
	X int
	Y int
}

func Change(dat *[]Point) {
	var pts = []Point{Point{X: 88, Y: 99}, Point{X: 333, Y: 444}}
	*dat = pts
}
