import React from 'react'
import { connect } from 'react-redux'
import { fetchMembersRequest, deleteMemberRequest } from '../../../../actions/action'
import BreadCrumb from '../../../common/Breadcrumb'
import AdminList from '../AdminList'

const mapStateToProps = (state) => {
    return {
        members: state.members
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        fetchRequest: () => dispatch(fetchMembersRequest()),
        deleteRequest: (id) => dispatch(deleteMemberRequest({id}))
    }
}

class ConnectedMemberList extends React.Component {
    render() {
        return (
            <div className="content">
                <BreadCrumb items={[{ path: "/", label: "管理者サイト" }, { path: "/members", label: "メンバー" }]} />
                <AdminList
                    items={this.props.members}
                    caption={"メンバー"}
                    path={"members"}
                    fetchRequest={this.props.fetchRequest}
                    deleteRequest={this.props.deleteRequest} />
            </div>
        )
    }
}

const AdminMemberList = connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedMemberList)

export default AdminMemberList