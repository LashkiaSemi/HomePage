import { fork } from 'redux-saga/effects'

import { watchJobs } from './api-jobs'
import { watchMembers } from './api-members'
import { watchActivities } from './api-activities'
import { watchSocieties } from './api-societies'
import { watchResearches } from './api-researches'
import { watchEquipments } from './api-equipments'
import { watchLectures } from './api-lectures'

// rootSaga
// sagaのwatcherをまとめておく
export default function* rootSaga() {
    yield fork(watchJobs)
    yield fork(watchMembers)
    yield fork(watchActivities)
    yield fork(watchSocieties)
    yield fork(watchResearches)
    yield fork(watchEquipments)
    yield fork(watchLectures)
}