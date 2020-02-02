import React from 'react'
import { connect } from 'react-redux'
import { fetchSocieties } from '../../../actions/action'

const mapStateToProps = state => {
    return {
        isLoading: state.isLoading,
        societies: state.societies
    }
}

class ConnectedSociey extends React.Component {
    componentDidMount(){
        this.props.fetchSocieties()
    }

    render(){
        return (
            <div className="content">
                <h1 className="content-title h1-block">学会発表</h1>
                {
                    // TODO: loading
                    <SocietyTable societies={this.props.societies} />
                }
            </div>
        )
    }
}

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
    { fetchSocieties }
)(ConnectedSociey)

export default Society