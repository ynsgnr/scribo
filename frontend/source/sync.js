class ScriboSync extends HTMLElement {
    constructor() {
        super();
        let template = document.createElement("template")
        template.innerHTML = `
        <ul>
            <li>Coffee</li>
            <li>Tea</li>
            <li>Milk</li>
        </ul> 
        `
        let shadowRoot = this.attachShadow({ mode: "open" });
        shadowRoot.appendChild(template.content);
    }
}
window.customElements.define("scribo-sync", ScriboSync);