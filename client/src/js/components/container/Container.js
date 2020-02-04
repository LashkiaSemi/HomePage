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
import Logout from './contents/Logout'
import PasswordEdit from './contents/PasswordEdit'

class Container extends React.Component {
    render() {
        return (
            <div className="container">
                <Route exact path="/" component={Home}/>
                <Route exact path="/members" component={Member}/>
                <Route exact path="/members/:id" component={MemberDetail}/>
                <Route exact path="/account/edit" component={MemberEdit} />
                <Route exact path="/account/edit_pass" component={PasswordEdit} />
                <Route exact path="/jobs" component={Job}/>
                <Route exact path="/activities" component={Activity} />
                <Route exact path="/societies" component={Society} />
                <Route exact path="/researches" component={Research}/>
                <Route exact path="/links" component={Link}/>
                <Route exact path="/equipments" component={Equipment}/>
                <Route exact path="/lectures" component={Lecture} />

                <Route exact path="/login" component={Login}/>
                <Route exact path="/logout" component={Logout}/>
            </div>
        )
    }
}

export default Container