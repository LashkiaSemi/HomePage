import React from 'react'
import { Route, Switch } from 'react-router-dom'
import { connect } from 'react-redux'

// main content
import Home from './contents/Home'
import Job from './contents/Job'
import Member from './contents/Member'
import Activity from './contents/Activity'
import Society from './contents/Society'
import Research from './contents/Research'
import Link from './contents/Link'
import Equipment from './contents/Equipment'
import Lecture from './contents/Lecture'
import LectureEdit from './contents/LectureEdit'
import MemberDetail from './contents/MemberDetail'
import MemberEdit from './contents/MemberEdit'
import Login from './contents/Login'
import Logout from './contents/Logout'
import PasswordEdit from './contents/PasswordEdit'

// admin site
import AdminHome from './admin/AdminHome'
import AdminActivityList from './admin/activity/AdminActivityList'
import AdminActivityEdit from './admin/activity/AdminActivityEdit'
import AdminSocietyList from './admin/society/AdminSocietyList'
import AdminSocietyEdit from './admin/society/AdminSocietyEdit'
import AdminResearchList from './admin/research/AdminResearchList'
import AdminResearchEdit from './admin/research/AdminResearchEdit'
import AdminMemberList from './admin/member/AdminMemberList'
import AdminMemberEdit from './admin/member/AdminMemberEdit'
import AdminJobList from './admin/job/AdminJobList'
import AdminJobEdit from './admin/job/AdminJobEdit'
import AdminEquipmentList from './admin/equipment/AdminEquipmentList'
import AdminEquipmentEdit from './admin/equipment/AdminEquipmentEdit'
import AdminLectureList from './admin/lecture/AdminLectureList'
import AdminLectureEdit from './admin/lecture/AdminLectureEdit'
import AdminTagList from './admin/tag/AdminTagList'
import AdminTagEdit from './admin/tag/AdminTagEdit'

// else
import Auth from '../common/Auth'
import Admin  from '../common/Admin'
import Error404 from './contents/Error'
import Download from './contents/Download'


class ConnectedContainer extends React.Component {
    constructor(props) {
        super(props)
    }
    render() {
        return (
            <div className="container">
                <Switch>
                    <Route path="/api" component={Error404} />
                    <Route exact path="/download" component={Download} />
                    {/* main pages */}
                    <Route exact path="/" component={Home}/>
                    <Route exact path="/members" component={Member}/>
                    <Route exact path="/members/:id" component={MemberDetail}/>
                    <Route exact path="/jobs" component={Job}/>
                    <Route exact path="/activities" component={Activity} />
                    <Route exact path="/societies" component={Society} />
                    <Route exact path="/researches" component={Research}/>
                    <Route exact path="/links" component={Link}/>
                    <Route exact path="/login" component={Login}/>

                    {/* Auth */}
                    <AuthRoute exact path="/equipments" component={Equipment} />
                    <AuthRoute exact path="/lectures" component={Lecture} />
                    <AuthRoute exact path="/lectures/new" component={LectureEdit}/>
                    <AuthRoute exact path="/lectures/:id/edit" component={LectureEdit} />
                    <AuthRoute exact path="/account/edit" component={MemberEdit} />
                    <AuthRoute exact path="/account/edit_pass" component={PasswordEdit} />
                    <AuthRoute exact path="/logout" component={Logout} />

                    {/* Admin */}
                    <AdminRoute exact path="/admin" component={AdminHome}/>
                    <AdminRoute exact path="/admin/activities" component={AdminActivityList} />
                    <AdminRoute exact path="/admin/activities/new" component={AdminActivityEdit} />
                    <AdminRoute exact path="/admin/activities/:id/edit" component={AdminActivityEdit} />
                    <AdminRoute exact path="/admin/societies" component={AdminSocietyList} />
                    <AdminRoute exact path="/admin/societies/new" component={AdminSocietyEdit} />
                    <AdminRoute exact path="/admin/societies/:id/edit" component={AdminSocietyEdit} />
                    <AdminRoute exact path="/admin/researches" component={AdminResearchList} />
                    <AdminRoute exact path="/admin/researches/new" component={AdminResearchEdit} />
                    <AdminRoute exact path="/admin/researches/:id/edit" component={AdminResearchEdit} />
                    <AdminRoute exact path="/admin/members" component={AdminMemberList} />
                    <AdminRoute exact path="/admin/members/new" component={AdminMemberEdit} />
                    <AdminRoute exact path="/admin/members/:id/edit" component={AdminMemberEdit} />
                    <AdminRoute exact path="/admin/jobs" component={AdminJobList} />
                    <AdminRoute exact path="/admin/jobs/new" component={AdminJobEdit} />
                    <AdminRoute exact path="/admin/jobs/:id/edit" component={AdminJobEdit} />
                    <AdminRoute exact path="/admin/equipments" component={AdminEquipmentList} />
                    <AdminRoute exact path="/admin/equipments/new" component={AdminEquipmentEdit} />
                    <AdminRoute exact path="/admin/equipments/:id/edit" component={AdminEquipmentEdit} />
                    <AdminRoute exact path="/admin/lectures" component={AdminLectureList} />
                    <AdminRoute exact path="/admin/lectures/new" component={AdminLectureEdit} />
                    <AdminRoute exact path="/admin/lectures/:id/edit" component={AdminLectureEdit} />
                    <AdminRoute exact path="/admin/tags" component={AdminTagList} />
                    <AdminRoute exact path="/admin/tags/new" component={AdminTagEdit} />
                    <AdminRoute exact path="/admin/tags/:id/edit" component={AdminTagEdit} />

    
                    <Route component={Error404} />
                </Switch>
            </div>
        )
    }
}

function AuthRoute({component: Component, ...rest}){
    return (
        <Route 
            {...rest}
            render={routeProps => (
                <Auth>
                    <Component {...routeProps} />
                </Auth>
            )}
        />
    )
}

function AdminRoute({component: Component, ...rest}) {
    return (
        <Route
            {...rest}
            render={routeProps => (
                <Admin>
                    <Component {...routeProps} />
                </Admin>
            )}
        />
    )
}

const Container = connect(
)(ConnectedContainer)

export default Container