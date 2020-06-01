class ScriboFooter extends HTMLElement {
        constructor() {
            super();
            let template = document.createElement("template")
            template.innerHTML = `
            <div style = "width:100%; height:10%; padding:3% 0% 3% 0%; border-top-style: solid; text-align: center; color:black; border-color:#f28c1f; border-width:1px 0px 0px 0px; position: absolute; bottom: 0;">
                <span>Scribo</span>
            </div>
            `
            let shadowRoot = this.attachShadow({ mode: "open" });
            shadowRoot.appendChild(template.content.cloneNode(true));
        }
    }
window.customElements.define("scribo-footer", ScriboFooter);