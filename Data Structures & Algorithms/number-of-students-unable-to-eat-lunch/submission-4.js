class Solution {
    /**
     * @param {number[]} students
     * @param {number[]} sandwiches
     * @return {number}
     */

    countStudents(students, sandwiches) {
        let n = students.length;
        let queue = new Queue();
        for (let student of students) {
            queue.push(student);
        }

        let res = n;
        for (let sand of sandwiches) {
            let count = 0;
            while (count < n && queue.front() !== sand) {
                queue.push(queue.pop());
                count++;
            }

            if (queue.front() === sand) {
                queue.pop();
                res--;
            } else {
                break;
            }
        }
        return res;
    }
}
