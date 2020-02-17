import React from 'react'
import { connect } from 'react-redux'
import AdminEdit from '../AdminEdit'
import BreadCrumb from '../../../common/Breadcrumb'
import { fetchResearchesRequest, createResearchRequest, updateResearchRequest, fetchMembersRequest } from '../../../../actions/action'

const mapStateToProps = (state) => {
    return {
        researches: state.researches,
        members: state.members
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
            },
            fields: [
                { label: "タイトル", type: "text", name: "title" },
                { label: "著者", type: "text", name: "author" },
                { label: "コメント", type: "textarea", name: "comment" },
                { label: "ファイル", type: "file", name: "file" },
            ],
            isInitOptions: false
        }
        this.fileInput = React.createRef()
    }

    // componentDidMount() {
    //     this.props.fetchMembersRequest()
    // }

    // componentDidUpdate() {
    //     // selectのoptionを作る
    //     if(this.state.isInitOptions) {
    //         return
    //     }
    //     if(this.props.members.length) {
    //         this.setState({
    //             fields: this.state.fields.concat({ label: "著者", type: "select", name: "author_id", options: this.props.members }),
    //             isInitOptions: true,
    //         })
    //     }
    // }

    render() {
        return (
            <div className="content">
                <BreadCrumb items={[{ path: "/", label: "管理者サイト" }, { path: "/researches", label: "卒業研究" }]} />
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