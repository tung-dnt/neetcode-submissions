class Solution {
    /**
     * @param {number[]} temperatures
     * @return {number[]}
     */
    dailyTemperatures(temperatures: number[]): number[] {
        const result: number[] = new Array(temperatures.length).fill(0);
        const stack: number[] = [temperatures[0]];
        const idxStack: number[] = [0];

        for (let i = 1; i < temperatures.length; i++) {
            while (temperatures[i] > stack[stack.length - 1]) {
                stack.pop();
                const index = idxStack.pop();
                result[index] = i - index;
            }

            stack.push(temperatures[i]);
            idxStack.push(i);
        }

        return result;
    }
}
