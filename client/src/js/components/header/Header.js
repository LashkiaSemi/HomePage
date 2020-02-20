import React from 'react'
import { Link } from 'react-router-dom'
import { STRAGE_KEY } from '../../constants/config'
/*
static data
MENU ナビゲーションの項目
obj: {
    id: 押した時にジャンプするpathになってる
    display: 表示名
    isLogin: bool。ログイン時に表示するか
    isNotLogin: bool。未ログイン時に表示するか
}
*/
const MENU = [
    {
        id: "activities",
        display: "活動記録",
        isLogin: true,
        isNotLogin: true,
    },
    {
        id: "societies",
        display: "学会発表",
        isLogin: true,
        isNotLogin: true,
    },
    {
        id: "researches",
        display: "卒業研究",
        isLogin: true,
        isNotLogin: true,
    },
    {
        id: "jobs",
        display: "就職先",
        isLogin: true,
        isNotLogin: true,
    },
    {
        id: "members",
        display: "メンバー",
        isLogin: true,
        isNotLogin: true,
    },
    {
        id: "links",
        display: "外部リンク",
        isLogin: true,
        isNotLogin: true,
    },
    {
        id: "equipments",
        display: "研究室備品",
        isLogin: true,
        isNotLogin: false,
    },
    {
        id: "lectures",
        display: "レクチャー",
        isLogin: true,
        isNotLogin: false,
    },
    {
        id: "login",
        display: "ログイン",
        isLogin: false,
        isNotLogin: true,
    },
    {
        id: "logout",
        display: "ログアウト",
        isLogin: true,
        isNotLogin: false,
    },
]


/*
Header ヘッダー。主にナビゲーションバーだけどな！
*/
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

/*
Navigation ナビゲーションバー
*/
const Navigation = () => (
    <nav>
        <Logo />
        <Menu />
    </nav>
)

/*
Logo ナビゲーションバーの一番左のロゴ。Lashkiaゼミって書いておくとこ
*/
const Logo = () => (
    <div className="logo">
        <Link to="/">Lashkia研究室</Link>
    </div>
)

/*
Menu ナビゲーションバーのメニュー部分
メニューの内容は MENU にて。
*/
const Menu = () => {
    // localStrageに値がなかったらnullになります
    const isLogin = localStorage.getItem(STRAGE_KEY)
    return (
        <ul className="menu">
            {
                MENU.map((menu) => (
                    isLogin
                    ? menu.isLogin
                        ? <MenuRow key={menu.id} menu={menu} />
                        : <></>
                    : menu.isNotLogin
                        ? <MenuRow key={menu.id} menu={menu} />
                        : <></>
                ))
            }
        </ul>
    )
}

/*
MenuRow メニューの項目1つ
props:
    menu = 1項目
*/
const MenuRow = (props) => {
    return (
        <li><Link to={`/${props.menu.id}`}>{props.menu.display}</Link></li>
    )
}

export default Header

