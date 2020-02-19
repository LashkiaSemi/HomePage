import React from 'react'
import { connect } from 'react-redux'
import AdminEdit from '../AdminEdit'
import BreadCrumb from '../../../common/Breadcrumb'
import { APIErrorList } from '../../../common/APIError'
import { fetchResearchesRequest, createResearchRequest, updateResearchRequest, fetchMembersRequest } from '../../../../actions/action'

const mapStateToProps = (state) => {
    return {
        researches: state.researches,
        apiError: state.apiError
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        fetchResearchesRequest: () => dispatch(fetchResearchesRequest()),
        fetchMembersRequest: () => dispatch(fetchMembersRequest()),
        createRequest: (body) => dispatch(createResearchRequest({body})),
        updateRequest: (id, body) => dispatch(updateResearchRequest({id, body})),
    }
}

class ConnectedResearchEdit extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            values: {
                title: "",
                author: "",
                comment: "",
                is_public: true,
            },
            fields: [
                { label: "タイトル", type: "text", name: "title", required: true },
                { label: "著者", type: "text", name: "author", required: true },
                { label: "コメント", type: "textarea", name: "comment" },
                { label: "ファイル", type: "file", name: "file" },
                { label: "公開 / 非公開", type: "checkbox", name: "is_public", requestType: "bool" }
            ],
        }
        this.fileInput = React.createRef()
    }

    render() {
        return (
            <div className="content">
                <BreadCrumb items={[{ path: "/", label: "管理者サイト" }, { path: "/researches", label: "卒業研究" }]} />
                <APIErrorList
                    apiError={this.props.apiError}/>
                <AdminEdit
                    items={this.props.researches}
                    itemID={this.props.match.params.id}
                    fields={this.state.fields}
                    values={this.state.values}
                    fileInput={this.fileInput}
                    fetchRequest={this.props.fetchResearchesRequest}
                    createRequest={this.props.createRequest}
                    updateRequest={this.props.updateRequest} />
            </div>
        )
    }
}

const AdminResearchEdit = connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedResearchEdit)

export default AdminResearchEdit