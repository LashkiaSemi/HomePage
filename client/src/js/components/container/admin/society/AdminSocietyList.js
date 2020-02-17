import React from 'react'
import { fetchSocietiesRequest, createSocietyRequest, deleteSocietyRequest } from '../../../../actions/action'
import { connect } from 'react-redux'
import AdminList from '../AdminList'
import BreadCrumb from '../../../common/Breadcrumb'

const mapStateToProps = (state) => {
    return {
        societies: state.societies,
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        fetchRequest: () => dispatch(fetchSocietiesRequest()),
        deleteRequest: (id) => dispatch(deleteSocietyRequest({id}))
    }
}

class ConnectedSocietyList extends React.Component {
    render() {
        return (
            <div className="content">
                <BreadCrumb items={[{ path: "/", label: "管理者サイト" }, { path: "/societies", label: "学会発表" }]} />
                <AdminList
                    items={this.props.societies}
                    caption={"学会発表"}
                    path={"societies"}
                    fetchRequest={this.props.fetchRequest}
                    deleteRequest={this.props.deleteRequest}/>
            </div>
        )
    }
}

const AdminSocietyList = connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedSocietyList)

export default AdminSocietyList
