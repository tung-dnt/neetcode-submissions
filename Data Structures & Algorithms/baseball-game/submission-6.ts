class Solution {
  /**
   * @param {string[]} operations
   * @return {number}
   */
  calPoints(operations: string[]): number {
    const stack = []

    for (let i in operations) {
      const ele = operations[i]
      if (ele === "+") {
        const first = stack.pop();
        const second = stack.pop();
        const sum = +first + +second;

        stack.push(second);
        stack.push(first);
        stack.push(sum);
      } else if (ele === "C") {
        stack.pop();
      } else if (ele === "D") {
        const num = stack.pop();

        stack.push(num)
        stack.push(+num * 2);
      } else {
        stack.push(ele);
      }
    }

    let sum = 0;
    for (let i in stack) {
      sum += +stack[i];
    }

    return sum;
  }
}
