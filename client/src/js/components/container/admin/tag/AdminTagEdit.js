import React from 'react'
import { connect } from 'react-redux'
import AdminEdit from '../AdminEdit'
import BreadCrumb from '../../../common/Breadcrumb'
import { fetchTagsRequest, createTagRequest, updateTagRequest } from '../../../../actions/action'

const mapStateToProps = state => {
    return {
        tags: state.tags
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        fetchRequest: () => dispatch(fetchTagsRequest()),
        createRequest: (body) => dispatch(createTagRequest({ body })),
        updateRequest: (id, body) => dispatch(updateTagRequest({ id, body }))
    }
}

class ConnectedTagEdit extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            values: {
                name: ""
            },
            fields: [
                { label: "タグ名", type: "text", name: "name", required: true },
            ],
        }
    }

    render() {
        return (
            <div className="content">
                <BreadCrumb items={[{ path: "/", label: "管理者サイト" }, { path: "/tags", label: "タグ" }]} />
                <AdminEdit
                    items={this.props.tags}
                    itemID={this.props.match.params.id}
                    fields={this.state.fields}
                    values={this.state.values}
                    fetchRequest={this.props.fetchRequest}
                    createRequest={this.props.createRequest}
                    updateRequest={this.props.updateRequest} />
            </div>
        )
    }
}

const AdminTagEdit = connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedTagEdit)

export default AdminTagEdit
