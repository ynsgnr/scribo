import './auth.js'
import './footer.js'
import './device.js'
import './header.js'
import './file.js'

import * as cookie from './cookie.js'
import { UserIDKey, UserKey } from './auth.js'

class App extends HTMLElement {
        constructor() {
            super();
            this.root = this.attachShadow({ mode: "open" });

            this.appTemplate = document.createElement("template")
            this.appTemplate.innerHTML = `
            <scribo-header id="header"></scribo-header>
            <div style="min-height: 100%; width:100%;">
                <div style = "display:none; margin: 0 auto; text-align: center; overflow:auto; height:100%;" id="login">
                    <div style = "float: right; width:30%; height:100%; padding: 1%; overflow:auto;">
                        <scribo-auth id="scribo-auth"></scribo-auth>
                    </div>
                </div>
                <div style = "display:none; margin: 0 auto; text-align: center; overflow:auto; height:100%;" id="app">
                </div>
            </div>
            <scribo-footer></scribo-footer>
            `

            this.authElem =  this.appTemplate.content.getElementById("scribo-auth")
            this.login =  this.appTemplate.content.getElementById("login")
            this.header = this.appTemplate.content.getElementById("header")
            this.app = this.appTemplate.content.getElementById("app")

            this.appContent = `<scribo-device></scribo-device>`

            this.header.addEventListener("signout",()=>{
                this.authElem.signOut(cookie.getCookie(UserKey))
                .then(()=>{
                    this.app.style.display = "none";this.login.style.removeProperty("display")})})

            this.authElem.addEventListener("signedin",(e)=>{
                this.app.innerHTML = this.appContent
                this.header.dispatchEvent(new e.constructor(e.type, e));
                this.app.style.removeProperty("display");
                this.login.style.display="none";
            })
            this.authElem.addEventListener("authrequired",()=>{
                this.app.innerHTML = ""
                this.app.style.display = "none";
                this.login.style.removeProperty("display")
            })

            this.root.appendChild(this.appTemplate.content)
        }
    }
window.customElements.define("app-main", App);