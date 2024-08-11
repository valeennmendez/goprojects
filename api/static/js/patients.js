console.log("Conectado...")

const urlApi = "http://localhost:8080/patients"


function ValidateSession(){
    fetch("http://localhost:8080/validate",{
        method: "GET",
        credentials: "include"
    })
    .then(response =>{
        if (!response.ok){
            window.location.href = "login.html"
        }
    })
    .catch(error => console.error(error))
}

function DeletePatients(){
    const tabla = document.querySelector(".tabla")

    tabla.addEventListener("click", function(e){
        const deleteButton = e.target.closest(".deleteBtn")

        if(deleteButton){
            const idpatients = deleteButton.getAttribute("data-id")
            const deleteUrl = `${urlApi}/${idpatients}`
        
            if(confirm("¿Seguro que desea eliminar al paciente?")){
                fetch(deleteUrl,{
                    method: "DELETE",
                    credentials: "include"
                })
                    .then(response =>{
                        if(!response.ok){
                            console.error("error:",response)
                            return
                        }
                        const row = e.target.closest("tr")
                        if(row){
                            row.remove()
                        }
                    })
                .catch(error => console.error("error:", error))
            }  
        }
    })
}

function loadPatientDataApp(id){
    fetch(`http://localhost:8080/patients/${id}`,{
        credentials: "include",
    })
    .then(response =>{
        if(!response.ok){
            console.error(response)
        }
        return response.json()
    })
    .then(data => {
        document.getElementById("emailApp").value = data.Email
        document.getElementById("dniApp").value = data.Dni
        document.getElementById("formContainerAppoinment").setAttribute("data-id", data.ID)
    })
}

function loadPatientData(id){
    fetch(`http://localhost:8080/patients/${id}`,{
        credentials: "include",
    })
    .then(response =>{
        if(!response.ok){
            console.error(response)
        }
        return response.json()
    })
    .then(data => {
        document.getElementById("nameEdit").value = data.FullName
        document.getElementById("emailEdit").value = data.Email
        document.getElementById("dniEdit").value = data.Dni
        document.getElementById("phoneEdit").value = data.Phone
        document.getElementById("formContainerEdit").setAttribute("data-id", data.ID)
    })
}

function EditPatients(){
    const formContainerEdit = document.getElementById("formContainerEdit")
    
    if(formContainerEdit){

        const closeFormEdit = document.getElementById("closeFormEditBtn")

        closeFormEdit.addEventListener("click", function(){
            formContainerEdit.classList.remove("active")
        })


        formContainerEdit.addEventListener("submit", function(){

            const idpatient = formContainerEdit.getAttribute("data-id")
            console.log("PACIENTE A EDITAR: ",idpatient)

            const fullnameP = document.getElementById("nameEdit").value
            const emailP = document.getElementById("emailEdit").value
            const dniP = document.getElementById("dniEdit").value
            const phoneP = document.getElementById("phoneEdit").value

            const patientData = {
                fullname: fullnameP,
                email: emailP,
                dni: parseInt(dniP,10),
                phone: parseInt(phoneP,10),
            }

            console.log("PATIENT EDIT DATA: ", patientData)

            fetch(`http://localhost:8080/edit/${idpatient}`,{
                method: "PUT",
                headers: {
                    "Content-Type": "application/json"},
                body: JSON.stringify(patientData),
                credentials: "include",
            })
            .then(response =>{
                if(!response.ok){
                    console.error(response)
                }
                return response.json()
            })
            .then(data =>{
                console.log("Patient edited...")

                const tableBody = document.querySelector(".tabla tbody")
                const row = document.createElement("tr")

                row.innerHTML = `
                        <td>${data.FullName}</td>
                        <td>${data.Email}</td>
                        <td>${data.Dni}</td>
                        <td>${data.Phone}</td>
                        <td>
                            <ul>
                                <div class="buttons">
                                    <a name="" id="" class="btn deleteBtn" data-id="${data.ID}" href="#" role="button">
                                        <i class="fa-regular fa-trash-can"></i>
                                        <span class="tooltip-text">Eliminar</span>
                                    </a>
                                </div>
                                <div class="buttons">
                                    <a name="" id="btnEdit"  class="btn" href="#"  role="button">
                                        <i class="fa-regular fa-pen-to-square"></i>
                                        <span class="tooltip-text">Editar</span>
                                    </a>
                                </div>
                                <div class="buttons">
                                    <a name="" id="" class="btn" href="#" role="button">
                                        <i class="fa-regular fa-calendar"></i>
                                        <span class="tooltip-text">Cita</span>
                                    </a>
                                </div>
                            </ul>
                        </td>
                `;
                tableBody.appendChild(row)
                window.location.reload()
            })
            .catch(error => console.error(error))
        }) 
    }
}

function CreatePatient(){
    const openForm = document.getElementById("openFormBtn")
    const closeForm = document.getElementById("closeFormBtn")
    const formContainer = document.getElementById("formContainer")

    openForm.addEventListener("click", function(){
        formContainer.classList.add("active")
    })

    closeForm.addEventListener("click", function(){
        formContainer.classList.remove("active")
    })

    if(formContainer){
        formContainer.addEventListener("submit", function(e){
            e.preventDefault()

            const fullnameP = document.getElementById("name").value
            const emailP = document.getElementById("email").value
            const dniP = document.getElementById("dni").value
            const phoneP = document.getElementById("phone").value


            const patientData ={
                fullname: fullnameP,
                email: emailP,
                dni: parseInt(dniP,10),
                phone: parseInt(phoneP,10)
            }

            console.log(fullnameP,emailP,dniP,phoneP)

            fetch(`http://localhost:8080/create`,{
                method: "POST",
                headers: {
                    "Content-Type": "application/json"},
                body: JSON.stringify(patientData),
                credentials: "include"
            })
                .then(response =>{
                    if(!response.ok){
                        console.error(response)
                    }
                    return response.json()
                })
                .then(data =>{
                    console.log("Patient created...")

                    const tableBody = document.querySelector(".tabla tbody")
                    const row = document.createElement("tr")

                    row.innerHTML = `
                        <td>${data.FullName}</td>
                            <td>${data.Email}</td>
                            <td>${data.Dni}</td>
                            <td>${data.Phone}</td>
                            <td>
                                <ul>
                                    <div class="buttons">
                                        <a name="" id="" class="btn deleteBtn" data-id="${data.ID}" href="#" role="button">
                                            <i class="fa-regular fa-trash-can"></i>
                                            <span class="tooltip-text">Eliminar</span>
                                        </a>
                                    </div>
                                    <div class="buttons">
                                        <a name="" id="btnEdit"  class="btn" href="#"  role="button">
                                            <i class="fa-regular fa-pen-to-square"></i>
                                            <span class="tooltip-text">Editar</span>
                                        </a>
                                    </div>
                                    <div class="buttons">
                                        <a name="" id="" class="btn" href="#" role="button">
                                            <i class="fa-regular fa-calendar"></i>
                                            <span class="tooltip-text">Cita</span>
                                        </a>
                                    </div>
                                </ul>
                            </td>
                    `;
                    tableBody.appendChild(row)
                    window.location.reload()
                })
                .catch(error => console.error(error))
        })
    }
}

function CreateAppointment(){
    const dayInput = document.getElementById("dayApp")
    const hourSelect = document.getElementById("hourApp")

    dayInput.addEventListener("change", function(){
        const selectedDay = dayInput.value;

        if (selectedDay){
            fetchAvilableHours(selectedDay)
        }
    })

    function fetchAvilableHours(date){
        fetch(`http://localhost:8080/available-hours?fecha=${date}`)
        .then(response => response.json())
        .then(data => {
            console.log(data)
            const availableHours = data.available_hours;
            updateHourOptions(availableHours);
        })
        .catch(error => console.error("Error: ", error))
    }


    function updateHourOptions(hours){
        hourSelect.innerHTML = "";
        console.log(hours)

        if(hours.length > 0){
            hours.forEach( hour => {
                const option = document.createElement('option')
                option.value = hour;
                option.textContent = hour
                hourSelect.appendChild(option)
            })
        }else {
            const option = document.createElement('option');
            option.textContent = "No available hours";
            hourSelect.appendChild(option)
        }
    }


    const btnsumbit = document.getElementById("btnsumbit")
    
    const formContainerAppointment = document.getElementById("formContainerAppoinment")

    if(formContainerAppointment){

        const closeFormApp = document.getElementById("closeFormAppBtn")

        closeFormApp.addEventListener("click", function(){
            formContainerAppointment.classList.remove("active")
        })


        formContainerAppointment.addEventListener("submit", function(e){
            e.preventDefault()

            const pacienteID = formContainerAppointment.getAttribute("data-id")
            console.log(pacienteID)

            const emailPatient = document.getElementById("emailApp").value
            const dayApp = document.getElementById("dayApp").value
            const hour = document.getElementById("hourApp").value
            const motivocons = document.getElementById("motivoApp").value

            console.log(email,dni,dayApp,hour)

            const appData = {
                pacienteid: parseInt(pacienteID),
                fecha: dayApp,
                hora: hour,
                motivoconsulta: motivocons,
            }

            fetch(`http://localhost:8080/create-appointment`,{
                method: "POST",
                headers:{
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(appData),
                credentials: "include",
            })
            .then(response =>{
                if(!response.ok){
                    console.error(response)
                }
                return response.json()
            })
            .then(data =>{
                console.log("Appointment created...")
                if(data.message === "Appoinment created succesfully"){
                    alert("Cita creada correctamente.")
                    window.location.reload()
                }
                
            })

        })
    }
}


function loadPatients(){
    fetch(urlApi,{
        credentials: "include",
    })
    .then(response =>{
        if(!response.ok){
            console.error(response)
        }
        return response.json()
    })
    .then(data => {
        //console.log(data)
        const tablaBody = document.querySelector(".tabla tbody")
        tablaBody.innerHTML = ""

        data.forEach(patient =>{
            const row = document.createElement("tr")

            //console.log(patient.ID)
            
            row.innerHTML = `
                <td>${patient.FullName}</td>
                <td>${patient.Email}</td>
                <td>${patient.Dni}</td>
                <td>${patient.Phone}</td>
                <td>
                    <ul>
                        <div class="buttons">
                            <a name="" id="" class="btn deleteBtn" data-id="${patient.ID}" href="#" role="button">
                                <i class="fa-regular fa-trash-can"></i>
                                <span class="tooltip-text">Eliminar</span>
                            </a>
                        </div>
                        <div class="buttons">
                            <a name="" id="" class="btn editBtn" href="#" data-id="${patient.ID}" role="button">
                                <i class="fa-regular fa-pen-to-square"></i>
                                <span class="tooltip-text">Editar</span>
                            </a>
                        </div>
                        <div class="buttons">
                            <a name="" id="" class="btn appointBtn" href="#" data-id="${patient.ID}" role="button">
                                <i class="fa-regular fa-calendar"></i>
                                <span class="tooltip-text">Cita</span>
                            </a>
                        </div>
                    </ul>
                </td>
        
            `;
            tablaBody.appendChild(row)
        })

        document.querySelectorAll(".editBtn").forEach(button =>{
            button.addEventListener("click", function(){
                
                const id = this.getAttribute("data-id")
                loadPatientData(id)
                formContainerEdit.classList.add("active")
                
            })
        })

        document.querySelectorAll(".appointBtn").forEach(button =>{
            button.addEventListener("click", function(){
                
                const formContainerAppointment = document.getElementById("formContainerAppoinment")
                const id = this.getAttribute("data-id")
                loadPatientDataApp(id)
                formContainerAppointment.classList.add("active")
                
            })
        })

    })
    .catch(error => console.error("Error", error))
}

function searchPatient(){
    const searchInput = document.getElementById("searchInput");
    const resultList = document.querySelector(".tabla tbody");

    searchInput.addEventListener("input", function(){
        const query = searchInput.value

        if(query.length > 0){
            fetch(`http://localhost:8080/search-patient?p=${encodeURIComponent(query)}`)
                .then(response => response.json())
                .then(data => {
                    resultList.innerHTML = "";
                    data.forEach(patient => {
                        const row = document.createElement("tr");
                        row.innerHTML = `
                            <td>${patient.FullName}</td>
                            <td>${patient.Email}</td>
                            <td>${patient.Dni}</td>
                            <td>${patient.Phone}</td>
                            <td>`
                        resultList.appendChild(row)
                    })
                })
                .catch(error => console.error(error))
        }else{
            loadPatients();
        }
    })
}


document.addEventListener("DOMContentLoaded",function(e){
    e.preventDefault();

    ValidateSession()

    loadPatients()

    EditPatients()

    DeletePatients()

    CreatePatient()

    CreateAppointment()

    searchPatient()

})