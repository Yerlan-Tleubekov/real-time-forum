import { createElementObj } from "../utils/create";
import { tagsOptions } from "../utils/options";
import Container from "./container";
import Sandwich from "./sandwich";
import NavItems from "./nav-items";

{/* <button type="button" class="btn btn-primary" data-bs-toggle="modal" data-bs-target="#exampleModal">
  Launch demo modal
</button> */}

const navItems = [
    {
        name: 'Home',
        href: '#',
        classNames: 'nav-link',
        attrs: []
    },
    {
        name: 'Posts',
        href: '#',
        classNames: 'nav-link',
        attrs: []
    },
    {
        name: 'Dialogs',
        href: '#',
        classNames: 'nav-link',
        attrs: []
    },
    {
        name: 'Auth',
        href: '#',
        classNames: 'nav-link',
        attrs: [['data-bs-toggle',"modal"], ['data-bs-target', "#exampleModal"]]
    },
]

const Header = () => {
    const headerTag = document.getElementById("header")
    const logo = createElementObj({ tagName: 'a', classNames: 'navbar-brand', attrs: [['href', '#']], textContent: 'Forum' })
    const sandwich = Sandwich().get()
    const items = NavItems(navItems).get()

    const container = Container(logo, sandwich, items).get()
    const nav = createElementObj({ ...tagsOptions.nav, children: [container] })

    return {
        init: () => {
            headerTag.append(nav)
        }
    }
}

export default Header;