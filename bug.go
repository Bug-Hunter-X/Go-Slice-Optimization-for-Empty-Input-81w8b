func myFunc(a []int) []int {
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
  doubledNums := myFunc(nums)
  fmt.Println(doubledNums) // Output: [2 4 6 8 10]
  
  emptyNums := []int{}
  doubledEmptyNums := myFunc(emptyNums)
  fmt.Println(doubledEmptyNums) // Output: []
}