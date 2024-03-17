package main

func SumAllTails(tailsToSum... []int) []int {
  var sums []int 
  for _, numbers := range tailsToSum {
    if len(numbers) == 0 {
      sums = append(sums, 0)
    } else {
      // sums = append(sums, numbers[len(numbers) - 1])
      // Slices can be sliced! The syntax is slice[low:high]
      sums = append(sums, Sum(numbers[1:]))
    }
  }
  return sums
}
