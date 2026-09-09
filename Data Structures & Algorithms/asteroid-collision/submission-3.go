/*
# Insight
For each element, compare nearest contrast-sighted number such that absolute value of it must be bigger, otherwise it's eliminated
-> Stack problem, push to stack, evaluate top items

- Negative number move to left, positive move to right
-> To make collisions happen:
  - left items(stack) must move to right  
  - right items(array) must move to left
# Invariant
Check for collisions and cleanup elements:
```
for top > 0 && asteroids[i] < 0:
- abs(top) > abs(asteroids[i]): dont push to stack
- abs(top) < abs(asteroids[i]): pop stack + push new ast
- abs(top) = abs(asteroids[i]): pop stack + ignore ast
```
*/
func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for _, ast := range asteroids {
		// When collision happens
		for len(stack) > 0 && ast < 0 && stack[len(stack)-1] > 0 {
			diff := stack[len(stack)-1] + ast
			if diff < 0 {
				stack = stack[:len(stack)-1]
			} else if diff > 0 {
				ast = 0
			} else if diff == 0 {
				stack = stack[:len(stack)-1]
				ast = 0
			}
		}
		if ast != 0 {
			stack = append(stack, ast)
		}
	}
	return stack
}