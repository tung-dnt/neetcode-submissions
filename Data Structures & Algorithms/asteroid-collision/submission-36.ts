class Solution {
    /**
     * @param {number[]} asteroids
     * @return {number[]}
     */
    asteroidCollision(asteroids: number[]): number[] {
        const stable: number[] = [];

        for (let a of asteroids) {
            if (a > 0) stable.push(a);
            else {
                while (true) {
                    if (stable.length === 0 || stable[stable.length - 1] < 0) {
                        stable.push(a);
                        break;
                    } else {
                        if (stable[stable.length - 1] === a * -1) {
                            stable.pop();
                            break;
                        } else if (stable[stable.length - 1] < a * -1) {
                            stable.pop();
                        } else break;
                    }
                }
            }
        }

        return stable;
    }
}
