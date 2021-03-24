class RecycleList extends HTMLElement {

    constructor() {
        super();
        let template = document.createElement("template")
        template.innerHTML = `
        <div id="list">
            <slot></slot>
        </div>
        `
        this._elements = [];

        let root = template.content
        this._list = root.getElementById("list")
        let shadowRoot = this.attachShadow({ mode: "open" })
        shadowRoot.appendChild(root)
    }


    /**
     * @param {HTMLElement} data
     */
    set base(baseElement) {
        this._base = baseElement
    }


    /**
     * @param {any[]} data
     */
    set items(data) {
        let objectData
        if(!Array.isArray(data)){
            if (typeof data === 'object'){
                objectData = data
                data = Object.keys(objectData)
            }else{
                throw "Unexpected input"
            }
        }
        for (var i=0;i<data.length;i++){
            var obj = data[i]
            if (typeof data[i] !== 'object'){
                //assume its the key
                obj = objectData[data[i]]
                if (!obj){
                    throw "Unexpected input: missing key "+data[i]
                }
            }
            if (i>=this._elements.length){
                let newElement = this.base.cloneNode(true)
                this._elements.push(newElement)
                this._list.appendChild(newElement)
            }
            let keys = Object.keys(obj)
            for (var j=0;j<keys.length;j++){
                if (typeof obj[keys[j]]==='number' || typeof obj[keys[j]]==='string'){
                    this._elements[i].setAttribute(keys[j],obj[keys[j]])
                }
            }
            this._elements[i].data = obj
            this._elements[i].onclick = (event)=>{this.elementOnclick(event.target)}
        }
        if (i<this._elements.length){
            for (var k = i; k<this._elements.length; k++){
                this._list.removeChild(this._elements[k])
            }
            this._elements = this._elements.slice(0,i)
        }
    }

    get selectedItem(){
        return this.selectedElement
    }
    
    elementOnclick(element){
        if (element && element != this.selectedElement){
            if(this.selectedElement){
                this.selectedElement.removeAttribute("selected")
            }
            element.setAttribute("selected","true")
            this.selectedElement = element
            
            this.dispatchEvent(new CustomEvent("itemSelect",{composed: true, detail:{selected:element}}))
        }
    }
}
window.customElements.define("recycle-list", RecycleList);