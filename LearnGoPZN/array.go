package main
import "fmt"

func main() {
	var contohArr [5]int
	contohArr = [5]int{10, 20, 30, 40, 50}
	fmt.Println(contohArr)
	fmt.Println()
	for i := 0; i < len(contohArr); i++ {
		fmt.Printf("Index ke-%d: %d\n", i, contohArr[i])
	}
	fmt.Println()
	fmt.Printf("Panjang array: %d, capacity array: %d\n", len(contohArr), cap(contohArr))

	var contohArr2 = [3]string{"apel", "jeruk", "mangga"}
	fmt.Println(contohArr2)
	for i := 0; i < len(contohArr2); i++ {
		for j := 0; j < len(contohArr2[i]); j++ {
			fmt.Printf("Index ke-%d dari elemen ke-%d: %c\n", j, i, contohArr2[i][j])
		}
		fmt.Println()
	}
}
