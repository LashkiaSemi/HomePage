import React from 'react'
import { connect } from 'react-redux'
import { fetchJobsRequest } from '../../../actions/action'
import { APIErrorList } from '../../common/APIError'

const mapDispatchToProps = dispatch => {
    return {
        fetchRequest: () => dispatch(fetchJobsRequest())
    }
}

const mapStateToProps = state => {
    return {
        isLoading: state.isLoading,
        jobs: state.jobs,
        apiError: state.apiError
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
                <APIErrorList
                    apiError={this.props.apiError}/>

                <div className="content-header">
                    <h2 className="h2">就職先一覧</h2>
                </div>
                <JobList jobs={this.props.jobs}/>
            </div>
        )
    }
}

/*
JobList 就職先一覧
props:
    jobs = 就職先のデータセット
*/
const JobList = (props) => {
    return (
        <div className="list">
            <ul>
                {
                    props.jobs.map((job) => (
                        <JobRow key={job.id} job={job} />
                    ))
                }
                <li className="list-item">etc...</li>
            </ul>
        </div>
    )
}

/*
JobRow 就職先一件
props:
    job = 就職先一件
*/
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

