import { createElementObj } from "../utils/create";
import { Form, Input } from "../components/form-control";

{/* <div class="modal" tabindex="-1">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title">Modal title</h5>
                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
                <p>Modal body text goes here.</p>
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
                <button type="button" class="btn btn-primary">Save changes</button>
            </div>
        </div>
    </div>
</div> */}

function OnSubmitSignInForm(e) {
    console.log(e.target)


    return false
}

function SignInBody() {

    const login = Input().get()
    const password = Input().get()

    const signInForm = Form(OnSubmitSignInForm, login, password).get()

    const b = createElementObj({
        tagName: 'div',
        classNames: 'modal-body',
        children: [signInForm]
    })




    return {
        get: () => {
            return b
        }
    }
}


function AuthModal(titleName) {
    const signInBody = SignInBody().get()
    console.log(signInBody)

    const signUp = createElementObj({
        tagName: 'button',
        classNames: 'button',
        textContent: "Sign Up",
    })

    const signIn = createElementObj({
        tagName: 'button',
        classNames: 'button',
        textContent: "Sign In",
    })

    const title = createElementObj({
        tagName: 'div',
        classNames: 'modal-title',
        textContent: titleName,
    })

    const header = createElementObj({
        tagName: 'div',
        classNames: 'modal-header',
        children: [title, signIn, signUp]
    })

    const content = createElementObj({
        tagName: 'div',
        classNames: 'modal-content',
        children: [header, signInBody]
    })

    const dialog = createElementObj({
        tagName: 'div',
        classNames: 'modal-dialog',
        children: [content]
    })


    const modal = createElementObj({
        tagName: "div",
        classNames: "modal fade",
        attrs: [['tabindex', '-1'], ['id', 'exampleModal'], ['aria-labelledby', 'exampleModalLabel'], ['aria-hidden', 'true']],
        children: [dialog]
    })


    const layout = document.getElementById("body")


    return {
        get: () => {

            return modal
        },
        init: () => {
            layout.append(modal)
        }
    }

}


export default AuthModal
