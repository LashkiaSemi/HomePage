import React from 'react'
import { fetchActivitiesRequest } from '../../../actions/action'
import { connect } from 'react-redux'

// tips: 
// 一応sagaもあるけど、使ってない。ここDB利用するなら
// 上手いことやってください

const mapDispatchToProps = dispatch => {
    return {
        fetchRequest: () => dispatch(fetchActivitiesRequest())
    }
}

const mapStateToProps = state => {
    return {
        isLoading: state.isLoading,
        activities: state.activities
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
        const acts = []
        var news = []
        var year = this.props.activities[0].date.substring(0, 4)
        this.props.activities.map(act => {
            if (act.date.substring(0, 4) !== year) {
                acts.push({ id: year, title: year + "年のニュース", news: news })
                year = act.date.substring(0, 4)
                news = []
            }
            news.push({ date: act.date.replace(/-/g, "/"), content: act.activity })
        })
        acts.push({ id: year+"news", title: year + "年のニュース", news: news })

        if (this.state.activities.length === 0) {
            this.setState({
                activities: acts
            })
        }
    }

    render() {
        return (
            <div className="content">
                <h1 className="content-title h1-block">活動記録</h1>
                <Toc activities={this.state.activities}/>
                <ActivityList activities={this.state.activities}/>
            </div>
        )
    }
}

// Toc
// 目次部分
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

const TocRow = (props) => {
    return (
        <li>
            <a href={`#${props.activity.id}`} className="list-item">{props.activity.title}</a>
        </li>
    )
}

// ActivityList
// 内容
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

const NewsRow = (props) => {
    return (
        <li>&lt;{props.news.date}&gt; {props.news.content}</li>
    )
}

const ACTIVITIES = [
    {
        id: "2018news",
        title: "2018年度のニュース",
        news: [
            {
                id: "20182",
                date: "2018/06/07",
                content: "ゼミ見学会",
            },
            {
                id: "20181",
                date: "2018/05/31",
                content: "質問対策講座",
            }
        ]
    },
    {
        id: "2017news",
        title: "2017年度のニュース",
        news: [
            {
                id: "20171",
                date: "2017/06/08",
                content: "ゼミ見学会1",
            },
            {
                id: "20172",
                date: "2017/06/15",
                content: "ゼミ見学会2",
            },
        ]
    },
    {
        id: "2015news",
        title: "2015年度のニュース",
        news: [
            {
                id: "20151",
                date: "2015/04/09",
                content: "ゼミ見学会1",
            },
            {
                id: "20152",
                date: "2015/04/09",
                content: "ゼミ見学会2",
            },
            {
                id: "20153",
                date: "2015/04/09",
                content: "ゼミ見学会2",
            },
        ]
    },

]

const Activity = connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedActivity)

export default Activity
