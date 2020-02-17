import React from 'react'
import { connect } from 'react-redux'
import { fetchJobsRequest, deleteJobRequest } from '../../../../actions/action'
import BreadCrumb from '../../../common/Breadcrumb'
import AdminList from '../AdminList'

const mapStateToProps = (state) => {
    return {
        jobs: state.jobs
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        fetchRequest: () => dispatch(fetchJobsRequest()),
        deleteRequest: (id) => dispatch(deleteJobRequest({ id }))
    }
}

class ConnectedJobList extends React.Component {
    render() {
        return (
            <div className="content">
                <BreadCrumb items={[{ path: "/", label: "管理者サイト" }, { path: "/jobs", label: "就職先" }]} />
                <AdminList
                    items={this.props.jobs}
                    caption={"就職先"}
                    path={"jobs"}
                    fetchRequest={this.props.fetchRequest}
                    deleteRequest={this.props.deleteRequest} />
            </div>
        )
    }
}

const AdminJobList = connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedJobList)

export default AdminJobList