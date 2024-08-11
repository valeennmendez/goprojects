console.log("Conectado...")

document.addEventListener("DOMContentLoaded", function(e){
    const loginForm = document.getElementById("formLogin")

    loginForm.addEventListener("submit",function(e){
        e.preventDefault()

        const email = document.getElementById("email").value
        const password = document.getElementById("password").value
        const error = document.getElementById("error")

        if(email === "" || password === ""){
            error.innerHTML = "Please fill in the fields"
            error.classList.add("active")
            return
        }

        console.log(email, password)

        fetch("http://localhost:8080/login",{
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                email: email,
                password: password,
            }),
            credentials: "include",
        })
        .then(response => response.json())
        .then(data =>{
            console.log(data.message)
            if(data.message === 'Login succesful'){
                window.location.href = "index.html"
                return
            }else{
                error.classList.add("active")
                return
            }
        })
        //.catch(error => console.error(error))
    })
})