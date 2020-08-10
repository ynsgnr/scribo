import './auth.js'
import './footer.js'
import './header.js'

class App extends HTMLElement {
        constructor() {
            super();
            this.root = this.attachShadow({ mode: "open" });

            this.loginTemplate =  document.createElement("template")
            this.loginTemplate.innerHTML = `
            <div id="scribo-login" style="min-height: 100%; width:100%;">
                <scribo-header></scribo-header>
                <div style = "margin: 0 auto; text-align: center; overflow:auto; height:100%; ">
                    <div style = "float: right; width:30%; height:100%; padding: 1%; overflow:auto;">
                        <scribo-auth id="scribo-auth"></scribo-auth>
                    </div>
                </div>
                <scribo-footer></scribo-footer>
            </div>
            `
            this.root.appendChild(this.loginTemplate.content.cloneNode(true));
            this.loginPage = this.root.getElementById("scribo-login")
            this.authElem = this.root.getElementById("scribo-auth")
            this.authElem.addEventListener("signedin",()=>{this.signedIn()})

            this.appTemplate = document.createElement("template")
            this.appTemplate.innerHTML = `
            <div id="scribo-app" style = "min-height: 100%; width:100%;">
                <scribo-header></scribo-header>
                <div style = "margin: 0 auto; text-align: center; overflow:auto; height:100%; ">
                    <span>App</span>
                </div>
                <scribo-footer></scribo-footer>
            </div>
            `
            this.appPage = this.appTemplate.content.cloneNode(true)
        }

        signedIn(){
            this.root.removeChild(this.loginPage)
            this.root.appendChild(this.appPage)
            this.appPage = this.root.getElementById("scribo-app")
        }
    }
window.customElements.define("app-main", App);