import * as Type from '../constants/action-types'

export function apiError(state={}, action) {
    switch(action.type) {
        case Type.API_ERROR:
            console.log("api error")
            return { error: action.payload.response }
        
        // failure
        case Type.LOGIN_FAILURE:
            return { error: action.payload.response }

        case Type.FETCH_ACTIVITIES_FAILURE:
        case Type.CREATE_ACTIVITY_FAILURE:
        case Type.UPDATE_ACTIVITY_FAILURE:
        case Type.DELETE_ACTIVITY_FAILURE:
            return { error: action.payload.response }

        case Type.FETCH_SOCIETIES_FAILURE:
        case Type.CREATE_SOCIETY_FAILURE:
        case Type.UPDATE_SOCIETY_FAILURE:
        case Type.DELETE_SOCIETY_FAILURE:
            return { error: action.payload.response }

        case Type.FETCH_RESEARCHES_FAILURE:
        case Type.CREATE_RESEARCH_FAILURE:
        case Type.UPDATE_RESEARCH_FAILURE:
        case Type.DELETE_RESEARCH_FAILURE:
            return { error: action.payload.response }

        case Type.FETCH_JOBS_FAILURE:
        case Type.CREATE_JOB_FAILURE:
        case Type.UPDATE_JOB_FAILURE:
        case Type.DELETE_JOB_FAILURE:
            return { error: action.payload.response }

        case Type.FETCH_MEMBERS_FAILURE:
        case Type.CREATE_MEMBER_FAILURE:
        case Type.UPDATE_MEMBER_FAILURE:
        case Type.DELETE_MEMBER_FAILURE:
            return { error: action.payload.response }

        case Type.FETCH_EQUIPMENTS_FAILURE:
        case Type.CREATE_EQUIPMENT_FAILURE:
        case Type.UPDATE_EQUIPMENT_FAILURE:
        case Type.DELETE_EQUIPMENT_FAILURE:
            return { error: action.payload.response }

        case Type.FETCH_LECTURES_FAILURE:
        case Type.CREATE_LECTURE_FAILURE:
        case Type.UPDATE_LECTURE_FAILURE:
        case Type.DELETE_LECTURE_FAILURE:
            return { error: action.payload.response }

        case Type.FETCH_ACCOUNT_FAILURE:
        case Type.UPDATE_ACCOUNT_FAILURE:
        case Type.UPDATE_ACCOUNT_PASSWORD_FAILURE:
            return { error: action.payload.response }
        
        // success
        case Type.LOGIN_SUCCESS:
            return {}

        case Type.FETCH_ACTIVITIES_SUCCESS:
        case Type.CREATE_ACTIVITY_SUCCESS:
        case Type.UPDATE_ACTIVITY_SUCCESS:
        case Type.DELETE_ACTIVITY_SUCCESS:
            return {}

        case Type.FETCH_SOCIETIES_SUCCESS:
        case Type.CREATE_SOCIETY_SUCCESS:
        case Type.UPDATE_SOCIETY_SUCCESS:
        case Type.DELETE_SOCIETY_SUCCESS:
            return {}

        case Type.FETCH_RESEARCHES_SUCCESS:
        case Type.CREATE_RESEARCH_SUCCESS:
        case Type.UPDATE_RESEARCH_SUCCESS:
        case Type.DELETE_RESEARCH_SUCCESS:
            return {}

        case Type.FETCH_JOBS_SUCCESS:
        case Type.CREATE_JOB_SUCCESS:
        case Type.UPDATE_JOB_SUCCESS:
        case Type.DELETE_JOB_SUCCESS:
            return {}

        case Type.FETCH_MEMBERS_SUCCESS:
        case Type.CREATE_MEMBER_SUCCESS:
        case Type.UPDATE_MEMBER_SUCCESS:
        case Type.DELETE_MEMBER_SUCCESS:
            return {}

        case Type.FETCH_EQUIPMENTS_SUCCESS:
        case Type.CREATE_EQUIPMENT_SUCCESS:
        case Type.UPDATE_EQUIPMENT_SUCCESS:
        case Type.DELETE_EQUIPMENT_SUCCESS:
            return {}

        case Type.FETCH_LECTURES_SUCCESS:
        case Type.CREATE_LECTURE_SUCCESS:
        case Type.UPDATE_LECTURE_SUCCESS:
        case Type.DELETE_LECTURE_SUCCESS:
            return {}

        case Type.FETCH_ACCOUNT_SUCCESS:
        case Type.UPDATE_ACCOUNT_SUCCESS:
        case Type.UPDATE_ACCOUNT_PASSWORD_SUCCESS:
            return {}

        default:
            return state
    }
}