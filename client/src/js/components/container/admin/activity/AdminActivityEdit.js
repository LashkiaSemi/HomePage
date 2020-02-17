import React from 'react'
import { connect } from 'react-redux'
import AdminEdit from '../AdminEdit'
import BreadCrumb from '../../../common/Breadcrumb'
import { fetchActivitiesRequest, createActivityRequest, updateActivityRequest } from '../../../../actions/action'
import { checkEmptyString } from '../../../../util/validation'
import { findItemByID } from '../../../../util/findItem'

const mapStateToProps = (state) => {
    return {
        activities: state.activities
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        fetchRequest: () => dispatch(fetchActivitiesRequest()),
        createRequest: (body) => dispatch(createActivityRequest({body})),
        updateRequest: (id, body) => dispatch(updateActivityRequest({id, body}))
    }
}

class ConnectedActivityEdit extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            values: {
                activity: "",
                date: "",
            },
            fields: [
                { label: "活動内容", type: "text", name: "activity", required: true},
                { label: "日付", type: "date", name: "date", required: true},
            ]
        }
    }

    render() {
        return (
            <div className="content">
                <BreadCrumb items={[{ path: "/", label: "管理者サイト" }, { path: "/activities", label: "活動記録" }]} />
                <AdminEdit
                    items={this.props.activities}
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

// class ConnectedActivityEdit extends React.Component {
//     constructor(props) {
//         super(props)
//         this.state = {
//             id: props.match.params.id,
//             date: "",
//             activity: "",
//             errors: [],
//             isInitialized: false,
//             isCreate: false
//         }

//         this.handleChange = this.handleChange.bind(this)
//         this.handleSubmit = this.handleSubmit.bind(this)
//     }

//     componentDidMount(){
//         if (typeof this.state.id === 'undefined') {
//             this.setState({
//                 isInitialized: true,
//                 isCreate: true
//             })
//             return
//         }
//         this.props.fetchRequest()
//     }

//     componentDidUpdate(){
//         if (this.state.isInitialized) {
//             return
//         }
//         if (!this.state.isInitialized && Object.keys(this.props.activities).length) {
//             const item = findItemByID(this.props.activities, this.state.id)
//             this.setState({
//                 activity: item.activity,
//                 date: item.date,
//                 isInitialized: true
//             })
//         }
//     }

//     handleChange(e) {
//         const field = e.target.name
//         this.setState({
//             [field]: e.target.value
//         })
//     }

//     handleSubmit(e) {
//         e.preventDefault()
//         const errors = checkEmptyString([
//             { id: "date", value: this.state.date, field: "日付" },
//             { id: "activity", value: this.state.activity, field: "活動内容" },
//         ])
//         if (errors.length > 0) {
//             this.setState({ errors })
//             return
//         }
    
//         const body = {
//             // todo: 日付のコンバータあったほうがいい？
//             date: this.state.date.replace(/-/g, '/'),
//             activity: this.state.activity
//         }

//         if(this.state.isCreate) {
//             this.props.createRequest(body)
//         } else {
//             this.props.updateRequest(this.state.id, body)
//         }
//     }

//     render() {
//         return (
//             <div className="content">
//                 <BreadCrumb items={[{ path: "/", label: "管理者サイト" }, { path: "/activities", label: "活動記録" }]} />
//                 {
//                     !this.state.isInitialized
//                     ? <></>
//                     : <form className="form-admin" onSubmit={this.handleSubmit}>
//                         {
//                             this.state.errors.length
//                                 ? <ErrorList errors={this.state.errors} />
//                                 : <></>
//                         }
//                         <table className="mb-20">
//                             <tbody>
//                                 <tr className="form-admin-item">
//                                     <td><label className="input-admin-label">活動内容</label></td>
//                                     <td><input type="text" className="input-admin-text" name="activity" value={this.state.activity} onChange={this.handleChange} /></td>
//                                 </tr>
//                                 <tr className="form-admin-item">
//                                     <td><label className="input-admin-label">日付</label></td>
//                                     <td><input type="date" className="input-admin-date" name="date" value={this.state.date} onChange={this.handleChange} /></td>
//                                 </tr>
//                             </tbody>
//                         </table>
//                         <div className="al-center">
//                             <button type="submit" className="btn btn-primary">保存</button>
//                         </div>
//                     </form>
//                 }
//             </div>
//         )
//     }
// }

const AdminActivityEdit = connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedActivityEdit)

export default AdminActivityEdit

