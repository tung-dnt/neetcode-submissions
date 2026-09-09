class Solution {
    /**
     * @param {string[]} operations
     * @return {number}
     */
    calPoints(operations) {
        let result = 0
        const record = []
        
        for(let op of operations) {
            if(op == "+"){
              record.push(record[record.length -1] +   (record[record.length- 2]))
            } else if(op =="D"){
              record.push(record[record.length - 1] * 2)
            } else if(op == "C"){
              record.pop()
            } else {
              record.push(parseInt(op))
            }
        }
        console.log(record)
        
        result = record.reduce((sum, e) => sum + e, 0)
      
        return result
    }
}
