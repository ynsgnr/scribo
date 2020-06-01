class ScriboAuth extends HTMLElement {
    constructor() {
        super();
        let template = document.createElement("template")
        template.innerHTML = `
        <div>
            <p>Custom element shadow tree content... Auth</p>
        </div>
        `
        let shadowRoot = this.attachShadow({ mode: "open" });
        shadowRoot.appendChild(template.content.cloneNode(true));
    }
}
window.customElements.define("scribo-auth", ScriboAuth);