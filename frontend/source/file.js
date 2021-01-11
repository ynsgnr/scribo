
import {upload} from './api.file.upload.js'

class FileUploader extends HTMLElement {
    constructor() {
        super();
        let template = document.createElement("template")
        template.innerHTML = `
        <input id="fileUpload" type="file" accept=".azw,.azw3,.azw4,.cbz,.cbr,.cbc,.chm,.djvu,.docx,.epub,.fb2,.fbz,.html,.htmlz,.lit,.lrf,.mobi,.odt,.pdf,.prc,.pdb,.pml,.rb,.rtf,.snb,.tcr,.txt,.txtz">
        `
        let shadowRoot = this.attachShadow({ mode: "open" });
        shadowRoot.appendChild(template.content);
        this.fileUploadInput = shadowRoot.getElementById("fileUpload")
        this.fileUploadInput.addEventListener("change",()=>{this.fileUpload(this.fileUploadInput.files)})
    }

    fileUpload(files){
        if (!files.length) {
            return
        }
        upload(files[0]);
    }
}
window.customElements.define("file-uploader", FileUploader);