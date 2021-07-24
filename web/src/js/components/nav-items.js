import { createElementObj } from "../utils/create"

{/* <div class="collapse navbar-collapse" id="navbarSupportedContent">
      <ul class="navbar-nav me-auto mb-2 mb-lg-0">
        <li class="nav-item">
          <a class="nav-link active" aria-current="page" href="#">Home</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="#">Link</a>
        </li>
      </ul>
      <form class="d-flex">
        <input class="form-control me-2" type="search" placeholder="Search" aria-label="Search">
        <button class="btn btn-outline-success" type="submit">Search</button>
      </form>
    </div> */}

const liCreator = (navItems) => {
    return navItems.map((navItem) => {
        const a = createElementObj({
            tagName: 'a',
            classNames: navItem.classNames,
            textContent: navItem.name,
            attrs: navItem.attrs
        })

        return createElementObj({
            tagName: 'li',
            classNames: 'nav-item',
            children: [a]
        })
    })

}

const NavItems = (navitems) => {



    const ul = createElementObj({
        tagName: 'ul',
        classNames: 'navbar-nav me-auto mb-2 mb-lg-0',
        children: []
    })

    const wrap = createElementObj({
        tagName: 'div',
        classNames: 'collapse navbar-collapse',
        attrs: [['id', 'navbarSupportedContent']],
        children: liCreator(navitems)
    })

    return {
        init: () => {

        },
        get: () => {
            return wrap
        }
    }
}

export default NavItems