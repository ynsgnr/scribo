class ScriboHeader extends HTMLElement {
        constructor() {
            super();
            let template = document.createElement("template")
            template.innerHTML = `
            <div style = "width:100%; height:10%; padding:3%; text-align: left; color:white; background-color:#f28c1f; ">
                <span>Scribo</span>
            </div>
            `
            let shadowRoot = this.attachShadow({ mode: "open" });
            shadowRoot.appendChild(template.content.cloneNode(true));
        }
    }
window.customElements.define("scribo-header", ScriboHeader);