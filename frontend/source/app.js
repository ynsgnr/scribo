class App extends HTMLElement {
    constructor(){
        super();

        const shadowRoot = this.attachShadow({mode: 'open'});
        shadowRoot.innerHTML = '<h1>Hello Shadow DOM</h1>';
    }
}

window.customElements.define('app-main', App);