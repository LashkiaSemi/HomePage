// idってプロパティから一つのitemを探す
export const findItemByID = (items, id) => {
    return items.find(item => item.id == id)
}

// admin siteのリストで、見出しつけるために使ってます。
// typeにpathというか、データの種類(adminListのpath)を送って、dataにitemを送る、みたいな
export function findCaptionByDataType(type, data) {
    switch(type) {
        case "activities":
            return data.activity
        case "societies":
            return data.title
        case "researches":
            return data.title
        case "members":
            return data.name
        case "jobs":
            return data.company
        case "equipments":
            return data.name
        case "lectures":
            return data.title
        case "tags":
            return data.name
        default:
            return data.id
    }
}
