import './scriboAuth.js'
import './scriboFooter.js'
import './scriboHeader.js'

class App extends HTMLElement {
        constructor() {
            super();
            let template = document.createElement("template")
            template.innerHTML = `
            <div style = "min-height: 100%; width:100%;">
                <scribo-header></scribo-header>
                <div style = "margin: 0 auto; text-align: center; overflow:auto; min-height: 100%;">
                    <div style = "float: right; width:30%; padding: 1%; overflow:auto; ">
                        <scribo-auth></scribo-auth>
                    </div>
                </div>
                <scribo-footer></scribo-footer>
            </div>
            `
            let shadowRoot = this.attachShadow({ mode: "open" });
            shadowRoot.appendChild(template.content.cloneNode(true));

        }
    }
window.customElements.define("app-main", App);