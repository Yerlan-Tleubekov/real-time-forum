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

const liCreator = (names) => {
    return names.map((name) => {
        const a = createElementObj({
            tagName: 'a',
            classNames: 'nav-link',
            textContent: name,
        })

        return createElementObj({
            tagName: 'li',
            classNames: 'nav-item',
            children: [a]
        })
    })

}

const NavItems = (...names) => {



    const ul = createElementObj({
        tagName: 'ul',
        classNames: 'navbar-nav me-auto mb-2 mb-lg-0',
        children: []
    })

    const wrap = createElementObj({
        tagName: 'div',
        classNames: 'collapse navbar-collapse',
        attrs: [['id', 'navbarSupportedContent']],
        children: liCreator(names)
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