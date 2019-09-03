package five

import "fmt"

type T1 interface {
	Get_Name() string
}

type Struct_One struct {
	Name string
}

func (this Struct_One) Get_Name() string {
	return this.Name
}

func Five() {
	var one Struct_One = Struct_One{
		Name: "Smurfs",
	}
	var two T1
	two = one

	res, ok := two.(T1)
	fmt.Println(res, ok)

	// 这里讲的不详细
	// if sv, ok := one.(T1); ok {
	// 	fmt.Println(ok)
	// 	fmt.Printf("v implements T1():%s\n", sv.T1())
	// }
}
