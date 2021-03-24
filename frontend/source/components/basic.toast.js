export function Toast(shadowRoot,element){
    let toast = document.createElement("basic-toast")
    let realElement
    if (!element.nodeName){
        realElement = document.createElement("p")
        realElement.style = "margin-top: 0px; margin-bottom: 0px;"
        realElement.innerHTML = element
    }else{
        realElement = element
    }
    toast.appendChild(realElement)
    let template = document.createElement("template")
    template.content.appendChild(toast)
    shadowRoot.appendChild(template.content)
    return toast
}

class BasicToast extends HTMLElement {
    constructor() {
        super();
        let template = document.createElement("template")
        template.innerHTML = `
        <style>
            .toast {
                padding: 1%;
                background-clip: padding-box;
                background-color: #000000;
                z-index: 1100;
                text-align:left;
                position: fixed;
                top: 10%;
                left: 50%;
                transform: translate(-50%, -50%);
            }
            .visible {
                visibility: visible;
                opacity: 1;
            }
            .hidden {
                visibility: hidden;
                opacity: 0;
            }
            .animation{
                transition: 0.25s all ease-in-out;
            }
        </style>
        <div class="animation toast hidden" id="wrapper">
                <slot>TOAST CONTEXT</slot>
        </div>
        `
        let root = template.content
        this.toastWrapper = root.getElementById("wrapper")
        setTimeout(()=>{
            this.toastWrapper.classList.add('hidden')
            setTimeout(()=>{
                this.remove()
            }, 500)//should be same with transition time
        }, 3000)
        let shadowRoot = this.attachShadow({ mode: "open" });
        shadowRoot.appendChild(root);
        setTimeout(()=>{
            //Fade in
            this.toastWrapper.classList.remove('hidden')
            this.toastWrapper.classList.add('visible')
        }, 25)
    }
}
window.customElements.define("basic-toast", BasicToast);