import React from 'react'
import { connect } from 'react-redux'
import AdminEdit from '../AdminEdit'
import BreadCrumb from '../../../common/Breadcrumb'
import { fetchJobsRequest, createJobRequest, updateJobRequest } from '../../../../actions/action'
import { APIErrorList } from '../../../common/APIError'

const mapStateToProps = state => {
    return {
        jobs: state.jobs,
        apiError: state.apiError
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        fetchRequest: () => dispatch(fetchJobsRequest()),
        createRequest: (body) => dispatch(createJobRequest({ body })),
        updateRequest: (id, body) => dispatch(updateJobRequest({ id, body }))
    }
}

class ConnectedJobEdit extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            values: {
                company: "",
                job: ""
            },
            fields: [
                { label: "企業名", type: "text", name: "company", required: true },
                { label: "職種", type: "text", name: "job" },
            ],
        }
    }

    render() {
        return (
            <div className="content">
                <BreadCrumb items={[{ path: "/", label: "管理者サイト" }, { path: "/jobs", label: "就職先" }]} />
                <APIErrorList
                    apiError={this.props.apiError}/>
                <AdminEdit
                    items={this.props.jobs}
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

const AdminJobEdit = connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedJobEdit)

export default AdminJobEdit
