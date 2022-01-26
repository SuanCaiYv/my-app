const parseBoolean = function (str: string): boolean {
    if (str === "false") {
        return false
    } else if (str === "true") {
        return true
    } else if (str === "") {
        return false
    } else {
        throw new Error("fuck input")
    }
}

export {parseBoolean}