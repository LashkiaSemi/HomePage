import React from 'react'
import { Link } from 'react-router-dom'

// TODO: navigationのリンク
// TODO: ログイン状態の保存?

// Header
// 頭のナビゲーションバーです
class Header extends React.Component {
    constructor(props) {
        super(props)
    }
    
    render() {
        return (
            <div className="header">
                <Navigation />
            </div>
        )
    }
}

// Navigation
// ナビゲーションバー
const Navigation = (props) => (
    <nav>
        <Logo />
        <Menu />
    </nav>
)

// Logo
// らしきあぜみって書いておく場所
const Logo = (props) => (
    <div className="logo">
        {/* TODO: ここLinkにして */}
        <Link to="/">Lashkia研究室</Link>
    </div>
)

// Menu
// TODO: ここ、あれ、storeから読めばいいのでは？
const Menu = (props) => {
    return (
        <ul className="menu">
            {
                MENU.map((menu) => (
                    <MenuRow key={menu.id} menu={menu} />
                ))
            }
        </ul>
    )
}

// MenuRow
// 項目一個
const MenuRow = (props) => {
    // TODO: 必要なら、何かしら、分岐のしょりをば
    return (
        <li><Link to={`/${props.menu.id}`}>{props.menu.display}</Link></li>
    )
}

// Static datas
const MENU = [
    {
        id: "activities",
        display: "活動記録",
    },
    {
        id: "societies",
        display: "学会発表",
    },
    {
        id: "researches",
        display: "卒業研究",
    },
    {
        id: "jobs",
        display: "就職先",
    },
    {
        id: "members",
        display: "メンバー",
    },
    {
        id: "links",
        display: "外部リンク",
    },
    {
        id: "equipments",
        display: "研究室備品",
        loginRequired: true,
    },
    {
        id: "lectures",
        display: "レクチャー",
        loginRequired: true,
    },
    {
        id: "login",
        display: "ログイン",
        loginRequired: false,
    },
    {
        id: "logout",
        display: "ログアウト",
        loginRequired: true,
    },
]

export default Header

