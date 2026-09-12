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
                    const last: number = stable[stable.length - 1];

                    if (stable.length === 0 || last < 0) {
                        stable.push(a);
                        break;
                    } else {
                        const abs = Math.abs(a);
                        if (last === abs) {
                            stable.pop();
                            break;
                        } else if (last < abs) {
                            stable.pop();
                        } else break;
                    }
                }
            }
        }

        return stable;
    }
}
