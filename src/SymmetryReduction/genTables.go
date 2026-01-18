package SymmetryReduction

import "fmt"

func genTable(name string, mapping [9]int) {
	fmt.Printf("var %s = [512]int32{\n", name)
	for mask := 0; mask < 512; mask++ {
		var out int32 = 0
		for i := 0; i < 9; i++ {
			if (mask>>mapping[i])&1 == 1 {
				out |= 1 << i
			}
		}
		fmt.Printf("%d,", out)
		if (mask+1)%16 == 0 {
			fmt.Println()
		}
	}
	fmt.Println("}")
}

func main() {
	genTable("ROT90TABLE", [9]int{6, 3, 0, 7, 4, 1, 8, 5, 2})
	genTable("ROT180TABLE", [9]int{8, 7, 6, 5, 4, 3, 2, 1, 0})
	genTable("ROT270TABLE", [9]int{2, 5, 8, 1, 4, 7, 0, 3, 6})
	genTable("FLIPHTABLE", [9]int{6, 7, 8, 3, 4, 5, 0, 1, 2})
	genTable("FLIPVTABLE", [9]int{2, 1, 0, 5, 4, 3, 8, 7, 6})
	genTable("DIAGTABLE", [9]int{0, 3, 6, 1, 4, 7, 2, 5, 8})
	genTable("ANTIDIAGTABLE", [9]int{8, 5, 2, 7, 4, 1, 6, 3, 0})
}

// hard-code the generated outputs tables to src/constants/Symmtables.go
