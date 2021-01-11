class ScriboHeader extends HTMLElement {
        constructor() {
            super();
            let template = document.createElement("template")
            template.innerHTML = `
            <div style = "width:100%; height:50px; padding:3px; justify-content: center; text-align: left; color:white; background-color:#f28c1f; ">
                <div style= "padding:3px 0px 3px 10px; align-items: center; display: flex;height: 100%;">
                    <span>Scribo</span>
                </div>
            </div>
            `
            let shadowRoot = this.attachShadow({ mode: "open" });
            shadowRoot.appendChild(template.content);
        }
    }
window.customElements.define("scribo-header", ScriboHeader);