import * as Type from '../constants/action-types'

// login
export const loginRequest = (payload) => { return { type: Type.LOGIN_REQUEST, payload} }
export const loginSuccess = (payload) => { return { type: Type.LOGIN_SUCCESS, payload} }
export const loginFailure = (payload) => { return { type: Type.LOGIN_FAILURE, payload} }

// logout
export const logoutRequest = () => { return {type: Type.LOGOUT_REQUEST} }
export const logoutSuccess = () => { return {type: Type.LOGOUT_SUCCESS} }
export const logoutFailure = (payload) => { return {type: Type.LOGOUT_FAILURE, payload} }

// account
export const fetchAccountRequest = () => { return {type: Type.FETCH_ACCOUNT_REQUEST} }
export const fetchAccountSuccess = (payload) => { return {type: Type.FETCH_ACCOUNT_SUCCESS, payload} }
export const fetchAccountFailure = (payload) => { return {type: Type.FETCH_ACCOUNT_FAILURE, payload} }
export const updateAccountRequest = (payload) => { return { type: Type.UPDATE_ACCOUNT_REQUEST, payload } }
export const updateAccountSuccess = (payload) => { return { type: Type.UPDATE_ACCOUNT_SUCCESS, payload } }
export const updateAccountFailure = (payload) => { return { type: Type.UPDATE_ACCOUNT_FAILURE, payload } }
export const updateAccountPasswordRequest = (payload) => { return { type: Type.UPDATE_ACCOUNT_PASSWORD_REQUEST, payload } }
export const updateAccountPasswordSuccess = (payload) => { return { type: Type.UPDATE_ACCOUNT_PASSWORD_SUCCESS, payload } }
export const updateAccountPasswordFailure = (payload) => { return { type: Type.UPDATE_ACCOUNT_PASSWORD_FAILURE, payload } }

// member
export const fetchMembersRequest = () => { return { type: Type.FETCH_MEMBERS_REQUEST } }
export const fetchMembersSuccess = (payload) => { return { type: Type.FETCH_MEMBERS_SUCCESS, payload } }
export const fetchMembersFailure = (payload) => { return { type: Type.FETCH_MEMBERS_FAILURE, payload } }
export const fetchMemberRequest = (payload) => { return { type: Type.FETCH_MEMBER_REQUEST, payload } }
export const fetchMemberSuccess = (payload) => { return { type: Type.FETCH_MEMBER_SUCCESS, payload } }
export const fetchMemberFailure = (payload) => { return { type: Type.FETCH_MEMBER_FAILURE, payload } }
export const createMemberRequest = (payload) => { return { type: Type.CREATE_MEMBER_REQUEST, payload } }
export const createMemberSuccess = (payload) => { return { type: Type.CREATE_MEMBER_SUCCESS, payload } }
export const createMemberFailure = (payload) => { return { type: Type.CREATE_MEMBER_FAILURE, payload } }
export const updateMemberRequest = (payload) => { return { type: Type.UPDATE_MEMBER_REQUEST, payload} }
export const updateMemberSuccess = (payload) => { return { type: Type.UPDATE_MEMBER_SUCCESS, payload} }
export const updateMemberFailure = (payload) => { return { type: Type.UPDATE_MEMBER_FAILURE, payload} }
export const deleteMemberRequest = (payload) => { return { type: Type.DELETE_MEMBER_REQUEST, payload } }
export const deleteMemberSuccess = (payload) => { return { type: Type.DELETE_MEMBER_SUCCESS, payload } }
export const deleteMemberFailure = (payload) => { return { type: Type.DELETE_MEMBER_FAILURE, payload } }

// activity
export const fetchActivitiesRequest = () => { return { type: Type.FETCH_ACTIVITIES_REQUEST } }
export const fetchActivitiesSuccess = (payload) => { return { type: Type.FETCH_ACTIVITIES_SUCCESS, payload } }
export const fetchActivitiesFailure = (payload) => { return { type: Type.FETCH_ACTIVITIES_FAILURE, payload } }
export const createActivityRequest = (payload) => { return { type: Type.CREATE_ACTIVITY_REQUEST, payload } }
export const createActivitySuccess = (payload) => { return { type: Type.CREATE_ACTIVITY_SUCCESS, payload } }
export const createActivityFailure = (payload) => { return { type: Type.CREATE_ACTIVITY_FAILURE, payload } }
export const updateActivityRequest = (payload) => { return { type: Type.UPDATE_ACTIVITY_REQUEST, payload } }
export const updateActivitySuccess = (payload) => { return { type: Type.UPDATE_ACTIVITY_SUCCESS, payload } }
export const updateActivityFailure = (payload) => { return { type: Type.UPDATE_ACTIVITY_FAILURE, payload } }
export const deleteActivityRequest = (payload) => { return { type: Type.DELETE_ACTIVITY_REQUEST, payload } }
export const deleteActivitySuccess = (payload) => { return { type: Type.DELETE_ACTIVITY_SUCCESS, payload } }
export const deleteActivityFailure = (payload) => { return { type: Type.DELETE_ACTIVITY_FAILURE, payload } }

// society
export const fetchSocietiesRequest = () => { return { type: Type.FETCH_SOCIETIES_REQUEST } }
export const fetchSocietiesSuccess = (payload) => { return { type: Type.FETCH_SOCIETIES_SUCCESS, payload } }
export const fetchSocietiesFailure = (payload) => { return { type: Type.FETCH_SOCIETIES_FAILURE, payload } }
export const createSocietyRequest = (payload) => { return { type: Type.CREATE_SOCIETY_REQUEST, payload } }
export const createSocietySuccess = (payload) => { return { type: Type.CREATE_SOCIETY_SUCCESS, payload } }
export const createSocietyFailure = (payload) => { return { type: Type.CREATE_SOCIETY_FAILURE, payload } }
export const updateSocietyRequest = (payload) => { return { type: Type.UPDATE_SOCIETY_REQUEST, payload } }
export const updateSocietySuccess = (payload) => { return { type: Type.UPDATE_SOCIETY_SUCCESS, payload } }
export const updateSocietyFailure = (payload) => { return { type: Type.UPDATE_SOCIETY_FAILURE, payload } }
export const deleteSocietyRequest = (payload) => { return { type: Type.DELETE_SOCIETY_REQUEST, payload } }
export const deleteSocietySuccess = (payload) => { return { type: Type.DELETE_SOCIETY_SUCCESS, payload } }
export const deleteSocietyFailure = (payload) => { return { type: Type.DELETE_SOCIETY_FAILURE, payload } }


// reserach
export const fetchResearchesRequest = () => { return { type: Type.FETCH_RESEARCHES_REQUEST } }
export const fetchResearchesSuccess = (payload) => { return { type: Type.FETCH_RESEARCHES_SUCCESS, payload } }
export const fetchResearchesFailure = (payload) => { return { type: Type.FETCH_RESEARCHES_FAILURE, payload } }
export const createResearchRequest = (payload) => { return { type: Type.CREATE_RESEARCH_REQUEST, payload } }
export const createResearchSuccess = (payload) => { return { type: Type.CREATE_RESEARCH_SUCCESS, payload } }
export const createResearchFailure = (payload) => { return { type: Type.CREATE_RESEARCH_FAILURE, payload } }
export const updateResearchRequest = (payload) => { return { type: Type.UPDATE_RESEARCH_REQUEST, payload } }
export const updateResearchSuccess = (payload) => { return { type: Type.UPDATE_RESEARCH_SUCCESS, payload } }
export const updateResearchFailure = (payload) => { return { type: Type.UPDATE_RESEARCH_FAILURE, payload } }
export const deleteResearchRequest = (payload) => { return { type: Type.DELETE_RESEARCH_REQUEST, payload } }
export const deleteResearchSuccess = (payload) => { return { type: Type.DELETE_RESEARCH_SUCCESS, payload } }
export const deleteResearchFailure = (payload) => { return { type: Type.DELETE_RESEARCH_FAILURE, payload } }


// equipment
export const fetchEquipmentsRequest = () => { return { type: Type.FETCH_EQUIPMENTS_REQUEST } }
export const fetchEquipmentsSuccess = (payload) => { return { type: Type.FETCH_EQUIPMENTS_SUCCESS, payload } }
export const fetchEquipmentsFailure = (payload) => { return { type: Type.FETCH_EQUIPMENTS_FAILURE, payload } }
export const createEquipmentRequest = (payload) => { return { type: Type.CREATE_EQUIPMENT_REQUEST, payload } }
export const createEquipmentSuccess = (payload) => { return { type: Type.CREATE_EQUIPMENT_SUCCESS, payload } }
export const createEquipmentFailure = (payload) => { return { type: Type.CREATE_EQUIPMENT_FAILURE, payload } }
export const updateEquipmentRequest = (payload) => { return { type: Type.UPDATE_EQUIPMENT_REQUEST, payload } }
export const updateEquipmentSuccess = (payload) => { return { type: Type.UPDATE_EQUIPMENT_SUCCESS, payload } }
export const updateEquipmentFailure = (payload) => { return { type: Type.UPDATE_EQUIPMENT_FAILURE, payload } }
export const deleteEquipmentRequest = (payload) => { return { type: Type.DELETE_EQUIPMENT_REQUEST, payload } }
export const deleteEquipmentSuccess = (payload) => { return { type: Type.DELETE_EQUIPMENT_SUCCESS, payload } }
export const deleteEquipmentFailure = (payload) => { return { type: Type.DELETE_EQUIPMENT_FAILURE, payload } }

// lecture
export const fetchLecturesRequest = () => { return { type: Type.FETCH_LECTURES_REQUEST } }
export const fetchLecturesSuccess = (payload) => { return { type: Type.FETCH_LECTURES_SUCCESS, payload } }
export const fetchLecturesFailure = (payload) => { return { type: Type.FETCH_LECTURES_FAILURE, payload } }
export const fetchLectureRequest = (payload) => { return { type: Type.FETCH_LECTURE_REQUEST, payload } }
export const fetchLectureSuccess = (payload) => { return { type: Type.FETCH_LECTURE_SUCCESS, payload } }
export const fetchLectureFailure = (payload) => { return { type: Type.FETCH_LECTURE_FAILURE, payload } }
export const createLectureRequest = (payload) => { return { type: Type.CREATE_LECTURE_REQUEST, payload } }
export const createLectureSuccess = (payload) => { return { type: Type.CREATE_LECTURE_SUCCESS, payload } }
export const createLectureFailure = (payload) => { return { type: Type.CREATE_LECTURE_FAILURE, payload } }
export const updateLectureRequest = (payload) => { return { type: Type.UPDATE_LECTURE_REQUEST, payload } }
export const updateLectureSuccess = (payload) => { return { type: Type.UPDATE_LECTURE_SUCCESS, payload } }
export const updateLectureFailure = (payload) => { return { type: Type.UPDATE_LECTURE_FAILURE, payload } }
export const deleteLectureRequest = (payload) => { return { type: Type.DELETE_LECTURE_REQUEST, payload } }
export const deleteLectureSuccess = (payload) => { return { type: Type.DELETE_LECTURE_SUCCESS, payload } }
export const deleteLectureFailure = (payload) => { return { type: Type.DELETE_LECTURE_FAILURE, payload } }

// job
export const fetchJobsRequest = () => { return { type: Type.FETCH_JOBS_REQUEST } }
export const fetchJobsSuccess = (payload) => { return { type: Type.FETCH_JOBS_SUCCESS, payload } }
export const fetchJobsFailure = (payload) => { return { type: Type.FETCH_JOBS_FAILURE, payload } }
export const createJobRequest = (payload) => { return { type: Type.CREATE_JOB_REQUEST, payload } }
export const createJobSuccess = (payload) => { return { type: Type.CREATE_JOB_SUCCESS, payload } }
export const createJobFailure = (payload) => { return { type: Type.CREATE_JOB_FAILURE, payload } }
export const updateJobRequest = (payload) => { return { type: Type.UPDATE_JOB_REQUEST, payload } }
export const updateJobSuccess = (payload) => { return { type: Type.UPDATE_JOB_SUCCESS, payload } }
export const updateJobFailure = (payload) => { return { type: Type.UPDATE_JOB_FAILURE, payload } }
export const deleteJobRequest = (payload) => { return { type: Type.DELETE_JOB_REQUEST, payload } }
export const deleteJobSuccess = (payload) => { return { type: Type.DELETE_JOB_SUCCESS, payload } }
export const deleteJobFailure = (payload) => { return { type: Type.DELETE_JOB_FAILURE, payload } }

// tag
export const fetchTagsRequest = () => { return { type: Type.FETCH_TAGS_REQUEST } }
export const fetchTagsSuccess = (payload) => { return { type: Type.FETCH_TAGS_SUCCESS, payload } }
export const fetchTagsFailure = (payload) => { return { type: Type.FETCH_TAGS_FAILURE, payload } }
export const createTagRequest = (payload) => { return { type: Type.CREATE_TAG_REQUEST, payload } }
export const createTagSuccess = (payload) => { return { type: Type.CREATE_TAG_SUCCESS, payload } }
export const createTagFailure = (payload) => { return { type: Type.CREATE_TAG_FAILURE, payload } }
export const updateTagRequest = (payload) => { return { type: Type.UPDATE_TAG_REQUEST, payload } }
export const updateTagSuccess = (payload) => { return { type: Type.UPDATE_TAG_SUCCESS, payload } }
export const updateTagFailure = (payload) => { return { type: Type.UPDATE_TAG_FAILURE, payload } }
export const deleteTagRequest = (payload) => { return { type: Type.DELETE_TAG_REQUEST, payload } }
export const deleteTagSuccess = (payload) => { return { type: Type.DELETE_TAG_SUCCESS, payload } }
export const deleteTagFailure = (payload) => { return { type: Type.DELETE_TAG_FAILURE, payload } }

// else
export const showLoading = () => { return { type: Type.SHOW_LOADING} }
export const hideLoading = () => { return { type: Type.HIDE_LOADING} }
export const cerateAPIError = (payload) => { return { type: Type.API_ERROR, payload} }