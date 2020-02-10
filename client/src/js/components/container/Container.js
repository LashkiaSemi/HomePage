import React from 'react'
import { Route, Switch } from 'react-router-dom'

import Home from './contents/Home'
import Job from './contents/Job'
import Member from './contents/Member'
import Activity from './contents/Activity'
import Society from './contents/Societies'
import Research from './contents/Researches'
import Link from './contents/Link'
import Equipment from './contents/Equipment'
import Lecture from './contents/Lecture'
import LectureEdit from './contents/LectureEdit'
import MemberDetail from './contents/MemberDetail'
import MemberEdit from './contents/MemberEdit'
import Login from './contents/Login'
import Logout from './contents/Logout'
import PasswordEdit from './contents/PasswordEdit'

import Auth from '../common/Auth'
import Error404 from './contents/Error'

class Container extends React.Component {
    render() {
        return (
            <div className="container">
                <Switch>
                    <Route exact path="/" component={Home}/>
                    <Route exact path="/members" component={Member}/>
                    <Route exact path="/members/:id" component={MemberDetail}/>
                    <Route exact path="/jobs" component={Job}/>
                    <Route exact path="/activities" component={Activity} />
                    <Route exact path="/societies" component={Society} />
                    <Route exact path="/researches" component={Research}/>
                    <Route exact path="/links" component={Link}/>
                    <Route exact path="/login" component={Login}/>

                    <Auth>
                        <Route exact path="/equipments" component={Equipment} />
                        <Route exact path="/lectures" component={Lecture} />
                        <Route exact path="/lectures/new" component={LectureEdit} />
                        <Route exact path="/lectures/:id/edit" component={LectureEdit} />
                        <Route exact path="/account/edit" component={MemberEdit} />
                        <Route exact path="/account/edit_pass" component={PasswordEdit} />
                        <Route exact path="/logout" component={Logout} />
                    </Auth>

                    <Route component={Error404} />
                </Switch>
            </div>
        )
    }
}

export default Container