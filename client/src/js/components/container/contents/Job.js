import React from 'react'
import { connect } from 'react-redux'
import { fetchJobsRequest } from '../../../actions/action'

const mapDispatchToProps = dispatch => {
    return {
        fetchRequest: () => dispatch(fetchJobsRequest())
    }
}

const mapStateToProps = state => {
    return {
        isLoading: state.isLoading,
        jobs: state.jobs
    }
}

class ConnectedJob extends React.Component {
    componentDidMount(){
        this.props.fetchRequest()
    }
    render(){
        return (
            <div className="content">
                <h1 className="content-title h1-block">就職先</h1>
                <div className="content-header">
                    <h2 className="h2">就職先一覧</h2>
                </div>
                <div className="list">
                    <ul>
                        {
                            this.props.isLoading
                            ? <li>Loading</li>
                            :
                            this.props.jobs.map((job) => (
                                <JobRow key={job.id} job={job} />
                            ))
                        }
                        <li className="list-item">etc...</li>
                    </ul>
                </div>
            </div>
        )
    }
}


const JobRow = (props) => {
    return (
        <li className="list-item">{props.job.company} / {props.job.job}</li>
    )
}

const Job = connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedJob)

export default Job

