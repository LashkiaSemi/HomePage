import React from 'react'

// Newsで表示する内容
const NEWS = [
    {
        id: 1,
        date: "2019/06/10",
        news: "ゼミ見学",
    },
]

/*
Home トップページ。めっちゃ静的
*/
class Home extends React.Component {
    render() {
        return (
            <>
                <Jumbotron text={"ようこそ　ラシキアゼミへ"} />
                <News />
                <AboutLab />
            </>
        )
    }
}

/*
Jumbotron
props:
    text = ジャンボトロンに表示するwelcome文
*/
const Jumbotron = (props) => {
    return (
        <div className="jumbotron">
            <label>{props.text}</label>
        </div>
    )
}

/*
News トップにカードとして表示してる今後の予定
Newsのデータは一番上のNEWSでいじる...
memo: ここをDBに移行しても面白いかも?
*/
const News = () => {
    return (
        <div className="content">
            <div className="card">
                <div className="card-header">
                    <h2 className="h2">お知らせ</h2>
                </div>
                <div className="card-content">
                    {
                        NEWS.map((news) => (
                            <NewsRow key={news.id} news={news} />
                        ))
                    }
                </div>
            </div>
        </div>
    )
}

/*
NewsRow 今後の予定を一件
props:
    news = ニュースを一件
        {
            date: 日付
            news: 内容
        }
*/
const NewsRow = (props) => {
    return (
        <li>{props.news.date}: {props.news.news}</li>
    )
}

/*
AboutLab ゼミについてのコンテンツ
ガッチガチのstatic。基本的に変更いらず
*/
const AboutLab = () => {
    return (
        <div className="content">
            <div className="mb-30">
                <h2 className="h2-underline">ラシキア研究室</h2>
                <p>ラシキア研究室はIT全般について勉強していき、最終的には役に立つソフトウェアを開発することを目的としています。</p>
            </div>

            <div className="mb-30">
                <h2 className="h2-underline">活動場所</h2>
                <p>ゼミの事案等の全体で活動するときは、15号館5Fを利用します。個人の研究は、11号館4Fのゼミ室で行っています。稀に、ゼミ室前のガーデン端末を利用することもあります。</p>
                <p><a href="https://www.chukyo-u.ac.jp/information/facility/g2.html">豊田キャンパスマップ</a><br />
                    ゼミ室：　　　　11号館4F ラシキア研究室<br />
                    ゼミ活動場所：　15号館(AI棟・人工知能高等研究所) 5F</p>
            </div>

            <div className="mb-30">
                <h2 className="h2-underline">学習内容</h2>
                <p>ゼミでは以下のことを勉強していきます。</p>
                <ul>
                    <li>プログラミング言語の学習(Java, C, C++, Android, ios, PHP ...)</li>
                    <ul>
                        <li>全体的にシステム開発の割合が高い(問題定義からシステムの構築へ)</li>
                        <li>本格的な研究も可(画像処理, ...)</li>
                    </ul>
                    <li>プレゼンテーション法</li>
                    <li>就職活動を成功させるための学習</li>
                    <ul>
                        <li>SPI対策</li>
                    </ul>
                    <li>勉強会(Android, iPhone, サーバー...)</li>
                </ul>
            </div>


            <div className="mb-30">
                <h2 className="h2-underline">ゼミ時間割(例)</h2>
                <p>ゼミは、毎週木曜日に以下のうような予定で進行していきます。</p>
                <table style={{ "borderSpacing": ".8rem 0rem" }}>
                    <tbody>
                        <tr>
                            <td className="al-center">11:10</td>
                            <td></td>
                        </tr>
                        <tr>
                            <td className="al-center">|</td>
                            <td>2年生レクチャー開始(5F)</td>
                        </tr>
                        <tr>
                            <td className="al-center">12:40</td>
                            <td></td>
                        </tr>
                        <br />
                        <tr>
                            <td className="al-center">13:30</td>
                            <td></td>
                        </tr>
                        <tr>
                            <td className="al-center">|</td>
                            <td>集合、出席確認、連絡事項(5F)</td>
                        </tr>
                        <tr>
                            <td className="al-center">13:45</td>
                            <td></td>
                        </tr>
                        <tr>
                            <td className="al-center">|</td>
                            <td>3年開発報告発表、2年ITニュース発表、4年研究発表(5F)</td>
                        </tr>
                        <tr>
                            <td className="al-center">15:30</td>
                            <td></td>
                        </tr>
                        <br />
                        <tr>
                            <td className="al-center">15:45</td>
                            <td></td>
                        </tr>
                        <tr>
                            <td className="al-center">|</td>
                            <td>2、3年レクチャー(5F)、4年研究発表(全員)(4F)</td>
                        </tr>
                        <tr>
                            <td className="al-center">18:30</td>
                            <td>自習、発表ダメ出し</td>
                        </tr>
                    </tbody>
                </table>
            </div>


            <div className="mb-30">
                <h2 className="h2-underline">注意事項</h2>
                <p>以下のことに注意してください。</p>
                <ul>
                    <li>時間厳守！</li>
                    <ul>
                        <li>集合時間、発表時間を守る</li>
                        <li>だらだらとした時間を作らない</li>
                    </ul>
                </ul>
                <ul>
                    <li>発表途中のダメ出しは禁止！</li>
                    <ul>
                        <li>技術的なことならともかく、個人的なことだと聞いている人に迷惑</li>
                        <li>個人的なダメ出しは18:00以降、先生と議論</li>
                        <li>質問されてわからないことがあればメモをする</li>
                        <li>先生に聞きたいことも18:00以降、先生と議論</li>
                    </ul>
                </ul>
                <ul>
                    <li>ゼミを遅刻or欠席する場合</li>
                    <ul>
                        <li>ゼミ長に理由と、どのくらい遅れるかを<b>必ず</b>連絡</li>
                    </ul>
                </ul>
            </div>
        </div>
    )
}

export default Home