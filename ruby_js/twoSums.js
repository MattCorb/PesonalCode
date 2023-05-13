const twoSums = (nums, target) =>{
    const numsObj = {}
    for (let i = 0; i < nums.length; i++){
        let compliment = target - nums[i]

        if (numsObj[compliment] != null){
            return [numsObj[compliment], i]
        }
        else{
            numsObj[nums[i]] = i
        }
    }
}

console.log(twoSums([2,7,11,15], 9))

//Ruby 

// def two_sum(nums, target)
//     numsObj = Hash.new
//     for i in 0..nums.length
//       compliment = target - nums[i]
//       if numsObj.fetch(compliment, nil) != nil
//         return [numsObj[compliment], i ]
//       else
//         numsObj[nums[i]] = i
//       end
//     end
// end