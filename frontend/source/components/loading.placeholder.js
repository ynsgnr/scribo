class LoadingPlaceHolder extends HTMLElement {

    static get observedAttributes() { return ['loaded']; }

    constructor() {
        super();
        let template = document.createElement("template")
        template.innerHTML = `
        <div style = "margin: 0 auto; text-align: center; overflow:auto; height:100%;" id="loading-display">
            <slot name="placeholder">Loading...<slot/>
        </div>
        <slot name="content" style="display:none" id="content"><slot/>
        `
        this._elements = [];

        let root = template.content
        this._loadingDisplays = root.getElementById("loading-display")
        this._content = root.getElementById("content")
        let shadowRoot = this.attachShadow({ mode: "open" })
        shadowRoot.appendChild(root)
    }

    update(){
        if (this.getAttribute("loaded") == "true"){
            this._content.style.display = null
            this._loadingDisplays.style.display="none"
        }else{
            this._content.style.display="none"
            this._loadingDisplays.style.display = null
        }
    }

    attributeChangedCallback(name, oldValue, newValue) {
        this.update();
    }
    connectedCallback(){
        this.update()
    }
    adoptedCallback() {
        this.update()
    }

}
window.customElements.define("loading-place-holder", LoadingPlaceHolder);