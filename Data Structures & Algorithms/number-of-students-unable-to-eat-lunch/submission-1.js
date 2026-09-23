class Solution {
    /**
     * @param {number[]} students
     * @param {number[]} sandwiches
     * @return {number}
     */

    countStudents(students, sandwiches) {
        const count = [0, 0]
        let remaining = students.length
        for(const student of students) {
            count[student]++
        }

        for(const s of sandwiches){
            if(count[s] > 0){
                count[s]--
                remaining--
            }else{
                break
            }
        }
    return remaining

    }
}
