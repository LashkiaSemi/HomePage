import { combineReducers } from 'redux'
import { account } from './account'
import { activities } from './activity'
import { equipments } from './equipment'
import { isLoading } from './isLoading'
import { jobs } from './job'
import { lectures, lecture } from './lecture'
import { logged } from './logged'
import { members, member } from './member'
import { researches } from './research'
import { societies } from './society'
import { tags } from './tag'
import { apiError } from './error'


const rootReducer = combineReducers({
    isLoading,
    activities,
    societies,
    researches,
    members,
    member,
    jobs,
    equipments,
    lectures,
    lecture,
    tags,
    logged,
    account,
    apiError,
})

export default rootReducer