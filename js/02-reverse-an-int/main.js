//function reverseInt(n) {
//
//    let result;
//
//    const nmb = Math.sign(n)
//
//    switch (nmb) {
//        case 1:
//            result = n.toString().split("").reverse().join("")
//            break;
//        case -1:
//            const arr = n.toString().split("").reverse();
//
//            const lastElement = arr.pop();
//            arr.unshift("-");
//
//            result = arr.join("");
//            console.log("RESULT", result)
//            break;
//        default:
//            result = "0"
//    }
//
//    return parseInt(result)
//}


// optimal solution
function reverseInt(n) {
    const reversed = n.toString().split("").reverse().join('')

    return parseInt(reversed) * Math.sign(n)
}