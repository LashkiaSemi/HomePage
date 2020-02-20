import React from 'react'
import { connect } from 'react-redux'
import { fetchResearchesRequest } from '../../../actions/action'
import { APIErrorList } from '../../common/APIError'

const mapDispatchToProps = dispatch => {
    return {
        fetchRequest: () => dispatch(fetchResearchesRequest())
    }
}

const mapStateToProps = state => {
    return {
        isLoading: state.isLoading,
        researches: state.researches,
        apiError: state.apiError
    }
}

class ConnectedResearch extends React.Component {
    componentDidMount(){
        this.props.fetchRequest()
    }

    render() {
        return (
            <div className="content">
                <h1 className="content-title h1-block">卒業研究</h1>
                <APIErrorList
                    apiError={this.props.apiError}/>
                <ResearchTable researches={this.props.researches} />
            </div>
        )
    }
}

/*
ResearchTable 卒研データのテーブル
props:
    researches = 卒研のデータセット
*/
const ResearchTable = (props) => {
    return (
        <table className="table-stripe">
            <thead>
                <tr>
                    <th>タイトル</th>
                    <th>著者</th>
                    <th>コメント</th>
                    <th>投稿日</th>
                    <th></th>
                </tr>
            </thead>
            <tbody>
                {
                    props.researches.map((res) => (
                        <ResearchRow key={res.id} research={res} />
                    ))
                }
            </tbody>
        </table>
    )
}

/*
ResearchRow 一件
props:
    research = 卒研一件
*/
const ResearchRow = (props) => {
    return (
        <tr>
            <td>{props.research.title}</td>
            <td>{props.research.author}</td>
            <td>{props.research.comment}</td>
            <td>{props.research.created_at}</td>
            {/* TODO: download script */}
            <td><button className="btn btn-primary">Download</button></td>
        </tr>
    )
}

const Research = connect(
    mapStateToProps,
    mapDispatchToProps,
)(ConnectedResearch)

export default Research