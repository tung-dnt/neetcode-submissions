class Solution {
    /**
     * @param {string} s
     * @return {string}
     */
    decodeString(s: string): string {
        const stack: string[] = [];

        for (let i of s) {
            if (i === "]") {
                let temp: string = "";
                let sub: string = "";
                let k: string = "";
                let flag: boolean = false;

                while (true) {
                    const item = stack.pop();

                    if (!item) break;

                    if (item === "[") {
                        sub = temp;
                        temp = "";
                        flag = true;
                    } else {
                        temp = `${item}${temp}`;

                        if (flag && Number.isNaN(Number(stack[stack.length - 1]))) {
                            k = temp;

                            break;
                        }
                    }
                }

                stack.push(sub.repeat(parseInt(k)));
            } else stack.push(i);
        }

        return stack.join("");
    }
}
