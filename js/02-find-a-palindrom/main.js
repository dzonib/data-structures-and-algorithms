function palindrom(s) {
    const reversed = s.split('').reverse().join("")

    return s === reversed
//    if (s === reversed) {
//        return true
//    } else {
//        return false
//    }
}

//function palindrome(s) {
//    const splited = s.split('')
//
//    const reversed = splited.reverse().join("")
//
//
//    return splited.every((letter, i) => {
//
//        return letter === reversed[i]
//    })
//}


//function palindrome(s) {
//return str.split('').every((char, i) => {
//    return char === str[str.length - i - 1];
//})
//    })
//}