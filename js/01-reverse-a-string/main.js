//function reverse(str) {
//
//    const strArr = str.split("")
//    const strLength = strArr.length
//
//
//    const result = []
//    for (let i = 0; i < strLength; i++) {
//
//        const index = strLength - 1 - i
//
//        result.push(strArr[index])
//    }
//    return result.join("")
//}


//function reverse(str) {
//    let reversed = ''
//
//    for (let char of str) {
//        reversed = char + reversed;
//    }
//
//    return reversed
//}


function reverse(str) {
    return str.split("").reduce((acc, c) => {
        return c + acc
    }, '')
}