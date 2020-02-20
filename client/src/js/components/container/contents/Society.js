import React from 'react'
import { connect } from 'react-redux'
import { fetchSocietiesRequest } from '../../../actions/action'
import { APIErrorList } from '../../common/APIError'

const mapDispatchToProps = dispatch => {
    return {
        fetchRequest: () => dispatch(fetchSocietiesRequest())
    }
}

const mapStateToProps = state => {
    return {
        isLoading: state.isLoading,
        societies: state.societies,
        apiError: state.apiError
    }
}

class ConnectedSociey extends React.Component {
    componentDidMount(){
        this.props.fetchRequest()
    }

    render(){
        return (
            <div className="content">
                <h1 className="content-title h1-block">学会発表</h1>
                <APIErrorList
                    apiError={this.props.apiError} />
                <SocietyTable societies={this.props.societies} />
            </div>
        )
    }
}

/*
SocietyTable 学会発表のテーブル
props:
    societies = データセット
*/
const SocietyTable = (props) => {
    return (
        <table className="table-basic">
            <thead>
                <tr>
                    <th>日付</th>
                    <th>名前</th>
                    <th>タイトル</th>
                    <th>発表学会</th>
                    <th>受賞</th>
                </tr>
            </thead>
            <tbody>
                {
                    props.societies.map((soc) => (
                        <SocietyRow key={soc.id} society={soc} />
                    ))
                }
            </tbody>
        </table>
    )
}

/*
SocietyRow 一件表示
props:
    society = 学会発表のデータ一件
*/
const SocietyRow = (props) => {
    return (
        <tr>
            <td>{props.society.date}</td>
            <td>{props.society.author}</td>
            <td>{props.society.title}</td>
            <td>{props.society.society}</td>
            <td>{props.society.award}</td>
        </tr>
    )
}

const Society = connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedSociey)

export default Society