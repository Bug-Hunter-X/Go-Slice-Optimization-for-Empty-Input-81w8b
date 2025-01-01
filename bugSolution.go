func myFuncOptimized(a []int) []int {
  if len(a) == 0 {
    return []int{}
  }
  result := make([]int, len(a))
  for i := range a {
    result[i] = a[i] * 2
  }
  return result
}

func main() {
  nums := []int{1, 2, 3, 4, 5}
  doubledNums := myFuncOptimized(nums)
  fmt.Println(doubledNums) // Output: [2 4 6 8 10]

  emptyNums := []int{}
  doubledEmptyNums := myFuncOptimized(emptyNums)
  fmt.Println(doubledEmptyNums) // Output: []
} 
//The solution avoids unnecessary allocation when the slice is empty by directly returning an empty slice instead of allocating and then returning it.