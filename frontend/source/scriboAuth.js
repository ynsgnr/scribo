class ScriboAuth extends HTMLElement {
    constructor() {
        super();
        let template = document.createElement("template")
        template.innerHTML = `
        <div style = "width:100%; height:100%; flex:1;">
            <form>
                <label for="email">Email:</label><br>
                <input type="text" id="ScriboAuthEmail" name="email" value="email"><br>
                <label for="pass">Password: </label><br>
                <input type="password" id="ScriboAuthPassword" name="pass" value="pass"><br>
                <input type="button" id="ScriboAuthLoginButton" value="Log In">
            </form>
        </div>
        `
        this.root = this.attachShadow({ mode: "open" });
        this.root.appendChild(template.content.cloneNode(true));
        this.emailInput = this.root.getElementById("ScriboAuthEmail")
        this.passInput = this.root.getElementById("ScriboAuthPassword")
        this.root.getElementById("ScriboAuthLoginButton").addEventListener("click",()=>this.login(this.emailInput.value,this.passInput.value))
    }

    login(username, password){
        console.log(username)
        console.log(password)
    }
}
window.customElements.define("scribo-auth", ScriboAuth);