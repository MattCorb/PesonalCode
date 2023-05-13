
const quicksort = (numbers) => {

    if (numbers.length <= 1){
        return numbers
    }

    let pivot = numbers[0]
    let left = []
    let right = []

    for (let i = 1; i < numbers.length; i++){
        if (numbers[i] < pivot){
            left.push(numbers[i])
        }
        else{
            right.push(numbers[i])
        }
    }

    return [...quicksort(left), pivot, ...quicksort(right)]
}

console.log(quicksort([3, 7, 2, 5, 1, 4, 6, 8]))

//ruby code
//more familiar with basic syntax espicially that surrounding arrays and their methods

// def quicksort (numbers)
//     if numbers.length <= 1
//         return numbers
//     end
    
//     pivot = numbers[0]
//     left = Array.new
//     right = Array.new
    
//     for i in 1..numbers.length() - 1
//         if numbers[i] < pivot
//             left.push(numbers[i])
//         else
//             right.push(numbers[i])
//         end
//     end
//     return [*quicksort(left), pivot, *quicksort(right)]
// end

// puts quicksort([3, 7, 2, 5, 1, 4, 6, 8])