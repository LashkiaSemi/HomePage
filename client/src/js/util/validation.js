// checkEmptyString items = [{id, value, field}]
export const checkEmptyString = (items = []) => {
    const errors = []
    items.forEach(item => {
        if(item.value === "") {
            errors.push({id: item.id+"Empty", content: item.field+"は必須です"})
        }
    })
    // items.forEach(item => console.log(item))
    return errors
}

// 空値であるか判断する。trueなら空
export function checkIsEmpty(value) {
    switch(typeof value) {
        case 'string':
            return value === ""
        case "number":
            return value === 0
        case "undefined":
            return true
        case "object":
            return Object.keys(value).length > 0
        default:
            console.log("checkIsEmpty: ", typeof value, value)
            return true
    }
}

// "true"のときのみtrueを変換する
export function parseStringToBool(value) {
    if (value === "true") {
        return true
    }
    return false
}