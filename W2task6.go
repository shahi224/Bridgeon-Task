package main
import (
	"fmt"
	"sort"
)

func mergeSortedarr(a,b,c []int)[]int {
	merged := append(a,b...)
	merged = append(merged, c...)
	sort.Ints(merged)
	return merged
}

func main(){

	arr1 :=[]int {1,3,2,5,4}
	arr2 :=[]int {9,6,8,7}
	arr3 :=[]int {10,13,12,11}

	fmt.Println(mergeSortedarr(arr1,arr2,arr3))

}