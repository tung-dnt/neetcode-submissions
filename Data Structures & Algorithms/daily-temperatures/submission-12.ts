class Solution {
    /**
     * @param {number[]} temperatures
     * @return {number[]}
     */
    dailyTemperatures(temperatures: number[]): number[] {
        const result: number[] = new Array(temperatures.length).fill(0);
        const stack: number[][] = [[temperatures[0], 0]];

        for (let i = 1; i < temperatures.length; i++) {
            while (stack.length > 0 && temperatures[i] > stack[stack.length - 1][0]) {
                const [_, index] = stack.pop();
                result[index] = i - index;
            }

            stack.push([temperatures[i], i]);
        }

        return result;
    }
}
