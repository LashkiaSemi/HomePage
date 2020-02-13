export const findActivityByID = (activities, id) => {
    return activities.find(activity => activity.id == id) 
}