import {GetDevices} from './api.js'

class ScriboDevice extends HTMLElement {
    constructor() {
        super();
        let template = document.createElement("template")
        template.innerHTML = `
        <div style="min-height: 100%; width:100%;">
            <div style = "margin: 0 auto; text-align: center; overflow:auto; height:100%;" id="loading-display">
                <p>Loading...<p/>
            </div>
            <div style = "display:none;" id="content">
                <div style = "display:flex; margin: 0 auto; text-align: center; overflow:auto; height:100%;">
                    <div>
                        <p>Content Here</p>
                    </div>
                    <div>
                        <p>Content Here</p>
                    </div>
                </div>
            </div>
        </div>
        `
        this.root = template.content
        this.loading = this.root.getElementById("loading-display")
        this.content = this.root.getElementById("content")
        let shadowRoot = this.attachShadow({ mode: "open" });
        shadowRoot.appendChild(this.root);
    }

    loaded(){
        this.content.style.removeProperty("display")
        this.loading.style.display="none"
    }

    connectedCallback(){
        GetDevices().then((result)=>{this.loaded();console.log(result)})
    }
}
window.customElements.define("scribo-device", ScriboDevice);