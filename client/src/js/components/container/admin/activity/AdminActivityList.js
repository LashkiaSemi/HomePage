import React from 'react'
import { connect } from 'react-redux'
import AdminList from '../AdminList'
import BreadCrumb from '../../../common/Breadcrumb'
import { fetchActivitiesRequest, deleteActivityRequest } from '../../../../actions/action'

const mapStateToProps = (state) => {
    return {
        activities: state.activities,
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        fetchRequest: () => dispatch(fetchActivitiesRequest()),
        deleteRequest: (id) => dispatch(deleteActivityRequest({id})),
    }
}

class ConnectedActivityList extends React.Component {
    render() {
        return (
            <div className="content">
                <BreadCrumb items={[{ path: "/", label: "管理者サイト" }, { path: "/activities", label: "活動記録" }]} />
                <AdminList
                    items={this.props.activities}
                    caption={"活動記録"}
                    path={"activities"}
                    fetchRequest={this.props.fetchRequest}
                    deleteRequest={this.props.deleteRequest} />
            </div>
        )
    }
}

const AdminActivityList = connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedActivityList)

export default AdminActivityList