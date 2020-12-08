import './auth.js'
import './footer.js'
import './device.js'
import './header.js'
import './file.js'

import * as cookie from './cookie.js'

class App extends HTMLElement {
        constructor() {
            super();
            this.root = this.attachShadow({ mode: "open" });

            this.appTemplate =  document.createElement("template")
            this.appTemplate.innerHTML = `
            <div style="min-height: 100%; width:100%;">
                <scribo-header></scribo-header>
                <div style = "display:none; margin: 0 auto; text-align: center; overflow:auto; height:100%;" id="login">
                    <div style = "float: right; width:30%; height:100%; padding: 1%; overflow:auto;">
                        <scribo-auth id="scribo-auth"></scribo-auth>
                    </div>
                </div>
                <div style = "display:none; margin: 0 auto; text-align: center; overflow:auto; height:100%;" id="app">
                    <scribo-device></scribo-device>
                    <file-uploader></file-uploader>
                </div>
                <scribo-footer></scribo-footer>
            </div>
            `
            this.root.appendChild(this.appTemplate.content.cloneNode(true))
            this.authElem =  this.root.getElementById("scribo-auth")
            this.login =  this.root.getElementById("login")
            this.app = this.root.getElementById("app")
            
            this.authElem.addEventListener("signedin",()=>{this.app.style.removeProperty("display");this.login.style.display="none";})
            this.authElem.addEventListener("authrequired",()=>{this.login.style.removeProperty("display")})
        }
    }
window.customElements.define("app-main", App);