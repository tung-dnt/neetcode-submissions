class Solution {
    /**
     * @param {number} target
     * @param {number[]} position
     * @param {number[]} speed
     * @return {number}
     */
    carFleet(target: number, position: number[], speed: number[]): number {
        const pair: number[][] = position.map((p, i) => [p, speed[i]]);
        pair.sort((a, b) => b[0] - a[0]);

        const timeStack: number[] = [];
        pair.forEach(([p, s]) => {
            const time = (target - p) / s;

            if (timeStack.length === 0 || time > timeStack[timeStack.length - 1])
                timeStack.push(time);
        });

        return timeStack.length;
    }
}
