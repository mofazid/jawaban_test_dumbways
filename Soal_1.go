package main 

import "fmt"

func main(){
	// HitungKembalian(1000, 100000) //
	LoopingDumbwaysID()
}

func HitungKembalian(totalBelanja int, jumlahUang int) {
	var kembalian int

	stockKembalian := []int{50000, 20000, 10000, 5000, 2000, 1000, 500}
	totalKembalian := jumlahUang - totalBelanja

	m := make(map[int][]int)

	for _, v := range stockKembalian {

		if v <= totalKembalian {

			if (v + kembalian) > totalKembalian {
				continue
			}

			kembalian += v
			m[v] = append(m[v], v)

			if (v + kembalian) <= totalKembalian-kembalian {
				for {
					kembalian += v
					m[v] = append(m[v], v)
					if (totalKembalian - kembalian) < v {
						break
					}
				}
			}
		}

		if kembalian == totalKembalian {
			break
		}
	}

	for _, v := range stockKembalian {
		res := len(m[v])
		if res != 0 {
			fmt.Printf("%d x %d \n", res, v)
		}
	}
}