import React from 'react'
import { connect } from 'react-redux'
import { fetchTagsRequest, deleteTagRequest } from '../../../../actions/action'
import BreadCrumb from '../../../common/Breadcrumb'
import AdminList from '../AdminList'
import { APIErrorList } from '../../../common/APIError'

const mapStateToProps = (state) => {
    return {
        tags: state.tags,
        apiError: state.apiError
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        fetchRequest: () => dispatch(fetchTagsRequest()),
        deleteRequest: (id) => dispatch(deleteTagRequest({ id }))
    }
}

class ConnectedTagList extends React.Component {
    render() {
        return (
            <div className="content">
                <BreadCrumb items={[{ path: "/", label: "管理者サイト" }, { path: "/tags", label: "タグ" }]} />
                <APIErrorList
                    apiError={this.props.apiError}/>
                <AdminList
                    items={this.props.tags}
                    caption={"タグ"}
                    path={"tags"}
                    fetchRequest={this.props.fetchRequest}
                    deleteRequest={this.props.deleteRequest} />
            </div>
        )
    }
}

const AdminTagList = connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedTagList)

export default AdminTagList