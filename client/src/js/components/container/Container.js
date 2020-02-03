import React from 'react'
import { Route } from 'react-router-dom'

import Home from './contents/Home'
import Job from './contents/Job'
import Member from './contents/Member'
import Activity from './contents/Activity'
import Society from './contents/Societies'
import Research from './contents/Researches'
import Link from './contents/Link'
import Equipment from './contents/Equipment'
import Lecture from './contents/Lecture'
import MemberDetail from './contents/MemberDetail'
import MemberEdit from './contents/MemberEdit'
import Login from './contents/Login'

class Container extends React.Component {
    render() {
        return (
            <div className="container">
                <Route exact path="/" render={props => <Home {...props}/>}/>
                <Route exact path="/members" render={props => <Member {...props}/>} />
                <Route exact path="/members/:id" component={MemberDetail}/>
                <Route exact path="/members/:id/edit" component={MemberEdit} />
                <Route exact path="/jobs" render={props => <Job {...props} />} />
                <Route exact path="/activities" component={Activity} />
                <Route exact path="/societies" component={Society} />
                <Route exact path="/researches" component={Research}/>
                <Route exact path="/links" component={Link}/>
                <Route exact path="/equipments" component={Equipment}/>
                <Route exact path="/lectures" component={Lecture} />

                <Route exact path="/login" component={Login}/>
            </div>
        )
    }
}

export default Container