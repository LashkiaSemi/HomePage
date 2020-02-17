import React from 'react'
import { connect } from 'react-redux'
import AdminEdit from '../AdminEdit'
import BreadCrumb from '../../../common/Breadcrumb'
import { fetchLecturesRequest, createLectureRequest, updateLectureRequest, fetchMembersRequest } from '../../../../actions/action'

const mapStateToProps = state => {
    return {
        lectures: state.lectures,
        members: state.members
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        fetchRequest: () => dispatch(fetchLecturesRequest()),
        fetchMembersRequest: () => dispatch(fetchMembersRequest()),
        createRequest: (body) => dispatch(createLectureRequest({ body })),
        updateRequest: (id, body) => dispatch(updateLectureRequest({ id, body }))
    }
}

class ConnectedLectureEdit extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            values: {
                user_id: 0,
                title: "",
                comment: ""
            },
            fields: [
                { label: "タイトル", type: "text", name: "title" },
                { label: "コメント", type: "textarea", name: "comment" },
                { label: "ファイル", type: "file", name: "file" }

            ],
            isInitialized: false,
        }
        this.fileInput = React.createRef()
    }

    componentDidMount() {
        this.props.fetchMembersRequest()
    }

    componentDidUpdate() {
        if (this.state.isInitialized) {
            return
        }
        if (Object.keys(this.props.members).length > 0) {
            var options = []
            this.props.members.map(member => {
                options.push({ label: member.name, value: parseInt(member.id) })
            })

            this.setState({
                fields: this.state.fields.concat({ label: "投稿者", type: "select", name: "user_id", requestType: "int", options }),
                isInitialized: true
            })
        }
    }

    render() {
        return (
            <div className="content">
                <BreadCrumb items={[{ path: "/", label: "管理者サイト" }, { path: "/lectures", label: "レクチャー" }]} />
                {
                    !this.state.isInitialized
                        ? <></>
                        : <AdminEdit
                            items={this.props.lectures}
                            itemID={this.props.match.params.id}
                            fields={this.state.fields}
                            values={this.state.values}
                            fileInput={this.fileInput}
                            fetchRequest={this.props.fetchRequest}
                            createRequest={this.props.createRequest}
                            updateRequest={this.props.updateRequest} />
                }
            </div>
        )
    }
}

const AdminLectureEdit = connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedLectureEdit)

export default AdminLectureEdit
