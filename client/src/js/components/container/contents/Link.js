import React from 'react'

// Linkは基本的に静的。
// 内容の編集はLABとかTECHとかを編集してください

// LAB 他研究室へのリンク
const LAB = [
    {
        name: '鈴木 研究室',
        path: 'http://www.suzuki.sist.chukyo-u.ac.jp/',
    },
    {
        name: '濱川 研究室',
        path: 'http://hamakawalab.sist.chukyo-u.ac.jp/',
    },
    {
        name: 'MDLAB（目加田・道満 研究室）',
        path: 'https://md.sist.chukyo-u.ac.jp/index.html',
    },
    {
        name: '鬼頭 研究室',
        path: 'http://kitolab.sist.chukyo-u.ac.jp/'
    },
    {
        name: 'オープンメディアラボ（宮崎・山田・中 研究室）',
        path: 'https://www.om.sist.chukyo-u.ac.jp/',
    }
]

// TECH 技術系のリンク
const TECH = [
    {
        name: 'Flutter',
        comment: 'Dart言語を用いた Android／iOSアプリ開発のフレームワーク',
        path: 'https://flutter.dev/',
    },
    {
        name: 'Docker',
        comment: '軽量なコンテナ型の仮想環境を提供するオープンソースソフトウェア',
        path: 'https://www.docker.com/',
    },
    {
        name: 'Electron',
        comment: 'クロスプラットホームなデスクトップアプリを開発できるフレームワーク',
        path: 'https://www.electronjs.org/',
    },
    {
        name: 'ARCore',
        comment: 'Googleが提供するARフレームワーク',
        path: 'https://developers.google.com/ar/',
    },
]


/*
Link リンク。めっちゃ静的
*/
class Link extends React.Component {
    render() {
        return (
            <div className="content">
                <h1 className="content-title h1-block">外部リンク</h1>
                <LinkList title={"他研究室へのリンク"} links={LAB} />
                <LinkList title={"面白いソフト・最新技術紹介"} links={TECH} />
            </div>
        )
    }
}

/*
LinkList リンクをリストで表示
props:
    title = リストのタイトル
    links = リンクのデータセット
        [{
            name: 見出し
            path: リンク先
            comment: リンクに対して備考があれば
        }]
*/
const LinkList = (props) => {
    return (
        <div className="list">
            <h3 className="list-title h3">{props.title}</h3>
            <ul>
                {
                    props.links.map((link) => (
                        <LinkRow key={link.path} link={link} />
                    ))
                }
            </ul>
        </div>
    )
}

/*
LinkRow リンク一件
props:
    link: 親からlink一件
*/
const LinkRow = (props) => {
    return (
        <li>
            <a href={props.link.path} className="list-item" target="_blank">
                <div>{props.link.name}</div>
                <div className="col-black">{props.link.comment}</div>
            </a>
        </li>
    )
}

export default Link