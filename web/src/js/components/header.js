import { createElementObj } from "../utils/create";
import { tagsOptions } from "../utils/options";
import Container from "./container";
import Sandwich from "./sandwich";
import NavItems from "./nav-items";

const Header = () => {
    const headerTag = document.getElementById("header")
    const logo = createElementObj({ tagName: 'a', classNames: 'navbar-brand', attrs: [['href', '#']], textContent: 'Forum' })
    const sandwich = Sandwich().get()
    const items = NavItems("Home", "Posts", "Dialogs", "Auth").get()

    const container = Container(logo, sandwich, items).get()
    const nav = createElementObj({ ...tagsOptions.nav, children: [container] })

    return {
        init: () => {
            headerTag.append(nav)
        }
    }
}

export default Header;