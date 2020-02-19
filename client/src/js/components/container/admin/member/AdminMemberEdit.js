import React from 'react'
import { connect } from 'react-redux'
import AdminEdit from '../AdminEdit'
import BreadCrumb from '../../../common/Breadcrumb'
import { fetchMembersRequest, createMemberRequest, updateMemberRequest } from '../../../../actions/action'
import { APIErrorList } from '../../../common/APIError'

const mapStateToProps = state => {
    return {
        members: state.members,
        apiError: state.apiError
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        fetchRequest: () => dispatch(fetchMembersRequest()),
        createRequest: (body) => dispatch(createMemberRequest({body})),
        updateRequest: (id, body) => dispatch(updateMemberRequest({id, body}))
    }
}

class ConnectedMemberEdit extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            values: {
                name: "",
                student_id: "",
                password: "",
                role: "",
                department: "",
                grade: 0,
                comment: ""
            },
            fields: [
                { label: "名前", type: "text", name: "name", required: true },
                { label: "学籍番号", type: "text", name: "student_id", required: true },
                { label: "権限", type: "select", name: "role", options: [{ label: "owner", value: "owner" }, { label: "admin", value: "admin" }, { label: "member", value: "member" },] }, // TODO: options
                { label: "学科", type: "text", name: "department" },
                { label: "学年", type: "select", name: "grade", requestType: "int", required: true, options: GRADE_OPTION },
                { label: "コメント", type: "textarea", name: "comment" },
            ],
        }
        
    }

    componentDidMount(){
        if (typeof this.props.match.params.id !== 'undefined') {
            return
        }
        var fields = this.state.fields
        fields.splice(2, 0, { label: "パスワード", type: "password", name: "password", required: true })
        this.setState({
            fields: fields,
        })
    }

    render() {
        return (
            <div className="content">
                <BreadCrumb items={[{ path: "/", label: "管理者サイト" }, { path: "/members", label: "メンバー" }]} />
                <APIErrorList
                    apiError={this.props.apiError}/>
                <AdminEdit
                    items={this.props.members}
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

const AdminMemberEdit = connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedMemberEdit)

export default AdminMemberEdit

// TODO: これ移行させようかなあ...
const GRADE_OPTION = [
    {
        value: 2,
        label: "学部2年",
    },
    {
        value: 3,
        label: "学部3年",
    },
    {
        value: 4,
        label: "学部4年",
    },
    {
        value: 5,
        label: "修士1年",
    },
    {
        value: 6,
        label: "修士2年",
    },
    {
        value: 7,
        label: "博士1年",
    },
    {
        value: 8,
        label: "博士2年",
    },
    {
        value: 0,
        label: "卒業生",
    },
]