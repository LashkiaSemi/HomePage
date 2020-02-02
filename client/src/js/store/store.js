import { createStore, applyMiddleware } from 'redux'
import rootReducer from '../reducers/reducer'
import createSagaMiddleware from 'redux-saga'
import rootSaga from '../sagas/api-sagas'

// init saga
const sagaMiddleware = createSagaMiddleware()

const store = createStore(
    rootReducer,
    applyMiddleware(sagaMiddleware)
)

sagaMiddleware.run(rootSaga)

export default store