import React from 'react'
import { connect } from 'react-redux'
import { fetchLecturesRequest, deleteLectureRequest } from '../../../../actions/action'
import BreadCrumb from '../../../common/Breadcrumb'
import AdminList from '../AdminList'

const mapStateToProps = (state) => {
    return {
        lectures: state.lectures
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        fetchRequest: () => dispatch(fetchLecturesRequest()),
        deleteRequest: (id) => dispatch(deleteLectureRequest({ id }))
    }
}

class ConnectedLectureList extends React.Component {
    render() {
        return (
            <div className="content">
                <BreadCrumb items={[{ path: "/", label: "管理者サイト" }, { path: "/lectures", label: "レクチャー" }]} />
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