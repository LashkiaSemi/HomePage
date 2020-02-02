import React from 'react'
import { Route } from 'react-router-dom'

import Home from './contents/Home'
import Job from './contents/Job'
import Member from './contents/Member'

class Container extends React.Component {
    render() {
        return (
            <div className="container">
                <Route exact path="/" render={props => <Home {...props}/>}/>
                <Route exact path="/members" render={props => <Member {...props}/>} />
                <Route exact path="/jobs" render={props => <Job {...props} />} />
            </div>
        )
    }
}

export default Container