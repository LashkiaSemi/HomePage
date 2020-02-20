import React from 'react'
import { connect } from 'react-redux'
import { fetchLecturesRequest, adminDeleteLectureRequest } from '../../../../actions/action'
import BreadCrumb from '../../../common/Breadcrumb'
import AdminList from '../AdminList'
import { APIErrorList } from '../../../common/APIError'

const mapStateToProps = (state) => {
    return {
        lectures: state.lectures,
        apiError: state.apiError
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        fetchRequest: () => dispatch(fetchLecturesRequest()),
        deleteRequest: (id) => dispatch(adminDeleteLectureRequest({ id }))
    }
}

class ConnectedLectureList extends React.Component {
    render() {
        return (
            <div className="content">
                <BreadCrumb items={[{ path: "/", label: "管理者サイト" }, { path: "/lectures", label: "レクチャー" }]} />
                <APIErrorList
                    apiError={this.props.apiError}/>
                <AdminList
                    items={this.props.lectures}
                    caption={"レクチャー"}
                    path={"lectures"}
                    fetchRequest={this.props.fetchRequest}
                    deleteRequest={this.props.deleteRequest} />
            </div>
        )
    }
}

const AdminLectureList = connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedLectureList)

export default AdminLectureList