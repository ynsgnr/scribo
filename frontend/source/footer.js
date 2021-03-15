class ScriboFooter extends HTMLElement {
        constructor() {
            super();
            let template = document.createElement("template")
            template.innerHTML = `
            <div style = "width:100%; height:50px; border-top-style: solid; text-align: center; color:black; border-color:#f28c1f; border-width:1px 0px 0px 0px; margin-top:5%;">
                <div style= "padding:3px 0px 3px 0px; justify-content: center;  align-items: center; display: flex;height: 100%;">
                    <span>Scribo</span>
                </div>
            </div>
            `
            let shadowRoot = this.attachShadow({ mode: "open" });
            shadowRoot.appendChild(template.content);
        }
    }
window.customElements.define("scribo-footer", ScriboFooter);