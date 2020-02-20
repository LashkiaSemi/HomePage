/*
checkEmptyString 文字列が空値でないことをチェックする
items: [{
    id: 何かしら値。。。
    value: 入力値
    field: フィールドの名前。エラーバーに表示されるのでできれば解読できるやつ
}]
*/
// もしかして使ってないwww
// export const checkEmptyString = (items = []) => {
//     const errors = []
//     items.forEach(item => {
//         if(item.value === "") {
//             errors.push({id: item.id+"Empty", content: item.field+"は必須です"})
//         }
//     })
//     return errors
// }

/*
checkIsEmpty 値が空値じゃないかチェックする
value: 入力値
*/
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

/*
parseStringToBool 文字列の'true'の時のみboolでtrueを返す
value: string
*/
export function parseStringToBool(value) {
    if (value === "true") {
        return true
    }
    return false
}