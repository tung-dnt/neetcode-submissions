class Solution {
    /**
     * @param {string} path
     * @return {string}
     */
    simplifyPath(path: string): string {
        const pathArr: string[] = path.split("/");
        const stack: string[] = [];

        for (let i of pathArr) {
            if (i === "..") {
                if (stack.length > 0) stack.pop();
            } else if (i !== "." && i !== "") stack.push(i);
        }

        return "/" + stack.join("/");
    }
}
