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