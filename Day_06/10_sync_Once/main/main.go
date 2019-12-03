package main

type T1 struct {
	left  string
	up    string
	right string
	down  string
}

var T2 map[string]T1

func loadT2() {
	T2["string"] = T1{
		left:  "left.png",
		up:    "up.png",
		right: "right.png",
		down:  "down.png",
	}

	// T2 = map[string]T1{
	// 	left:  "left.png",
	// 	up:    "up.png",
	// 	right: "right.png",
	// 	down:  "down.png",
	// }
}

func main() {

}
