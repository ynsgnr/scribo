export function Popup(shadowRoot,element){
    let popup = document.createElement("basic-pop-up")
    popup.appendChild(element)
    let template = document.createElement("template")
    template.content.appendChild(popup)
    shadowRoot.appendChild(template.content)
    return popup
}

class BasicPopUp extends HTMLElement {
    constructor() {
        super();
        let template = document.createElement("template")
        template.innerHTML = `
        <div style="
            padding: 2%;
            background-clip: padding-box;
            background-color: #000000;
            z-index: 1050;
            text-align:left;
            position: fixed;
            top: 50%;
            left: 50%;
            transform: translate(-50%, -50%);
            ">
                <slot>POP UP CONTEXT</slot>
        </div>
        <div id="background-dark-filter" style="
            background-clip: padding-box;
            background-color: #000000AA;
            position: fixed; top: 0; left: 0; right: 0; bottom: 0;
            width: 100%;
            height: 100%;
            z-index: 1000;
        "></div>
        `
        let root = template.content
        this.bgFilter = root.getElementById("background-dark-filter")
        this.bgFilter.onclick = ()=>this.remove()
        let shadowRoot = this.attachShadow({ mode: "open" });
        shadowRoot.appendChild(root);
    }
}
window.customElements.define("basic-pop-up", BasicPopUp);