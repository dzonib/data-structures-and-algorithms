function maxChar(str) {
    let result;

    let maxValue = -Infinity

    const holder = {}

    // NOTE: WE CAN ITERATE OVER A STRING, but I prefer this way (more expensive though)
    // in other languages you cannot do that

    const strArr = str.split("")

    // make an object of keys (char) and values (number)
    for (let c of strArr) {
        if (!holder[c]) {
            holder[c] = 1
        } else {
            holder[c] = holder[c] + 1
        }
    }

    for (const key in holder) {
        if (holder[key] > maxValue) {
            maxValue = holder[key]
            result = key
        }
    }

    return result
}