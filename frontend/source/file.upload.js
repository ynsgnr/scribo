class FileUpload extends HTMLElement {

    static get observedAttributes() { return ['disabletop','disablebottom','disablehint']; }

    constructor() {
        super();
        let template = document.createElement("template")
        template.innerHTML = `
        <div id="drop-zone">
            <slot id="top-slot">TOP</slot>
            <input type="file" id="upload" hidden/>
            <label id="upload-hint" for="upload" style="color:gray;">
                Drag and drop file or click here to send new file
            </label>
            <slot id="bottom-slot">BOTTOM</slot>
        </div>
        `
        let root = template.content

        this.manualUpload = root.getElementById("upload")
        this.manualUpload.oninput = ()=>this.fileUpload(this.manualUpload.files)
        this.manualUploadHint = root.getElementById("upload-hint")

        this.dropZone = root.getElementById("drop-zone")
        this.dropZone.addEventListener("dragenter",(e)=>{
            e.preventDefault()
            e.stopPropagation()
            this.dropZone.style.background = "#cecece"
        });
        this.dropZone.addEventListener("dragleave", (e)=>{
            e.preventDefault()
            e.stopPropagation()
            this.dropZone.style.removeProperty("background")
        });
        this.dropZone.addEventListener("dragover", (e)=>{
            e.preventDefault()
            e.stopPropagation()
        });
        this.dropZone.addEventListener("drop", (e)=>{
            e.preventDefault()
            e.stopPropagation()
            this.dropZone.style.removeProperty("background")
            console.log(e.dataTransfer)
            this.fileUpload(e.dataTransfer.files)
        });

        
        this.topSlot = root.getElementById("top-slot")
        this.bottomSlot = root.getElementById("bottom-slot")

        let shadowRoot = this.attachShadow({ mode: "open" })
        shadowRoot.appendChild(root)
    }

    fileUpload(files){
        this.dispatchEvent(new CustomEvent("filedrop",{composed: true, detail:files}))
        this.manualUpload.value = null
    }

    attributeChangedCallback(name, oldValue, newValue) {
        switch (name){
            case "disabletop":
                if (newValue || newValue=="true"){
                    this.topSlot.style.display = "none"
                }else{
                    this.topSlot.style.removeProperty("display")
                }
                break
            case "disablebottom":
                if (newValue || newValue=="true"){
                    this.bottomSlot.style.display = "none"
                }else{
                    this.bottomSlot.style.removeProperty("display")
                }
                break
            case "disablehint":
                if (newValue || newValue=="true"){
                    this.manualUploadHint.style.display = "none"
                }else{
                    this.manualUploadHint.style.removeProperty("display")
                }
                break

        }
    }
}
window.customElements.define("file-upload", FileUpload);