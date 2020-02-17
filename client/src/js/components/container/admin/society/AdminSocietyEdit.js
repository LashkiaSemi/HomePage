import React from 'react'
import { connect } from 'react-redux'
import AdminEdit from '../AdminEdit'
import { fetchSocietiesRequest, createSocietyRequest, updateSocietyRequest } from '../../../../actions/action'
import BreadCrumb from '../../../common/Breadcrumb'

const mapStateToProps = (state) => {
    return {
        societies: state.societies,
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        fetchRequest: () => dispatch(fetchSocietiesRequest()),
        createRequest: (body) => dispatch(createSocietyRequest({body})),
        updateRequest: (id, body) => dispatch(updateSocietyRequest({id, body}))
    }
}

class ConnectedSocietyEdit extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            values: {
                title: "",
                society: "",
                date: "",
                award: "",
                author: "",
            },
            fields: [
                { label: "タイトル", type: "text", name: "title" },
                { label: "日付", type: "date", name: "date" },
                { label: "著者", type: "text", name: "author" },
                { label: "発表学会", type: "text", name: "society" },
                { label: "受賞", type: "text", name: "award" }
            ]
        }
    }

    render() {
        return (
            <div className="content">
                <BreadCrumb items={[{ path: "/", label: "管理者サイト" }, { path: "/societies", label: "学会発表" }]}/>
                <AdminEdit
                    items={this.props.societies}
                    itemID={this.props.match.params.id}
                    fields={this.state.fields}
                    values={this.state.values}
                    fetchRequest={this.props.fetchRequest}
                    createRequest={this.props.createRequest}
                    updateRequest={this.props.updateRequest}/>
            </div>
        )
    }
}

const AdminSocietyEdit = connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedSocietyEdit)

export default AdminSocietyEdit