import * as Type from '../constants/action-types'

export const loginRequest = (payload) => { return { type: Type.LOGIN_REQUEST, payload} }
export const loginSuccess = (payload) => { return { type: Type.LOGIN_SUCCESS, payload} }
export const loginFailure = (payload) => { return { type: Type.LOGIN_FAILURE, payload} }

export const logoutRequest = () => { return {type: Type.LOGOUT_REQUEST} }
export const logoutSuccess = () => { return {type: Type.LOGOUT_SUCCESS} }
export const logoutFailure = (payload) => { return {type: Type.LOGOUT_FAILURE, payload} }

export const fetchAccountRequest = () => { return {type: Type.FETCH_ACCOUNT_REQUEST} }
export const fetchAccountSuccess = (payload) => { return {type: Type.FETCH_ACCOUNT_SUCCESS, payload} }
export const fetchAccountFailure = (payload) => { return {type: Type.FETCH_ACCOUNT_FAILURE, payload} }

export const updateAccountRequest = (payload) => { return { type: Type.UPDATE_ACCOUNT_REQUEST, payload } }
export const updateAccountSuccess = (payload) => { return { type: Type.UPDATE_ACCOUNT_SUCCESS, payload } }
export const updateAccountFailure = (payload) => { return { type: Type.UPDATE_ACCOUNT_FAILURE, payload } }
export const updateAccountPasswordRequest = (payload) => { return { type: Type.UPDATE_ACCOUNT_PASSWORD_REQUEST, payload } }
export const updateAccountPasswordSuccess = (payload) => { return { type: Type.UPDATE_ACCOUNT_PASSWORD_SUCCESS, payload } }
export const updateAccountPasswordFailure = (payload) => { return { type: Type.UPDATE_ACCOUNT_PASSWORD_FAILURE, payload } }

export const updateMemberRequest = (payload) => { return { type: Type.UPDATE_MEMBER_REQUEST, payload} }
export const updateMemberSuccess = (payload) => { return { type: Type.UPDATE_MEMBER_SUCCESS, payload} }
export const updateMemberFailure = (payload) => { return { type: Type.UPDATE_MEMBER_FAILURE, payload} }

export const fetchMembersRequest = () => { return { type: Type.FETCH_MEMBERS_REQUEST } }
export const fetchMembersSuccess = (payload) => { return { type: Type.FETCH_MEMBERS_SUCCESS, payload } }
export const fetchMembersFailure = (payload) => { return { type: Type.FETCH_MEMBERS_FAILURE, payload } }

export const fetchMemberRequest = (payload) => { return { type: Type.FETCH_MEMBER_REQUEST, payload } }
export const fetchMemberSuccess = (payload) => { return { type: Type.FETCH_MEMBER_SUCCESS, payload } }
export const fetchMemberFailure = (payload) => { return { type: Type.FETCH_MEMBER_FAILURE, payload } }

export const fetchActivitiesRequest = () => { return { type: Type.FETCH_ACTIVITIES_REQUEST } }
export const fetchActivitiesSuccess = (payload) => { return { type: Type.FETCH_ACTIVITIES_SUCCESS, payload } }
export const fetchActivitiesFailure = (payload) => { return { type: Type.FETCH_ACTIVITIES_FAILURE, payload } }

export const fetchSocietiesRequest = () => { return { type: Type.FETCH_SOCIETIES_REQUEST } }
export const fetchSocietiesSuccess = (payload) => { return { type: Type.FETCH_SOCIETIES_SUCCESS, payload } }
export const fetchSocietiesFailure = (payload) => { return { type: Type.FETCH_SOCIETIES_FAILURE, payload } }

export const fetchResearchesRequest = () => { return { type: Type.FETCH_RESEARCHES_REQUEST } }
export const fetchResearchesSuccess = (payload) => { return { type: Type.FETCH_RESEARCHES_SUCCESS, payload } }
export const fetchResearchesFailure = (payload) => { return { type: Type.FETCH_RESEARCHES_FAILURE, payload } }

export const fetchEquipmentsRequest = () => { return { type: Type.FETCH_EQUIPMENTS_REQUEST } }
export const fetchEquipmentsSuccess = (payload) => { return { type: Type.FETCH_EQUIPMENTS_SUCCESS, payload } }
export const fetchEquipmentsFailure = (payload) => { return { type: Type.FETCH_EQUIPMENTS_FAILURE, payload } }

export const fetchLecturesRequest = () => { return { type: Type.FETCH_LECTURES_REQUEST } }
export const fetchLecturesSuccess = (payload) => { return { type: Type.FETCH_LECTURES_SUCCESS, payload } }
export const fetchLecturesFailure = (payload) => { return { type: Type.FETCH_LECTURES_FAILURE, payload } }
export const fetchLectureRequest = (payload) => { return { type: Type.FETCH_LECTURE_REQUEST, payload } }
export const fetchLectureSuccess = (payload) => { return { type: Type.FETCH_LECTURE_SUCCESS, payload } }
export const fetchLectureFailure = (payload) => { return { type: Type.FETCH_LECTURE_FAILURE, payload } }
export const createLectureRequest = (payload) => { return { type: Type.CREATE_LECTURE_REQUEST, payload } }
export const createLectureSuccess = (payload) => { return { type: Type.CREATE_LECTURE_SUCCESS, payload } }
export const createLectureFailure = (payload) => { return { type: Type.CREATE_LECTURE_FAILURE, payload } }

export const fetchJobsRequest = () => { return { type: Type.FETCH_JOBS_REQUEST } }
export const fetchJobsSuccess = (payload) => { return { type: Type.FETCH_JOBS_SUCCESS, payload } }
export const fetchJobsFailure = (payload) => { return { type: Type.FETCH_JOBS_FAILURE, payload } }

export const showLoading = () => { return { type: Type.SHOW_LOADING} }
export const hideLoading = () => { return { type: Type.HIDE_LOADING} }
export const apiError = (payload) => { return { type: Type.API_ERROR, payload} }