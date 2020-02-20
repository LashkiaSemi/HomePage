import React from 'react'
import { connect } from 'react-redux'
import AdminEdit from '../AdminEdit'
import BreadCrumb from '../../../common/Breadcrumb'
import { fetchActivitiesRequest, createActivityRequest, updateActivityRequest } from '../../../../actions/action'
import { APIErrorList } from '../../../common/APIError'

const mapStateToProps = (state) => {
    return {
        activities: state.activities,
        apiError: state.apiError,
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
                <APIErrorList
                    apiError={this.props.apiError}/>
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

const AdminActivityEdit = connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedActivityEdit)

export default AdminActivityEdit

