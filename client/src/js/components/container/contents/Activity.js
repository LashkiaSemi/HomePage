import React from 'react'
import { fetchActivitiesRequest } from '../../../actions/action'
import { connect } from 'react-redux'
import { APIErrorList } from '../../common/APIError'

const mapStateToProps = state => {
    return {
        isLoading: state.isLoading,
        activities: state.activities,
        apiError: state.apiError,
    }
}

const mapDispatchToProps = dispatch => {
    return {
        fetchRequest: () => dispatch(fetchActivitiesRequest())
    }
}

class ConnectedActivity extends React.Component {
    constructor(props) {
        super(props)
        this.state={
            activities: []
        }
    }

    componentDidMount() {
        this.props.fetchRequest()
    }

    componentDidUpdate(){
        if (this.state.activities.length > 0) {
            return
        }
        if (this.props.activities.length < 1) {
            return
        }
        
        /*
        年ごとのデータセットを作る
        { id: 2019newsって感じで年+news
            title: 見出し。2019年のニュースって感じになってる
            news: [{
            id: activityID
            date: 日付
            content: 内容 }]
        }
        という感じ
        */
        const acts = []
        var news = []
        var year = this.props.activities[0].date.substring(0, 4)
        this.props.activities.map(act => {
            if (act.date.substring(0, 4) !== year) {
                acts.push({ id: year+"news", title: year + "年のニュース", news: news })
                year = act.date.substring(0, 4)
                news = []
            }
            news.push({ id: act.id, date: act.date.replace(/-/g, "/"), content: act.activity })
        })
        acts.push({ id: year+"news", title: year + "年のニュース", news: news })

        this.setState({
            activities: acts
        })
    }

    render() {
        return (
            <div className="content">
                <h1 className="content-title h1-block">活動記録</h1>
                <APIErrorList 
                    apiError={this.props.apiError}/>
                <Toc activities={this.state.activities}/>
                <ActivityList activities={this.state.activities}/>
            </div>
        )
    }
}

/*
Toc 目次部分
props:
    activities = 親で作ったデータセット
*/
const Toc = (props) => {
    return (
        <div className="list">
            <ul>
                {
                    props.activities.map((act) => (
                        <TocRow key={act.id} activity={act} />
                    ))
                }
            </ul>
        </div>
    )
}

/*
TocRow 目次一件
props:
    activity = 一年分のデータ
*/
const TocRow = (props) => {
    return (
        <li>
            <a href={`#${props.activity.id}`} className="list-item">{props.activity.title}</a>
        </li>
    )
}

/*
ActivityList 内容部分
props:
    activities = 活動記録のデータセット
*/
const ActivityList = (props) => {
    return (
        <div className="list-stripe">
            {
                props.activities.map((act) => (
                    <ActivityRow key={act.id} activity={act} />
                ))
            }
        </div>
    )
}

/*
ActivityRow 一年の活動内容を表示
props:
    activity = 一年分のデータセット { id, title, news[] }ってやつ
*/
const ActivityRow = (props) => {
    return (
        <div className="list-item" id={props.activity.id}>
            <h2 className="list-title">{props.activity.title}</h2>
            <ul>
                {
                    props.activity.news.map((news) => (
                        <NewsRow key={news.id} news={news} />
                    ))
                }
            </ul>
        </div>
    )
}

/*
NewsRow 活動内容を一件表示
props:
    news = news[]の一件
*/
const NewsRow = (props) => {
    return (
        <li>&lt;{props.news.date}&gt; {props.news.content}</li>
    )
}

const Activity = connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedActivity)

export default Activity
