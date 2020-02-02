import React from 'react'

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
        <a href="/">Lashkia研究室</a>
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
        // TODO: link
        <li><a>{props.menu.display}</a></li>
    )
}

// Static datas
const MENU = [
    {
        id: "activity",
        display: "活動記録",
    },
    {
        id: "society",
        display: "学会発表",
    },
    {
        id: "research",
        display: "卒業研究",
    },
    {
        id: "job",
        display: "就職先",
    },
    {
        id: "member",
        display: "メンバー",
    },
    {
        id: "link",
        display: "外部リンク",
    },
    {
        id: "equipment",
        display: "研究室備品",
        loginRequired: true,
    },
    {
        id: "lecture",
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

